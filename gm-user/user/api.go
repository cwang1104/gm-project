package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
}

func (*HandlerUser) GetCap(c *gin.Context) {
	c.JSON(http.StatusOK, "success")
}
