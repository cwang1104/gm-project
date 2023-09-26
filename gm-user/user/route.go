package user

import "github.com/gin-gonic/gin"

type RouterUser struct {
}

func (*RouterUser) Route(r *gin.Engine) {
	h := &HandlerUser{}

	r.POST("/project/login", h.GetCap)
}
