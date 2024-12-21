package controllers

import (
	"golang-jwt-project/database"
	"net/http"

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

			if err := helper.MatchUserByUid(c, userId); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			}

			
	}
}