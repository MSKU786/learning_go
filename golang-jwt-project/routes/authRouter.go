package routes

import (
	controllers "golang-jwt-project/controllers"

	"github.com/gin-gonic/gin"
)


func AuthRoutes (r *gin.Engine) {
	r.POST("users/signup", controllers.SignUp())
	r.POST("users/login", controllers.Login())	 
}