package ums_resource_cate

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

	// ListAll 查询所有后台资源分类
	// @Tags UmsResourceController
	// @Router /resourceCategory/listAll [get]
	ListAll(*gin.Context)

	// Create 添加后台资源分类
	// @Tags UmsResourceController
	// @Router /resourceCategory/create [post]
	Create(*gin.Context)

	// Update 修改后台资源分类
	// @Tags UmsResourceController
	// @Router /resourceCategory/update/{id} [post]
	Update(*gin.Context)

	// Delete 根据ID删除后台资源分类
	// @Tags UmsResourceController
	// @Router /resourceCategory/delete/{id} [post]
	Delete(*gin.Context)
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
