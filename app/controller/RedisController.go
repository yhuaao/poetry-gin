package controller

import (
	"fmt"
	"poetry/app/util"
	"poetry/global"

	"github.com/gin-gonic/gin"
)

type RedisController struct{}

func (RedisController) Set(c *gin.Context) {
	// name:=c.Query("name")
	// fmt.Print(name)
	global.Redis.Set(global.Redis.Context(),"name",1212,0)
	// util.Success(c,name)
	return
}


func (RedisController) Get(c *gin.Context){
	// name:=c.Query("name")
	namev:=global.Redis.Get(global.Redis.Context(),"name")
	fmt.Printf("%v",namev)
	util.Success(c,namev.Val())
	// c.JSON(http.StatusOK,map[string]interface{}{
	// 	"code":1,
	// 	"data":namev.Val(),
	// })
	return
}