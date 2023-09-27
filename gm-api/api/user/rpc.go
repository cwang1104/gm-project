package user

import (
	"gm-user/pkg/service/login.service.v1"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var GrpcUserClient login_service.LoginServiceClient

func InitRpcUserClient() {
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.L().Fatal("did not connect grpc user client", zap.Error(err))
	}

	GrpcUserClient = login_service.NewLoginServiceClient(conn)

}
