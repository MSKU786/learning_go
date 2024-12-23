package controllers

import (
	"context"
	"golang-jwt-project/database"
	"golang-jwt-project/models"
	"net/http"

	"golang-jwt-project/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")
var validate = validator.New()

func Singup() {

}

func Login() {

}

func GetUsers() {

}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
			userId := c.Param("user_id")
			if err := helpers.MatchUserByUid(c, userId); err != nil {
			if err := helper.MatchUserByUid(c, userId); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
				return
			}

			var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
			
			var user models.User
			err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user) 
			defer cancel()

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
				return
			}

			c.JSON(http.StatusOK, user)
	}
}