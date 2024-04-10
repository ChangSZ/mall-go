package admin

import (
	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/services/admin"
	"github.com/ChangSZ/mall-go/pkg/hash"
	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Login 管理员登录
	// @Tags API.admin
	// @Router /api/login [post]
	Login(*gin.Context)

	// Logout 管理员登出
	// @Tags API.admin
	// @Router /api/admin/logout [post]
	Logout(*gin.Context)

	// ModifyPassword 修改密码
	// @Tags API.admin
	// @Router /api/admin/modify_password [patch]
	ModifyPassword(*gin.Context)

	// Detail 个人信息
	// @Tags API.admin
	// @Router /api/admin/info [get]
	Detail(*gin.Context)

	// ModifyPersonalInfo 修改个人信息
	// @Tags API.admin
	// @Router /api/admin/modify_personal_info [patch]
	ModifyPersonalInfo(*gin.Context)

	// Create 新增管理员
	// @Tags API.admin
	// @Router /api/admin [post]
	Create(*gin.Context)

	// List 管理员列表
	// @Tags API.admin
	// @Router /api/admin [get]
	List(*gin.Context)

	// Delete 删除管理员
	// @Tags API.admin
	// @Router /api/admin/{id} [delete]
	Delete(*gin.Context)

	// Offline 下线管理员
	// @Tags API.admin
	// @Router /api/admin/offline [patch]
	Offline(*gin.Context)

	// UpdateUsed 更新管理员为启用/禁用
	// @Tags API.admin
	// @Router /api/admin/used [patch]
	UpdateUsed(*gin.Context)

	// ResetPassword 重置密码
	// @Tags API.admin
	// @Router /api/admin/reset_password/{id} [patch]
	ResetPassword(*gin.Context)

	// CreateAdminMenu 提交菜单授权
	// @Tags API.admin
	// @Router /api/admin/menu [post]
	CreateAdminMenu(*gin.Context)

	// ListAdminMenu 菜单授权列表
	// @Tags API.admin
	// @Router /api/admin/menu/{id} [get]
	ListAdminMenu(*gin.Context)
}

type handler struct {
	hashids      hash.Hash
	adminService admin.Service
}

func New() Handler {
	return &handler{
		hashids:      hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		adminService: admin.New(),
	}
}

func (h *handler) i() {}
