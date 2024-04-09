package ums_resource

import (
	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/services/menu"
	"github.com/ChangSZ/mall-go/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 添加后台资源
	// @Tags UmsResourceController
	// @Router /resource/create [post]
	Create() core.HandlerFunc

	// Update 修改后台资源
	// @Tags UmsResourceController
	// @Router /resource/update/{id} [post]
	Update() core.HandlerFunc

	// Get 根据ID获取资源详情
	// @Tags UmsResourceController
	// @Router /resource/{id} [get]
	Get() core.HandlerFunc

	// Delete 根据ID删除后台资源
	// @Tags UmsResourceController
	// @Router /resource/delete/{id} [post]
	Delete() core.HandlerFunc

	// List 分页模糊查询后台资源
	// @Tags UmsResourceController
	// @Router /resource/list [get]
	List() core.HandlerFunc

	// ListAll 查询所有后台资源
	// @Tags UmsResourceController
	// @Router /resource/listAll [get]
	ListAll() core.HandlerFunc
}

type handler struct {
	logger      *zap.Logger
	hashids     hash.Hash
	menuService menu.Service
}

func New(logger *zap.Logger, db mysql.Repo) Handler {
	return &handler{
		logger:      logger,
		hashids:     hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		menuService: menu.New(db),
	}
}

func (h *handler) i() {}
