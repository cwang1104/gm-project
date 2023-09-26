package main

import (
	"log"
	"os"

	common "gm-common"
	"gm-common/logs"
	"gm-user/router"

	"github.com/gin-gonic/gin"
)

func main() {

	dir, err := os.Getwd()
	if err != nil {
		log.Println("获取root路径失败", err)
		panic(err)
	}

	lc := &logs.LogConfig{
		DebugFileName: dir + "/logs/debug/gm-debug.log",
		InfoFileName:  dir + "/logs/info/gm-info.log",
		WarnFileName:  dir + "/logs/error/gm-error.log",
		MaxSize:       500,
		MaxAge:        28,
		MaxBackups:    3,
	}

	err = logs.InitLogger(lc)
	if err != nil {
		log.Panicln(err)
	}

	r := gin.Default()

	router.InitRouter(r)
	common.Run(r, ":8080", "gm-user")

}
