package main

import (
	routes "golang-jwt-project/routes"
	"os"

	"github.com/gin-conic/gin"
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

		router.get('/api/v1', func(c *gin.Context) {
			
		})
}