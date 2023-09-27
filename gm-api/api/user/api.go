package user

import (
	"context"
	common "gm-common"
	"gm-user/pkg/dao"
	"gm-user/pkg/model"
	"gm-user/pkg/repo"
	"gm-user/pkg/service/login.service.v1"
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

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := GrpcUserClient.GetCaptcha(ctx, &login_service.CaptchaMessage{
		Mobile: mobile,
	})

	if err != nil {
		c.JSON(http.StatusOK, resp.Fail(2001, err.Error()))
		return
	}

	c.JSON(http.StatusOK, resp.Success(nil))
}
