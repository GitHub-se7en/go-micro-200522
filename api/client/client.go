package client

import (
	"github.com/micro/go-micro/client"
	"go-micro-200522/user/proto"
)
//user服务客户端
func UserClient() proto.UserServiceClient {
	return proto.NewUserServiceClient("com.weiguo.srv.user",client.DefaultClient)
}