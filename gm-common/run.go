package common

import (
	"context"
	"errors"
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
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln("http listen failed;err:", err)
		}
		log.Printf("%s running in %s", serverName, addr)
	}()

	//只有sigint sigterm 两个信号可以被监听
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("%s shtdown,case by %s", addr, err.Error())
	}

	select {
	case <-ctx.Done():
		log.Println("wait timeout...")
	}

	log.Printf("%s stop success...", serverName)

}
