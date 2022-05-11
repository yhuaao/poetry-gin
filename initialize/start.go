package initialize

import (
	"fmt"
	"poetry/global"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RunServer(Router *gin.Engine) {
	color.Cyan("go-gin服务开始了")
	// 启动gin,并配置端口,global.Settings.Port是yaml配置过的
	err := Router.Run(fmt.Sprintf(":%d", global.Settings.Port))
	if err != nil {
		zap.L().Info("this is hello func", zap.String("error", "启动错误!"))
	}
}