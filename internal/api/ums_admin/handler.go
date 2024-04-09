package ums_admin

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/services/ums_admin"
	"github.com/ChangSZ/mall-go/internal/services/ums_role"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Register 用户注册
	// @Tags UmsAdminController
	// @Router /admin/register [post]
	Register() core.HandlerFunc

	// Login 登录以后返回token
	// @Tags UmsAdminController
	// @Router /admin/login [post]
	Login() core.HandlerFunc

	// RefreshToken 刷新token
	// @Tags UmsAdminController
	// @Router /admin/refreshToken [get]
	RefreshToken() core.HandlerFunc

	// Info 获取当前登录用户信息
	// @Tags UmsAdminController
	// @Router /admin/info [get]
	Info() core.HandlerFunc

	// Logout 登出功能
	// @Tags UmsAdminController
	// @Router /admin/logout [post]
	Logout() core.HandlerFunc

	// Get 获取指定用户信息
	// @Tags UmsAdminController
	// @Router /admin/{id} [get]
	Get() core.HandlerFunc

	// Update 修改指定用户信息
	// @Tags UmsAdminController
	// @Router /admin/update/{id} [post]
	Update() core.HandlerFunc

	// UpdatePassword 修改指定用户密码
	// @Tags UmsAdminController
	// @Router /admin/updatePassword [post]
	UpdatePassword() core.HandlerFunc

	// Delete 删除指定用户信息
	// @Tags UmsAdminController
	// @Router /admin/delete/{id} [post]
	Delete() core.HandlerFunc

	// UpdateStatus 修改帐号状态
	// @Tags UmsAdminController
	// @Router /admin/updateStatus/{id} [post]
	UpdateStatus() core.HandlerFunc

	// UpdateRole 给用户分配角色
	// @Tags UmsAdminController
	// @Router /admin/role/update [post]
	UpdateRole() core.HandlerFunc

	// GetRole 获取指定用户的角色
	// @Tags UmsAdminController
	// @Router /admin/role/{adminId} [get]
	GetRole() core.HandlerFunc
}

type handler struct {
	logger          *zap.Logger
	umsAdminService ums_admin.Service
	umsRoleService  ums_role.Service
}

func New(logger *zap.Logger) Handler {
	return &handler{
		logger:          logger,
		umsAdminService: ums_admin.New(),
		umsRoleService:  ums_role.New(),
	}
}

func (h *handler) i() {}
