package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/degarzonm/go-back-portfolio/middleware"
	"github.com/degarzonm/go-back-portfolio/database"
	"github.com/degarzonm/go-back-portfolio/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var FortCollection *mongo.Collection = database.OpenCollection(database.DBClient, "fort")

func GetFortByID(fortID string) (*models.Fort, error) {
	var fort models.Fort 

	// Define a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find the user with the given ID
	err := FortCollection.FindOne(ctx, bson.M{"fort_id": fortID}).Decode(&fort)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("Fort not found")
		}
		return nil, err
	}

	return &fort, nil
}


func GetMyFort() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user's ID from the token
		userID := c.MustGet("user_id").(string)

		// Get the user's forts
		user, err := GetUserByID(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting Soldier: %v", err)})
			return
		}

		if len(user.FortIDs) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Soldier has no forts"})
			return
		}

		// Get the first fort's information
		firstFortID := user.FortIDs[0]
		fort, err := GetFortByID(firstFortID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting fort: %v", err)})
			return
		}

		c.JSON(http.StatusOK, fort)
	}
}


func CreateFort() gin.HandlerFunc{
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var body_fort models.Fort
		generalID := c.MustGet("user_id").(string)//get the user's id from the token
		user, err := GetUserByID(generalID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting Soldier: %v", err)})
			return
		}

		// Check if the user is a general
		if user.Role != middleware.RoleGeneral {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to create a fort"})
			return
		}

		//bind the request body to the struct and validate it 
		if err := c.BindJSON(&body_fort); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(body_fort)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		//check if the fort's name or id number already exists
		dataRepeated, err := FortCollection.CountDocuments(ctx, bson.M{"name": body_fort.Name})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the fort's name"})
			return
		}
		dataRepeated, err = FortCollection.CountDocuments(ctx, bson.M{"fort_id": body_fort.FortID})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the fort's id number"})
			return
		}
		
		if dataRepeated > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this name or id number already exists for this fort"})
			return
		}

		//set the fort's commander to the user's id
		body_fort.Commander = user.UserID
		
		resultInsertionNumber, insertErr := FortCollection.InsertOne(ctx, body_fort)
		if insertErr != nil {
			msg := fmt.Sprintf("Fort was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}

func SetFortCommander() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestBody struct {
			FortID    string `json:"fort_id" binding:"required"`
			Commander string `json:"commander" binding:"required"`
		}

		generalID := c.MustGet("user_id").(string)
		user, err := GetUserByID(generalID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting Soldier: %v", err)})
			return
		}

		// Check if the user is a general
		if user.Role != middleware.RoleGeneral {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to set a fort commander"})
			return
		}

		// Bind the request body to the struct and validate it
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(requestBody)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		// Check if the commander is a valid user and has RoleOfficer
		commanderUser, err := GetUserByID(requestBody.Commander)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting Commander: %v", err)})
			return
		}

		if commanderUser == nil || commanderUser.Role != middleware.RoleOfficer {
			c.JSON(http.StatusBadRequest, gin.H{"error": "The provided commander ID is not a valid officer"})
			return
		}

		// Update the fort's commander with the given commander and fort_id
		update := bson.M{"$set": bson.M{"commander": requestBody.Commander}}
		filter := bson.M{"fort_id": requestBody.FortID}
		updateResult, updateErr := FortCollection.UpdateOne(ctx, filter, update)

		if updateErr != nil {
			msg := fmt.Sprintf("Error updating fort: %v", updateErr)
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		if updateResult.ModifiedCount == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No fort found with the provided fort_id"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Fort commander updated successfully"})
	}
}