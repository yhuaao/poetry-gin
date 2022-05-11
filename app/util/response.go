package util

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data interface{}) {
	code:=1;
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": code, // 自定义code
		"data": data, // 数据
	})
	return
}

func Error(c *gin.Context,  msg interface{}) {
	code:=0;
	dataType , _ := json.Marshal(msg)
	dataString := string(dataType)

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": code,
		"msg":  dataString,
	})
	return
}
