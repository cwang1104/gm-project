package user

import "github.com/gin-gonic/gin"

type RouterUser struct {
}

func (*RouterUser) Route(r *gin.Engine) {
	InitRpcUserClient()
	h := New()

	r.POST("/project/login", h.GetCaptcha)
}
