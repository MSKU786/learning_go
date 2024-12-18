package main

import (
	routes "golang-jwt-project/routes"
	"os"

	"github.com/gin-gonic/gin"
)


func main() {
		port := os.Getenv("PORT")

		if port == "" {
			port = "8000"
		}

		router := gin.New()
		router.Use(gin.Logger())

		routes.AuthRoutes(router);
		routes.UserRoutes(router);

		router.get("/api/v1", func(c *gin.Context) {
				c.JSON(200, gin.H{"success", "Access granted for API v1"})
		})

		router.get("/api/v2", func(c *gin.Context)  {
				c.JSON(200, gin.H("Success", "Access granted for API v2"))
		})

		router.Run(":" port);
}