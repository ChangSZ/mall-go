package ums_menu

import (
	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/services/menu"
	"github.com/ChangSZ/mall-go/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 添加后台菜单
	// @Tags UmsMenuController
	// @Router /menu/create [post]
	Create() core.HandlerFunc

	// Update 修改后台菜单
	// @Tags UmsMenuController
	// @Router /menu/update/{id} [post]
	Update() core.HandlerFunc

	// Get 根据ID获取菜单详情
	// @Tags UmsMenuController
	// @Router /menu/{id} [get]
	Get() core.HandlerFunc

	// Delete 根据ID删除后台菜单
	// @Tags UmsMenuController
	// @Router /menu/delete/{id} [post]
	Delete() core.HandlerFunc

	// List 分页查询后台菜单
	// @Tags UmsMenuController
	// @Router /menu/list/{parentId} [get]
	List() core.HandlerFunc

	// TreeList 树形结构返回所有菜单列表
	// @Tags UmsMenuController
	// @Router /menu/treeList [get]
	TreeList() core.HandlerFunc

	// UpdateHidden 修改菜单显示状态
	// @Tags UmsMenuController
	// @Router /menu/updateHidden/{id} [post]
	UpdateHidden() core.HandlerFunc
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
