package routes


func AuthRoutes (r *gin.Engine) {
	r.POST("users/signup", controllers.SingUp)
	r.POST("users/login", controllers.LogIn)	 
}