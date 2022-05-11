package initialize

import (
	"fmt"

	"poetry/app/util"
	"poetry/global"

	"go.uber.org/zap"
)

func InitLogger() {

	//创建文件夹
	util.CreateLogDir();

	cfg := zap.NewDevelopmentConfig()
        // 注意global.Settings.LogsAddress是在settings-dev.yaml配置过的
        // 配置日志的输出地址
	cfg.OutputPaths = []string{
		fmt.Sprintf("%slog_%s.log", global.Settings.LogsAddress, util.GetNowFormatTodayTime()), //
		"stdout",
	}
        // 创建logger实例
	logg, _ := cfg.Build()
	// logg.Sugar()
	zap.ReplaceGlobals(logg) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
    global.Lg = logg  // 注册到全局变量中

}
