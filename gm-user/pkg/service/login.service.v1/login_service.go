package login_service

import (
	"context"
	common "gm-common"
	"gm-user/pkg/dao"
	"gm-user/pkg/model"
	"gm-user/pkg/repo"
	"go.uber.org/zap"
	"time"
)

type LoginService struct {
	UnimplementedLoginServiceServer
	Cache repo.Cache
}

func New() *LoginService {
	return &LoginService{
		Cache: dao.Rdb,
	}
}

func (l *LoginService) GetCaptcha(ctx context.Context, msg *CaptchaMessage) (*CaptchaResp, error) {

	mobile := msg.Mobile
	if !common.VerifyMobile(msg.Mobile) {
		return nil, model.IllegalMobile
	}

	testCode := "123456"
	//todo: 接入短信验证码发送平台，并存入缓存中
	go func() {
		time.Sleep(time.Second * 2)
		zap.L().Info("调用短信平台,发送验证码", zap.String("mobile", mobile))
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		err := l.Cache.Put(ctx, "REGISTER_"+mobile, testCode, time.Minute*15)
		if err != nil {
			zap.L().Warn("验证码存储错误", zap.Error(err))
		}
	}()

	return &CaptchaResp{}, nil
}
