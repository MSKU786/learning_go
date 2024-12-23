package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)


func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("Unauthorized to access the resouce")	
		return err;
	}

	return err;
}


func MatchUserByUid(c *gin.Context, userId string) (err error) {
	// userId := c.Param("user_id")
	// if userId != c.GetString("user_id") {
	// 	return errors.New("Unauthorized")
	// }

	userType := c.GetString("user_type")
	uId := c.GetString("user_id")

	err = nil 
	if userType == "USER" && userId != uId {
		err = errors.New("Unauthorized to access the resouce")
		return err
	}

	err = CheckUserType(c, userId)
	return err;
}