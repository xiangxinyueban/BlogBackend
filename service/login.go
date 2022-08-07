package service

import (
	"blogbackend/helper"
	"blogbackend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// Login
// @Summary 用户登录
// @Schemes
// @Description do ping
// @Tags 用户
// @Accept json
// @Produce json
// @Param username formData string false "username"
// @Param password formData string false "password"
// @Success 200 {string} json"{"code":"200", "data":""}
// @Router /user/login [POST]
func Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	verifyId := ctx.PostForm("verifyId")
	verifyCode := ctx.PostForm("verifyCode")
	log.Printf("验证ID：%v, 验证码:%v", verifyId, verifyCode)
	if !VerifyCaptcha(verifyId, verifyCode) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg":  "验证码错误",
		})
		return
	}
	if username == "" || password == "" {
		log.Printf("异常登录，空用户名或密码")
		ctx.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg":  "用户名密码不能为空",
		})
		return
	}
	password = helper.GetMd5(password)
	data := new(models.User)
	err := models.DB.Where("name = ? AND password = ?", username, password).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusOK, gin.H{
				"code": "403",
				"msg":  "用户名或密码错误",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": "404",
			"msg":  "Get UserBasic Error" + err.Error(),
		})
		return
	}
	token, err := helper.GenerateToken(data.Id, data.Name)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusForbidden,
			"msg":  "Generate token failed" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": map[string]interface{}{
			"token": token,
		},
	})

}
