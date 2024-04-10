package ums_member_level

import (
	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/services/menu"
	"github.com/ChangSZ/mall-go/pkg/hash"
	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// List 查询所有会员等级
	// @Tags UmsMemberLevelController
	// @Router /memberLevel/list [get]
	List(*gin.Context)
}

type handler struct {
	logger      *zap.Logger
	hashids     hash.Hash
	menuService menu.Service
}

func New(logger *zap.Logger) Handler {
	return &handler{
		logger:      logger,
		hashids:     hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		menuService: menu.New(),
	}
}

func (h *handler) i() {}
