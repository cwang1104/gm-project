package user

import (
	"context"
	common "gm-common"
	"gm-user/pkg/dao"
	"gm-user/pkg/model"
	"gm-user/pkg/repo"
	"go.uber.org/zap"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	cache repo.Cache
}

func New() *HandlerUser {
	return &HandlerUser{
		cache: dao.Rdb,
	}
}

func (h *HandlerUser) GetCaptcha(c *gin.Context) {
	resp := &common.Result{}

	mobile := c.PostForm("mobile")
	if !common.VerifyMobile(mobile) {
		c.JSON(http.StatusOK, resp.Fail(model.IllegalMobile, "手机号不合法"))
		return
	}

	testCode := "123456"
	//todo: 接入短信验证码发送平台，并存入缓存中
	go func() {
		time.Sleep(time.Second * 2)
		zap.L().Info("调用短信平台,发送验证码", zap.String("mobile", mobile))
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		err := h.cache.Put(ctx, "REGISTER_"+mobile, testCode, time.Minute*15)
		if err != nil {
			zap.L().Warn("验证码存储错误", zap.Error(err))
		}
	}()

	c.JSON(http.StatusOK, resp.Success(testCode))
}
