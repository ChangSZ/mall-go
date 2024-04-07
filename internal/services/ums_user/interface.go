package ums_user

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_admin"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_resource"
)

type Service interface {
	i()

	/**
	 * 根据用户名获取后台管理员
	 */
	GetAdminByUsername(ctx core.Context, username string) (*ums_admin.UmsAdmin, error)

	/**
	 * 根据用户id获取用户
	 */
	GetItem(ctx core.Context, id int64) (*ums_admin.UmsAdmin, error)

	/**
	* 获取用户信息
	 */
	LoadUserByUsername(ctx core.Context, username string) (*AdminUserDetails, error)

	/**
	* 获取指定用户的可访问资源
	 */
	GetResourceList(ctx core.Context, adminId int64) ([]*ums_resource.UmsResource, error)
}

type UmsUserCacheService interface {
	DelAdmin(ctx core.Context, adminId int64)
	DelResourceList(ctx core.Context, adminId int64)
	DelResourceListByRole(ctx core.Context, roleId int64)
	DelResourceListByRoleIds(ctx core.Context, roleIds []int64)
	DelResourceListByResource(ctx core.Context, resourceId int64)
	GetAdmin(ctx core.Context, username string) *ums_admin.UmsAdmin
	SetAdmin(ctx core.Context, admin *ums_admin.UmsAdmin)
	GetResourceList(ctx core.Context, adminId int64) []ums_resource.UmsResource
	SetResourceList(ctx core.Context, adminId int64, resourceList []ums_resource.UmsResource)
}
