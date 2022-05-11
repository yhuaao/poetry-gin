package initialize

import (
	"poetry/app/middleware"
	"poetry/router"

	"github.com/gin-gonic/gin"
)
func Routers() *gin.Engine {
	Router := gin.Default()

	// Router.StaticFile("/", "./static/dist/index.html")
    // Router.Static("/assets", "./static/dist/assets")
    // Router.StaticFile("/favicon.ico", "./static/dist/favicon.ico")
    // // 其他静态资源
    // Router.Static("/public", "./static")
    // Router.Static("/storage", "./storage/app/public")

	//使用中间件
	Router.Use(
		middleware.LoggerMiddleware(),
		middleware.RecoveryMiddleware(true),
		middleware.CorsMiddleware());
	// 路由分组
	ApiGroup := Router.Group("/v1/")
	router.UserRouter(ApiGroup)
	router.CatpchaRouter(ApiGroup)
	return Router
}
