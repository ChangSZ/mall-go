package ums_role

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

	// Create 添加角色
	// @Tags UmsRoleController
	// @Router /role/create [post]
	Create() core.HandlerFunc

	// Update 修改角色
	// @Tags UmsRoleController
	// @Router /role/update/{id} [post]
	Update() core.HandlerFunc

	// Delete 批量删除角色
	// @Tags UmsRoleController
	// @Router /role/delete [post]
	Delete() core.HandlerFunc

	// ListAll 获取所有角色
	// @Tags UmsRoleController
	// @Router /role/listAll [get]
	ListAll() core.HandlerFunc

	// List 根据角色名称分页获取角色列表
	// @Tags UmsRoleController
	// @Router /role/list [get]
	List() core.HandlerFunc

	// UpdateStatus 修改角色状态
	// @Tags UmsRoleController
	// @Router /role/updateStatus/{id} [post]
	UpdateStatus() core.HandlerFunc

	// ListMenu 获取角色相关菜单
	// @Tags UmsRoleController
	// @Router /role/listMenu/{roleId} [get]
	ListMenu() core.HandlerFunc

	// ListResource 获取角色相关资源
	// @Tags UmsRoleController
	// @Router /role/listResource/{roleId} [get]
	ListResource() core.HandlerFunc

	// AllocMenu 给角色分配菜单
	// @Tags UmsRoleController
	// @Router /role/allocMenu [post]
	AllocMenu() core.HandlerFunc

	// AllocResource 给角色分配资源
	// @Tags UmsRoleController
	// @Router /role/allocResource [post]
	AllocResource() core.HandlerFunc
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
