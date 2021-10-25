package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"web-service-gin/models"
)

type Login interface {
	GetJWTToken() string
	LogOut(ctx *gin.Context) bool
}
type User struct {
	models.User
}

func (user *User) GetJWTToken() string {
	json, err := user.ToJson()
	if err != nil {
		return ""
	}
	fmt.Println(json)
	fmt.Println(user.UserName)
	return ""
}

func (user *User) LogOut(ctx *gin.Context) (res bool, err error) {
	return true, nil
}
