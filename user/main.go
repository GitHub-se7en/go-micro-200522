package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config/cmd"
	"go-micro-200522/user/handler"
	"go-micro-200522/user/proto"
	"go-micro-200522/user/wrapper"
	"log"
)

func main() {
	//本地运行和docker运行是完全不一样的，docker的话，是运行在一个虚拟机里面，从虚拟机里面访问windows本地的mongodb，当然是会报错的
	service := micro.NewService(
		micro.Name("com.weiguo.srv.user"),
		micro.Version("latest"),
		micro.WrapHandler(wrapper.HandlerWrapper),
		micro.WrapClient(wrapper.LogClientWrap),
	)

	err := cmd.Init()
	if err != nil {
		log.Fatalf("cmd init error: %v", err)
	}

	service.Init()

	proto.RegisterUserServiceHandler(service.Server(), new(handler.UserHandler))


	err = service.Run()

	if err != nil {
		log.Fatalln("err", err)
	}

}
