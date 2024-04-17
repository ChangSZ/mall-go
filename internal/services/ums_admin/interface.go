package ums_admin

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_admin"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_resource"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_role"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	/**
	 * 注册功能
	 */
	Register(ctx context.Context, umsAdminParam *UmsAdminParam) (*ums_admin.UmsAdmin, error)

	/**
	 * 登录功能
	 * @param username 用户名
	 * @param password 密码
	 * @return 生成的JWT的token
	 */
	Login(ctx context.Context, username, password string) (string, error)

	/**
	 * 刷新token的功能
	 * @param oldToken 旧的token
	 */
	RefreshToken(ctx context.Context, oldToken string) (string, error)

	/**
	 * 根据用户名或昵称分页查询用户
	 */
	List(ctx context.Context, keyword string, pageSize, pageNum int) ([]ums_admin.UmsAdmin, int64, error)

	/**
	 * 修改指定用户信息
	 */
	Update(ctx context.Context, id int64, admin *ums_admin.UmsAdmin) (int64, error)

	// /**
	// * 删除指定用户
	//  */
	// Delete(ctx context.Context, id int64) (int64, error)

	// /**
	// * 修改用户角色关系
	//  */
	// UpdateRole(ctx context.Context, adminId int64, roleIds []int64) (int64, error)

	/**
	* 获取用户对应角色
	 */
	GetRoleList(ctx context.Context, adminId int64) ([]ums_role.UmsRole, error)

	// /**
	// * 修改密码
	//  */
	// UpdatePassword(ctx context.Context, updatePasswordParam dto.UpdateAdminPasswordParam) (int64, error)

	/**
	* 获取缓存服务
	 */
	// UmsAdminCacheService getCacheService();

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

// UmsAdminCacheService interface for the cache service
type UmsAdminCacheService interface {
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
