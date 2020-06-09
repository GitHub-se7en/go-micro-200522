package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-micro-200522/api/router/api/v1"
)

func Router() *gin.Engine {
	router := gin.New()
	api := router.Group("api/v1")
	{
		api.POST("/user/registry",v1.RegistryUser)
		api.POST("/user/login",v1.UserLogin)
	}
	return router
}
