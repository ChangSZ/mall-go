package ums_role

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_menu"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	/**
	 * 根据管理员ID获取对应菜单
	 */
	GetMenuList(ctx context.Context, adminId int64) ([]ums_menu.UmsMenu, error)
}

// UmsAdminCacheService interface for the cache service
