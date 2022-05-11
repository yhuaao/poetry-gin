package middleware

import (
	"fmt"
	"poetry/app/util"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/color"
)


func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		// token:="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MTgsIlBob25lIjoiMTY2MDU1MDE3MjAiLCJBdXRob3JpdHlJZCI6MSwiZXhwIjoxNjU0ODI1NjAwLCJpc3MiOiJ0ZXN0IiwibmJmIjoxNjUyMjMzNjAwfQ.7ke-9YIs4n_Q43JbraGX39l44zKT05Ab7YqKGzzzPp4";
		color.Yellow(token)
		if token == "" {
			util.Error(c,"请登录")
			c.Abort()
			return
		}
		jwt := util.NewJWT()
		claims, err := jwt.ParseToken(token)
		if err != nil {
			if err == util.TokenExpired {
				if err == util.TokenExpired {
					util.Error(c, "授权已过期")
					c.Abort()
					return
				}
			}
			util.Error(c,"未登陆")
			c.Abort()
			return
		}
		fmt.Println(c)
		// gin的上下文记录claims和userId的值
		c.Set("claims", claims)
		c.Set("userId", claims.ID)
		c.Next()
	}
}