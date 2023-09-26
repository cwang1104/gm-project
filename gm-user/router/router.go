package router

import (
	"github.com/gin-gonic/gin"
	"gm-user/user"
)

// Router 路由接口
type Router interface {
	Route(r *gin.Engine)
}

type RegisterRouter struct {
}

func New() *RegisterRouter {
	return &RegisterRouter{}
}

func (*RegisterRouter) Route(ro Router, r *gin.Engine) {
	ro.Route(r)
}

func InitRouter(r *gin.Engine) {
	rg := New()
	r.Group("/api")
	rg.Route(&user.RouterUser{}, r)
}
