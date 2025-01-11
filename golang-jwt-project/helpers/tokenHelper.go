package helpers

import (
	"fmt"
	"context"
	"golang-jwt-project/database"
	"log"
	"os"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
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

	updateObj = append(updateObj, bson.E{Key: "token", Value: signedToken});
	updateObj = append(updateObj, bson.E{Key: "refresh_token", Value: signedRefreshToken});
	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: time.Now().Local()});

	upsert := true
	filter := bson.M{"user_id": userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	defer cancel();

	_, err := userCollection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateObj}}, &opt);

	if (err != nil) {
		log.Panic(err)
		return;
	}
	 
	return;
}


func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken, 
		&SignedDetails{}, 
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error();
		return;
	}

	claims, ok := token.Claims.(*SignedDetails);

	if !ok {
		msg = fmt.Sprintf("the token is invalid");
		msg = err.Error();
		return;
	}	

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token has expired");
		msg = err.Error();
		return;
	}

	return claims, msg;
}