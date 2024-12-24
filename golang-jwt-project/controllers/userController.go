package controllers

import (
	"context"
	"golang-jwt-project/database"
	"golang-jwt-project/models"
	"log"
	"net/http"

	"golang-jwt-project/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")
var validate = validator.New()

func Singup() gin.HandlerFunc{
		return func(c *gin.Context) {
			var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
			var user models.User
			if err := c.BindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if err := validate.Struct(user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})

			defer cancel()

			if err != nil {
				log.Panic(err);
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while checking for the user"})
				return
			}

			if count > 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "User email already exists"})
				return
			}

			count,err := userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})

			if err != nil {
				log.Panic(err);
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while checking for the user"})
				return
			}

			if count > 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "User phone number already exists"})
				return
			}

			user.Password, _ = helpers.HashPassword(user.Password)
			user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			user.ID = primitive.NewObjectID()
			user.UserId = user.ID.Hex()

			result, insertErr := userCollection.InsertOne(ctx, user)
			defer cancel()

			if insertErr != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating user"})
				return
			}

			c.JSON(http.StatusOK, result)
		}  
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