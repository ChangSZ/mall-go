package ums_role

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_menu"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	/**
	 * 根据管理员ID获取对应菜单
	 */
	GetMenuList(ctx core.Context, adminId int64) ([]ums_menu.UmsMenu, error)
}

// UmsAdminCacheService interface for the cache service
