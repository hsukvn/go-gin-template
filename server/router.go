package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hsukvn/go-gin-template/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	ping := new(controller.PingController)

	r.GET("/ping", ping.Ping)

	return r
}
