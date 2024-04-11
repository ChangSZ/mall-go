package ums_menu

import (
	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/services/menu"
	"github.com/ChangSZ/mall-go/pkg/hash"
	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 添加后台菜单
	// @Tags UmsMenuController
	// @Router /menu/create [post]
	Create(*gin.Context)

	// Update 修改后台菜单
	// @Tags UmsMenuController
	// @Router /menu/update/{id} [post]
	Update(*gin.Context)

	// Get 根据ID获取菜单详情
	// @Tags UmsMenuController
	// @Router /menu/{id} [get]
	Get(*gin.Context)

	// Delete 根据ID删除后台菜单
	// @Tags UmsMenuController
	// @Router /menu/delete/{id} [post]
	Delete(*gin.Context)

	// List 分页查询后台菜单
	// @Tags UmsMenuController
	// @Router /menu/list/{parentId} [get]
	List(*gin.Context)

	// TreeList 树形结构返回所有菜单列表
	// @Tags UmsMenuController
	// @Router /menu/treeList [get]
	TreeList(*gin.Context)

	// UpdateHidden 修改菜单显示状态
	// @Tags UmsMenuController
	// @Router /menu/updateHidden/{id} [post]
	UpdateHidden(*gin.Context)
}

type handler struct {
	hashids     hash.Hash
	menuService menu.Service
}

func New() Handler {
	return &handler{
		hashids:     hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		menuService: menu.New(),
	}
}

func (h *handler) i() {}
