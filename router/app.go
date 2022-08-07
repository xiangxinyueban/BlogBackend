package router

import (
	docs "blogbackend/docs"
	"blogbackend/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.Version = "0.1"
	r.GET("/", service.ConfigureCheck)
	r.POST("/user/register", service.Register)
	r.POST("/user/login", service.Login)
	r.GET("/safety/captcha", service.GenerateCaptcha)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
