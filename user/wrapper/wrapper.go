package wrapper

import (
	"context"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	"log"
)

/**
wrapper的目的应该是封装函数
 */
//封装处理器函数，案例是用log包着真实的处理器函数
func HandlerWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Println("HandlerWrapper,在调用真是函数之前")
		//你可以这样进行理解，相当于在这里调用了hello函数
		return fn(ctx,req,rsp)
	}
}

type clientWrapper struct {
	client.Client
}

func(c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	log.Println("wrapper client")
	return c.Client.Call(ctx,req,rsp)
}
// 返回一个包装过的客户端
func LogClientWrap(c client.Client) client.Client {
	return &clientWrapper{c}
}