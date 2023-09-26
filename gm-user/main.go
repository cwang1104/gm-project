package main

import (
	"log"
	"os"

	common "gm-common"
	"gm-common/logs"
	"gm-user/pkg/config"
	"gm-user/pkg/dao"
	"gm-user/router"

	"github.com/gin-gonic/gin"
)

func main() {

	err := config.Init()
	if err != nil {
		log.Panicln("config unmarshal failed", err)
	}

	err = dao.RedisInit()
	if err != nil {
		log.Panicln("redis init failed", err)
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Println("获取root路径失败", err)
		panic(err)
	}

	lc := &logs.LogConfig{
		DebugFileName: dir + config.ZapConf().DebugFileName,
		InfoFileName:  dir + config.ZapConf().InfoFileName,
		WarnFileName:  dir + config.ZapConf().WarnFileName,
		MaxSize:       config.ZapConf().MaxSize,
		MaxAge:        config.ZapConf().MaxAge,
		MaxBackups:    config.ZapConf().MaxBackup,
	}

	err = logs.InitLogger(lc)
	if err != nil {
		log.Panicln(err)
	}

	r := gin.Default()

	router.InitRouter(r)
	common.Run(r, config.ServerConf().GetAddr(), config.ServerConf().Name)

}
