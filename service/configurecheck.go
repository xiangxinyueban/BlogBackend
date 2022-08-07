package service

import (
	"blogbackend/conf"
	"blogbackend/models"
	"github.com/gin-gonic/gin"
	"log"
)

// ConfigureCheck
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags 配置检查
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router / [get]
func ConfigureCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"port":     conf.Port,
		"user":     conf.MySqlUser,
		"password": conf.MySqlPassword,
	})
	log.Println(models.DB)
}
