package middleware

import (
	"fmt"
	"net/http"
	"restaurant-management/helpers"

	"github.com/gin-gonic/gin"
)

func Authentification() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")

		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No authorization header provided")})
			c.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)

		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf(err)})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("firstname", claims.Firstname)
		c.Set("lastname", claims.Lastname)
		c.Set("uid", claims.Uid)

		c.Next()
	}
}
