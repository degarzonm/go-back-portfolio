package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/degarzonm/go-back-portfolio/database"
	"github.com/degarzonm/go-back-portfolio/middleware"
	"github.com/degarzonm/go-back-portfolio/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var AgendaCollection *mongo.Collection = database.OpenCollection(database.DBClient, "agenda")

func ModifyAgenda() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestBody struct {
			File string `json:"file" binding:"required"`
		}

		officerID := c.MustGet("user_id").(string)
		user, err := GetUserByID(officerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting Officer: %v", err)})
			return
		}

		// Check if the user is an officer
		if user.Role != middleware.RoleOfficer && user.Role != middleware.RoleGeneral {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to modify the agenda"})
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

		// Find the fort where the officer is the commander
		fortFilter := bson.M{"commander": officerID}
		var fort models.Fort
		fortErr := FortCollection.FindOne(ctx, fortFilter).Decode(&fort)
		if fortErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error finding Fort: %v", fortErr)})
			return
		}

		// Update or create the agenda
		agendaFilter := bson.M{"fort_id": fort.FortID}
		update := bson.M{"$set": bson.M{"file": requestBody.File}}
		upsert := options.Update().SetUpsert(true)
		updateResult, updateErr := AgendaCollection.UpdateOne(ctx, agendaFilter, update, upsert)

		if updateErr != nil {
			msg := fmt.Sprintf("Error updating Agenda: %v", updateErr)
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		if updateResult.UpsertedCount > 0 {
			c.JSON(http.StatusOK, gin.H{"message": "Agenda created successfully"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Agenda updated successfully"})
		}
	}
}
