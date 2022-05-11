package controller

import (
	"poetry/app/util"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type CatpchaController struct{}

var store = base64Captcha.DefaultMemStore
func (CatpchaController) GetCaptcha(c *gin.Context) {
	//
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	// b64s是图片的base64编码
	id, b64s, err := cp.Generate()
	if err != nil {
		zap.S().Errorf("生成验证码错误,:%s ", err.Error())
		util.Error(c,  "生成验证码错误")
		return
	}
	util.Success(c, gin.H{
		"captchaId": id,
		"picPath":   b64s,
	})
	return
}
