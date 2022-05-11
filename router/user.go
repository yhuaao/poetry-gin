package router

import (
	"poetry/app/controller"
	"poetry/app/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("list", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "pong",
			})
		})
		
		UserRouter.POST("login", controller.UserController{}.UserLogin)

		UserRouter.POST("/user_list",controller.UserController{}.UserList)

		UserRouter.POST("/jwt",controller.JwtController{}.GetJwt)

		UserRouter.POST("paser",middleware.JwtMiddleware(),controller.JwtController{}.PaserToken)
	
	
	}
}
