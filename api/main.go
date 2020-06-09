package main

import (
	"github.com/micro/go-micro/web"
	"go-micro-200522/api/router"
	"net/http"
	"time"
)

func main() {
	service := web.NewService(
		web.Name("com.weiguo.api"),
	)
	service.Init()

	engine := router.Router()

	server := http.Server{
		Addr:           ":8080",
		Handler:        engine,
		ReadTimeout:    time.Second * 30,
		WriteTimeout:   time.Second * 30,
		MaxHeaderBytes: 1 << 20,
	}

	service.Handle("/",engine)

	server.ListenAndServe()

	service.Run()

}
