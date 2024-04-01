package ums_admin

import (
	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/redis"
	"github.com/ChangSZ/mall-go/internal/services/menu"
	"github.com/ChangSZ/mall-go/pkg/hash"

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

	// ListAll 获取所有角色
	// @Tags UmsAdminController
	// @Router /admin/listAll [get]
	ListAll() core.HandlerFunc

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
	logger      *zap.Logger
	cache       redis.Repo
	hashids     hash.Hash
	menuService menu.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:      logger,
		cache:       cache,
		hashids:     hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		menuService: menu.New(db, cache),
	}
}

func (h *handler) i() {}