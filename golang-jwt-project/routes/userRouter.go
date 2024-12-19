package routes

import (
	"golang-jwt-project/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingR *gin.Engine) {
    incomingR.Use(middleware.Authenticate())
    incomingR.Get("/users", controllers.GetUsers())
    incomingR.Get("/users/:user_id", controllers.GetUser())
}