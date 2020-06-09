package client

import (
	"context"
	"go-micro-200522/user/proto"
	"log"
	"testing"
)

//测试client究竟能不能连接的上
func TestUserClient(t *testing.T) {
	client := UserClient()
	response, e := client.CreateUser(context.Background(), &proto.User{
		Nickname: "se7en",
		Username: "weiguozhao",
		Password: "1234",
		Phone:    "17090081561",
		Email:    "movieqizongzui@163.com",
		Birthday: 19900217,
	})
	if e != nil {
		log.Println("create user error",e)
	}
	log.Println(response.Msg)
	log.Println(response.Success)
}
