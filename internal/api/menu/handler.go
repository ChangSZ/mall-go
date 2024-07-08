package menu

import (
	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/services/menu"

	"github.com/ChangSZ/golib/hash"
	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 创建/编辑菜单
	// @Tags API.menu
	// @Router /api/menu [post]
	Create(*gin.Context)

	// Detail 菜单详情
	// @Tags API.menu
	// @Router /api/menu/{id} [get]
	Detail(*gin.Context)

	// Delete 删除菜单
	// @Tags API.menu
	// @Router /api/menu/{id} [delete]
	Delete(*gin.Context)

	// UpdateUsed 更新菜单为启用/禁用
	// @Tags API.menu
	// @Router /api/menu/used [patch]
	UpdateUsed(*gin.Context)

	// UpdateSort 更新菜单排序
	// @Tags API.menu
	// @Router /api/menu/sort [patch]
	UpdateSort(*gin.Context)

	// List 菜单列表
	// @Tags API.menu
	// @Router /api/menu [get]
	List(*gin.Context)

	// CreateAction 创建功能权限
	// @Tags API.menu
	// @Router /api/menu_action [post]
	CreateAction(*gin.Context)

	// ListAction 功能权限列表
	// @Tags API.menu
	// @Router /api/menu_action [get]
	ListAction(*gin.Context)

	// DeleteAction 删除功能权限
	// @Tags API.menu
	// @Router /api/menu_action/{id} [delete]
	DeleteAction(*gin.Context)
}

type handler struct {
	hashids hash.Hash
	service menu.Service
}

func New() Handler {
	return &handler{
		hashids: hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		service: menu.New(),
	}
}

func (h *handler) i() {}
