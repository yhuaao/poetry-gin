package router

import (
	"poetry/app/controller"

	"github.com/gin-gonic/gin"
)

func CatpchaRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("catpcha")
	{
		UserRouter.GET("get_captcha",controller.CatpchaController{}.GetCaptcha)


		UserRouter.GET("set",controller.RedisController{}.Set);

		UserRouter.GET("get",controller.RedisController{}.Get);

	}
}