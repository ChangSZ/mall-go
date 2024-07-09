package helper

import (
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/services/authorized"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Md5 加密
	// @Tags Helper
	// @Router /helper/md5/{str} [get]
	Md5(*gin.Context)

	// Sign 签名
	// @Tags Helper
	// @Router /helper/sign [post]
	Sign(*gin.Context)
}

type handler struct {
	service authorized.Service
}

func New() Handler {
	return &handler{
		service: authorized.New(),
	}
}

func (h *handler) i() {}
