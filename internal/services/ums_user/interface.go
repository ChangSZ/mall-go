package ums_user

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_admin"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_resource"
)

type Service interface {
	i()

	/**
	 * 根据用户名获取后台管理员
	 */
	GetAdminByUsername(ctx context.Context, username string) (*ums_admin.UmsAdmin, error)

	/**
	 * 根据用户id获取用户
	 */
	GetItem(ctx context.Context, id int64) (*ums_admin.UmsAdmin, error)

	/**
	* 获取用户信息
	 */
	LoadUserByUsername(ctx context.Context, username string) (*AdminUserDetails, error)

	/**
	* 获取指定用户的可访问资源
	 */
	GetResourceList(ctx context.Context, adminId int64) ([]*ums_resource.UmsResource, error)
}

type UmsUserCacheService interface {
	DelAdmin(ctx context.Context, adminId int64)
	DelResourceList(ctx context.Context, adminId int64)
	DelResourceListByRole(ctx context.Context, roleId int64)
	DelResourceListByRoleIds(ctx context.Context, roleIds []int64)
	DelResourceListByResource(ctx context.Context, resourceId int64)
	GetAdmin(ctx context.Context, username string) *ums_admin.UmsAdmin
	SetAdmin(ctx context.Context, admin *ums_admin.UmsAdmin)
	GetResourceList(ctx context.Context, adminId int64) []ums_resource.UmsResource
	SetResourceList(ctx context.Context, adminId int64, resourceList []ums_resource.UmsResource)
}
