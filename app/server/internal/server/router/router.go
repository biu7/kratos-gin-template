package router

import (
	"github.com/gin-gonic/gin"
	"kratos-gin-template/app/server/internal/middleware"
	"kratos-gin-template/app/server/internal/service"
)

func RegisterHelloRouter(router *gin.Engine, mid *middleware.Middleware, svc *service.GreeterService) {
	router.GET("/hello", svc.SayHello)
}
