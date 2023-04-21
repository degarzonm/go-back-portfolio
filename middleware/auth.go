package middleware

import (
	"fmt"
	"net/http"

	utils "github.com/degarzonm/go-back-portfolio/utils"

	"github.com/gin-gonic/gin"
)

// Authz validates token and authorizes users
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("uid", claims.Uid)
		c.Set("name", claims.Name)
		c.Set("city", claims.City)

		c.Next()

	}
}
