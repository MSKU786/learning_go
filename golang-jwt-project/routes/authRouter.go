package routes

import (
	"golang-jwt-project/controllers"
)


func AuthRoutes (r *gin.Engine) {
	r.POST("users/signup", controllers.SingUp())
	r.POST("users/login", controllers.LogIn())	 
}