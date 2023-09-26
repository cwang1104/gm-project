package common

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Run(r *gin.Engine, addr, serverName string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		zap.L().Info("server running", zap.String("serverName", serverName), zap.String("addr", addr))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.L().Fatal("http listen failed", zap.Error(err))
		}
	}()

	//只有sigint sigterm 两个信号可以被监听
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("%s shtdown,case by %s", addr, err.Error())
	}

	log.Printf("%s is shutting down", serverName)

	select {
	case <-ctx.Done():
		log.Println("wait timeout...")
	}

	log.Printf("%s stop success...", serverName)

}
