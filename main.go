package main
import (
	
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the MongoDB connection
	database.Connect()

	// Set up the Gin Gonic server
	r := gin.Default()

	// Use JWT authentication middleware
	r.Use(middleware.JWTAuthMiddleware())

	// Set up the routes
	controllers.SetUpRoutes(r)

	// Run the server
	r.Run(":8080") 
}
