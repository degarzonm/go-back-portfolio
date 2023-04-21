package controllers

import (
	"context"
	"fmt"
	"log"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/degarzonm/go-back-portfolio/database"
	"github.com/degarzonm/go-back-portfolio/middleware"

	"github.com/degarzonm/go-back-portfolio/models"
	utils "github.com/degarzonm/go-back-portfolio/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var SoldierCollection *mongo.Collection = database.OpenCollection(database.DBClient, "soldier")
var validate = validator.New()

// GetUserByID retrieves a user by their ID from the database.
func GetUserByID(userID string) (*models.Soldier, error) {
	var soldier models.Soldier 

	// Define a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find the user with the given ID
	err := SoldierCollection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&soldier)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("User not found")
		}
		return nil, err
	}

	return &soldier, nil
}


// HashPassword is used to encrypt the password before it is stored in the DB
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

// VerifyPassword checks the input password while verifying it with the passward in the DB.
func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("login or password is incorrect")
		check = false
	}

	return check, msg
}

// CreateUser is the api used to tget a single user
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.Soldier

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		dataRepeated, err := SoldierCollection.CountDocuments(ctx, bson.M{"name": user.Name})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the name"})
			return
		}

		dataRepeated, err = SoldierCollection.CountDocuments(ctx, bson.M{"user_id": user.UserID})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the id number"})
			return
		}

		if dataRepeated > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this name or id number already exists"})
			return
		}
		password := HashPassword(*user.Password)
		user.Password = &password

		user.ID = primitive.NewObjectID()
		token, refreshToken, _ := utils.GenerateAllTokens(*user.Name, *user.UserID, *user.City)
		user.Token = &token
		user.RefreshToken = &refreshToken
		user.Role = middleware.RoleSoldier
		user.FortIDs = []string{"reserva"} 

		resultInsertionNumber, insertErr := SoldierCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("Soldier was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, resultInsertionNumber)

	}
}

// Login is the api used to get a single user
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.Soldier
		var foundUser models.Soldier

		// bind the json body to the user struct
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := SoldierCollection.FindOne(ctx, bson.M{"user_id": user.UserID}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "login or password is incorrect"})
			return
		}

		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if passwordIsValid != true {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		token, refreshToken, _ := utils.GenerateAllTokens(*foundUser.Name, *foundUser.UserID, *foundUser.City)

		utils.UpdateAllTokens(token, refreshToken, *foundUser.UserID)

		responseData := models.SoldierLoginResponse{
			UserID:       foundUser.UserID,
			Name:         foundUser.Name,
			City:         foundUser.City,
			FortIDs:      foundUser.FortIDs,
			Token:        &token,
			RefreshToken: &refreshToken,
		}

		c.JSON(http.StatusOK, responseData)

	}
}


