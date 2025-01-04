package helpers

import (
	"context"
	"golang-jwt-project/database"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type SignedDetails struct {
	Email string
	UserId string
	First_Name string
	Last_Name string
	User_Type string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")	//database connection
var SECRET_KEY string = os.Getenv("SECRET_KEY")


func GenerateAllTokens(email string, firstName string, lastName string, userType string, userId string) (string, string) {
	var err error
	//Creating Access Token
	claims := &SignedDetails{
		Email: email,
		UserId: userId,
		First_Name: firstName,
		Last_Name: lastName,
		User_Type: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshtokenClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(8760)).Unix(),
	},
}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY));

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshtokenClaims).SignedString([]byte(SECRET_KEY));	

	if err != nil {
		log.Panic(err);
		return "", ""
	}

	return token, refreshToken, err;

}


func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second);
	
	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{"token", signedToke});
	updateObj = append(updateObj, bson.E{"refresh_token", signedRefreshToke});
	updateObj = append(updateObj, bson.E{"updated_at", time.Now().Local()});

	upsert := true
	filter := bson.M{"user_id": userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := userCollection.UpdateOne(ctx, filter, bson.D{{"$set", updateObj}}, &opt);

	if (err != nil) {
		log.Panic(err)
		return;
	}
	 
	return;
}