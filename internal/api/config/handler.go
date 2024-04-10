package config

import (
	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Email 修改邮件配置
	// @Tags API.config
	// @Router /api/config/email [patch]
	Email(*gin.Context)
}

type handler struct{}

func New() Handler {
	return &handler{}
}

func (h *handler) i() {}
