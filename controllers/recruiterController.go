package controllers 


import (
	"context"
	"fmt"
	"net/http"
	"time"
	"github.com/degarzonm/go-back-portfolio/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)



func AscendSoldier() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestBody struct {
			UserID string `json:"user_id" binding:"required"`
		}

		recruiterID := c.MustGet("user_id").(string)
		user, err := GetUserByID(recruiterID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting Recruiter: %v", err)})
			return
		}

		// Check if the user is a recruiter
		if user.Role != middleware.RoleRecruiter {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to ascend a soldier"})
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

		// Get the soldier to be ascended
		soldier, err := GetUserByID(requestBody.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting Soldier: %v", err)})
			return
		}

		// Ascend the soldier
		var newRole middleware.Role
		switch soldier.Role {
		case middleware.RoleSoldier:
			newRole = middleware.RoleOfficer
		case middleware.RoleOfficer:
			newRole = middleware.RoleGeneral
		case middleware.RoleGeneral:
			c.JSON(http.StatusBadRequest, gin.H{"error": "The user is already a general"})
			return
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "The user has an invalid role"})
			return
		}

		// Update the soldier's role
		update := bson.M{"$set": bson.M{"role": newRole}}
		filter := bson.M{"user_id": requestBody.UserID}
		updateResult, updateErr := SoldierCollection.UpdateOne(ctx, filter, update)

		if updateErr != nil {
			msg := fmt.Sprintf("Error updating soldier: %v", updateErr)
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		if updateResult.ModifiedCount == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No soldier found with the provided user_id"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Soldier ascended successfully"})
	}
}

func Jubilate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestBody struct {
			UserID string `json:"user_id" binding:"required"`
		}

		recruiterID := c.MustGet("user_id").(string)
		user, err := GetUserByID(recruiterID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting Recruiter: %v", err)})
			return
		}

		// Check if the user is a recruiter
		if user.Role != middleware.RoleRecruiter {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to jubilate a soldier"})
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

		// Delete the soldier with the given user_id
		filter := bson.M{"user_id": requestBody.UserID}
		deleteResult, deleteErr := SoldierCollection.DeleteOne(ctx, filter)

		if deleteErr != nil {
			msg := fmt.Sprintf("Error deleting soldier: %v", deleteErr)
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		if deleteResult.DeletedCount == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No soldier found with the provided user_id"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Soldier jubilated successfully"})
	}
}