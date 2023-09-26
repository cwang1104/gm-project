package login_service

import "context"

type LoginService struct {
	UnimplementedLoginServiceServer
}

func (*LoginService) GetCaptcha(context.Context, *CaptchaMessage) (*CaptchaResp, error) {
	return nil, nil
}
