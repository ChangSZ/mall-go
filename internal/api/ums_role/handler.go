package ums_role

import (
	"github.com/ChangSZ/mall-go/internal/services/ums_role"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 添加角色
	// @Tags UmsRoleController
	// @Router /role/create [post]
	Create(*gin.Context)

	// Update 修改角色
	// @Tags UmsRoleController
	// @Router /role/update/{id} [post]
	Update(*gin.Context)

	// Delete 批量删除角色
	// @Tags UmsRoleController
	// @Router /role/delete [post]
	Delete(*gin.Context)

	// ListAll 获取所有角色
	// @Tags UmsRoleController
	// @Router /role/listAll [get]
	ListAll(*gin.Context)

	// List 根据角色名称分页获取角色列表
	// @Tags UmsRoleController
	// @Router /role/list [get]
	List(*gin.Context)

	// UpdateStatus 修改角色状态
	// @Tags UmsRoleController
	// @Router /role/updateStatus/{id} [post]
	UpdateStatus(*gin.Context)

	// ListMenu 获取角色相关菜单
	// @Tags UmsRoleController
	// @Router /role/listMenu/{roleId} [get]
	ListMenu(*gin.Context)

	// ListResource 获取角色相关资源
	// @Tags UmsRoleController
	// @Router /role/listResource/{roleId} [get]
	ListResource(*gin.Context)

	// AllocMenu 给角色分配菜单
	// @Tags UmsRoleController
	// @Router /role/allocMenu [post]
	AllocMenu(*gin.Context)

	// AllocResource 给角色分配资源
	// @Tags UmsRoleController
	// @Router /role/allocResource [post]
	AllocResource(*gin.Context)
}

type handler struct {
	service ums_role.Service
}

func New() Handler {
	return &handler{
		service: ums_role.New(),
	}
}

func (h *handler) i() {}
