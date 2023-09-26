package main

import (
	"github.com/gin-gonic/gin"
	common "gm-common"
	"gm-user/router"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	common.Run(r, ":8080", "gm-user")

}
