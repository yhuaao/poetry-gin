package initialize

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"poetry/global"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func InitReloadServer(router *gin.Engine) {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", global.Settings.Port),
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
			// color.Cyan("go-gin服务开始了")
		}
	}()
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}