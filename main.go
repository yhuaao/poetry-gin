package main

import (
	"poetry/global"
	"poetry/initialize"
)

func main() {
	// 初始化配置
    initialize.InitConfig()
    initialize.InitLogger()
    if err := initialize.InitTrans("zh"); err != nil {
        panic(err)
    }
    initialize.InitMySqlGorm()
    initialize.InitRedisPool()
    Router := initialize.Routers()
	defer func() {
        if global.DB != nil {
            db, _ := global.DB.DB()
            db.Close()
        }
    }()
    // color.Cyan("go-gin服务开始了")
	initialize.InitReloadServer(Router)
}

