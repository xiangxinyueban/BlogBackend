package service

import (
	"github.com/gin-gonic/gin"
	"log"
)

// Register
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags 用户管理
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /user/register [post]
func Register(ctx *gin.Context) {
	name := ctx.PostForm("username")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	log.Println(name, email, password)

	ctx.JSON(200, gin.H{
		"msg": "注册成功",
	})
}
