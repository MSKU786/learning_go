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

			status, err := helpers.Vali(clientToken)
		}
}