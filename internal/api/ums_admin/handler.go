package ums_admin

import (
	"github.com/ChangSZ/mall-go/internal/services/ums_admin"
	"github.com/ChangSZ/mall-go/internal/services/ums_role"
	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Register 用户注册
	// @Tags UmsAdminController
	// @Router /admin/register [post]
	Register(*gin.Context)

	// Login 登录以后返回token
	// @Tags UmsAdminController
	// @Router /admin/login [post]
	Login(*gin.Context)

	// RefreshToken 刷新token
	// @Tags UmsAdminController
	// @Router /admin/refreshToken [get]
	RefreshToken(*gin.Context)

	// Info 获取当前登录用户信息
	// @Tags UmsAdminController
	// @Router /admin/info [get]
	Info(*gin.Context)

	// Logout 登出功能
	// @Tags UmsAdminController
	// @Router /admin/logout [post]
	Logout(*gin.Context)

	// Get 获取指定用户信息
	// @Tags UmsAdminController
	// @Router /admin/{id} [get]
	Get(*gin.Context)

	// Update 修改指定用户信息
	// @Tags UmsAdminController
	// @Router /admin/update/{id} [post]
	Update(*gin.Context)

	// UpdatePassword 修改指定用户密码
	// @Tags UmsAdminController
	// @Router /admin/updatePassword [post]
	UpdatePassword(*gin.Context)

	// Delete 删除指定用户信息
	// @Tags UmsAdminController
	// @Router /admin/delete/{id} [post]
	Delete(*gin.Context)

	// UpdateStatus 修改帐号状态
	// @Tags UmsAdminController
	// @Router /admin/updateStatus/{id} [post]
	UpdateStatus(*gin.Context)

	// UpdateRole 给用户分配角色
	// @Tags UmsAdminController
	// @Router /admin/role/update [post]
	UpdateRole(*gin.Context)

	// GetRole 获取指定用户的角色
	// @Tags UmsAdminController
	// @Router /admin/role/{adminId} [get]
	GetRole(*gin.Context)
}

type handler struct {
	umsAdminService ums_admin.Service
	umsRoleService  ums_role.Service
}

func New() Handler {
	return &handler{
		umsAdminService: ums_admin.New(),
		umsRoleService:  ums_role.New(),
	}
}

func (h *handler) i() {}
