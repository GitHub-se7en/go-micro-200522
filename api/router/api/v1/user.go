package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-micro-200522/api/cache"
	userClient "go-micro-200522/api/client"
	"go-micro-200522/api/util"
	"go-micro-200522/user/proto"
	"log"
	"time"
)

var userServiceClient = userClient.UserClient()

func RegistryUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	nickname := c.PostForm("nickname")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	user := &proto.User{
		Nickname: nickname,
		Username: username,
		Password: password,
		Phone:    phone,
		Email:    email,
	}
	response, err := userServiceClient.CreateUser(context.Background(), user)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "error",
			"code":    -1,
		})
		return
	}
	result := gin.H{}
	if response.Success {
		result["message"] = "success"
		result["code"] = 0
	} else {
		result[""] = response.Msg
		result["code"] = -1
	}
	c.JSON(200, result)

}
func UserLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	response, err := userServiceClient.UserLogin(context.Background(), &proto.LoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		c.JSON(500, gin.H{
			"message": "fail",
		})
		return
	}
	expiration := time.Hour * 24 * 3
	token, err := util.GenerateToken(response.Nickname, expiration)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "fail",
		})
		return
	}
	cache.Redis.Set(token,response.UserId,expiration)
	c.JSON(200,gin.H{"token":token})

}
