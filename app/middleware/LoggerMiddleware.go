package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// 请求路径
		path := c.Request.URL.Path
		// 请求参数
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		// 若response的状态码不是200为异常
		if c.Writer.Status() != 200 {
			// 记录异常信息
			zap.L().Info(path,
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.Duration("cost", cost),
			)
		}
	}
}




 