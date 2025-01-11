package routes

import (
	middleware "golang-jwt-project/middleware"

	controllers "golang-jwt-project/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingR *gin.Engine) {
    incomingR.Use(middleware.AuthMiddleware())
    incomingR.GET("/users", controllers.GetUsers())
    incomingR.GET("/users/:user_id", controllers.GetUser())
}