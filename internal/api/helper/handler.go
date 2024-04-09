package helper

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/services/authorized"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Md5 加密
	// @Tags Helper
	// @Router /helper/md5/{str} [get]
	Md5() core.HandlerFunc

	// Sign 签名
	// @Tags Helper
	// @Router /helper/sign [post]
	Sign() core.HandlerFunc
}

type handler struct {
	logger            *zap.Logger
	authorizedService authorized.Service
}

func New(logger *zap.Logger) Handler {
	return &handler{
		logger:            logger,
		authorizedService: authorized.New(),
	}
}

func (h *handler) i() {}
