package middleware

import (
	"fmt"
	"golang-jwt-project/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
		return func(c * gin.Context) {
			clientToken := c.Request.Header.Get("token")
			if clientToken == "" {
				c.JSON(http.StatusInternalServerError, gin.H{"error", fmt.Sprintf("No authorization header found")})
				c.abort()
				return;
			}

			claims, err := helpers.ValidateToken(clientToken)	//validate the token
			if err != "" {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				c.abort()
				return;
			}

			c.Set("email", claims.Email);
			c.Set("user_id", claims.UserId);
			c.Set("first_name", claims.First_Name);
			c.Set("last_name", claims.Last_Name);
			c.Set("user_type", claims.User_Type);
			c.Next();
			
		}
}