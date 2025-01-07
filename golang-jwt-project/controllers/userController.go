package controllers

import (
	"context"
	"golang-jwt-project/database"
	"golang-jwt-project/models"
	"log"
	"net/http"
	"os/user"
	"strconv"

	"golang-jwt-project/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")
var validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

		if err != nil {
			log.Panic(err)
		}	
			
		return string(bytes)
}


func VerifyPassword(userPassword string, providePassword string)(bool string) {
		err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providePassword))
		check := true
		msg := ""

		if err != nil {	
			msg = "Password does not match"
			check = false
		}	

		return check, msg;
	}


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

			token, refreshToken := helpers.GenerateAllTokens(*user.Email, *user.First_Name, *user.Last_Name, *user.User_Type,  *&user.UserId)
			user.Token = &token
			user.RefreshToken = &refreshToken

			result, insertErr := userCollection.InsertOne(ctx, user)
			defer cancel()

			if insertErr != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating user"})
				return
			}

			c.JSON(http.StatusOK, result)
		}  
}

func Login() ginHandler {
	return func(c *gin.Context) { 
			ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)	;
			var user models.User;
			var foundUser models.User;
			if err := c.BindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}	

			if err := validate.Struct(user); err != nil {	
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}	

			err :- userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser);

			if err != nil {
				log.Panic(err);
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Email or password is incorrect"})
				return
			}

			passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password);
			defer cancel()

			if (passwordIsValid != true) {
				c.JSON(http.internalServerError, gin.H{"error": msg})	
				return
			}

			token, refreshToken, _  := helpers.GenerateAllTokens(*foundUser.Email, *foundUser.First_Name, *foundUser.Last_Name, *foundUser.User_Type, *foundUser.UserId)

			err := userCollection.FindOne(ctx, bson.M{"user_id": foundUser.User_id}).Decode(&foundUser) 
			defer cancel()

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
				return
			}

			 
			c.JSON(http.StatusOK, foundUser)

	}
}

func GetUsers() gin.HandlerFunc{
		return func(c *gin.Context) {
			err := helpers.CheckUserType(c, "ADMIN"); 
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return;
			}

			var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second);

			recordPerPage, err := strconv.Atomic.query("recordPerPage")
			if err != nil || recordPerPage < 1 {
				recordPerPage = 10
			}

			page, err1 := strconv.Atomic.query("page")

			if err1 != nil || page < 1 {
				page = 1
			}
			
			startIndex := (page - 1) * recordPerPage;
			startIndex, err = strconv.Atoi(c.Query("startIndex"));


			matchStage := bson.D{{Key: "$match", Value: bson.D{{}}}};
			groupStage := bson.D{{Key: "$group", Value: bson.D{
				{Key: "_id", Value: bson.D{{Key: "_id", Value: "null"}}}, 
				{Key: "total", Value: bson.D{{Key: "$sum", Value: 1}}}, 
				{Key: "data", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}},
			}}};
			projectStage := bson.D{
				{Key: "$project", Value: bson.D{
					{Key: "_id", Value: 0},
					{Key: "total", Value: 1},
					{Key: "user_items", Value: bson.D{{Key: "$slice", Value: []interface{}{"$data", startIndex, recordPerPage}}}},
				},
			}}
			
			results , err := userCollection.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage, projectStage});
			defer cancel();
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while fetching users"})
				return
			}

			var allUsers []bson.M;

			if err = results.All(ctx, &allUsers); err != nil {
				log.Fatal(err);
		
			}
			c.JSON(http.StatusOK, allUsers[0]);

		}
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