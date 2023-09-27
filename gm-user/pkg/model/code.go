package model

import (
	"gm-common/errs"
)

var (
	IllegalMobile = errs.NewError(2001, "手机号不合法")
)
