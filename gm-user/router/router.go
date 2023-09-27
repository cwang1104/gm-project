package router

import (
	"github.com/gin-gonic/gin"
	"gm-user/api/user"
	"gm-user/pkg/config"
	"gm-user/pkg/dao"
	login_service "gm-user/pkg/service/login.service.v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
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
	rg.Route(&user.RouterUser{}, r)
}

type grpcConfig struct {
	Addr         string
	RegisterFunc func(*grpc.Server)
}

func RegisterGrpc() *grpc.Server {
	c := grpcConfig{
		Addr: config.GrpcConf().Addr,
		RegisterFunc: func(server *grpc.Server) {
			login_service.RegisterLoginServiceServer(server, &login_service.LoginService{
				Cache: dao.Rdb,
			})
		},
	}

	s := grpc.NewServer()
	c.RegisterFunc(s)

	lis, err := net.Listen("tcp", config.GrpcConf().Addr)
	if err != nil {
		zap.L().Warn("grpc cannot listen", zap.Error(err))
	}

	go func() {
		err = s.Serve(lis)
		if err != nil {
			zap.L().Warn("grpc server start failed", zap.Error(err))
			return
		}
	}()

	return s
}
