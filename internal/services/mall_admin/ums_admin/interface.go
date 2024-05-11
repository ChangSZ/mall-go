package ums_admin

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_admin"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_resource"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	/**
	 * 注册功能
	 */
	Register(ctx context.Context, param dto.UmsAdminParam) (*dto.UmsAdmin, error)

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
	List(ctx context.Context, keyword string, pageSize, pageNum int) ([]dto.UmsAdmin, int64, error)

	/**
	 * 修改指定用户信息
	 */
	Update(ctx context.Context, id int64, admin dto.UmsAdmin) (int64, error)

	/**
	 * 修改指定用户的状态
	 */
	UpdateStatus(ctx context.Context, id int64, status int32) (int64, error)

	/**
	* 删除指定用户
	 */
	Delete(ctx context.Context, id int64) (int64, error)

	/**
	* 修改用户角色关系
	 */
	UpdateRole(ctx context.Context, adminId int64, roleIds []int64) (int64, error)

	/**
	* 获取用户对应角色
	 */
	GetRoleList(ctx context.Context, adminId int64) ([]dto.UmsRole, error)

	/**
	* 修改密码
	 */
	UpdatePassword(ctx context.Context, username, oldPassword, newPassword string) (int64, error)

	/**
	 * 根据用户名获取后台管理员
	 */
	GetAdminByUsername(ctx context.Context, username string) (*ums_admin.UmsAdmin, error)

	/**
	 * 根据用户id获取用户
	 */
	GetItem(ctx context.Context, id int64) (*dto.UmsAdmin, error)

	/**
	 * 获取用户信息
	 */
	LoadUserByUsername(ctx context.Context, username string) (*AdminUserDetails, error)

	/**
	 * 获取指定用户的可访问资源
	 */
	GetResourceList(ctx context.Context, adminId int64) ([]*ums_resource.UmsResource, error)
}

// 后台用户缓存管理Service
type UmsAdminCacheServiceI interface {
	/**
	 * 删除后台用户缓存
	 */
	DelAdmin(ctx context.Context, adminId int64)

	/**
	 * 删除后台用户资源列表缓存
	 */
	DelResourceList(ctx context.Context, adminId int64)

	/**
	 * 当角色相关资源信息改变时删除相关后台用户缓存
	 */
	DelResourceListByRole(ctx context.Context, roleId int64)

	/**
	 * 当角色相关资源信息改变时删除相关后台用户缓存
	 */
	DelResourceListByRoleIds(ctx context.Context, roleIds []int64)

	/**
	 * 当资源信息改变时，删除资源项目后台用户缓存
	 */
	DelResourceListByResource(ctx context.Context, resourceId int64)

	/**
	 * 获取缓存后台用户信息
	 */
	GetAdmin(ctx context.Context, username string) *ums_admin.UmsAdmin

	/**
	 * 设置缓存后台用户信息
	 */
	SetAdmin(ctx context.Context, admin *ums_admin.UmsAdmin)

	/**
	 * 获取缓存后台用户资源列表
	 */
	GetResourceList(ctx context.Context, adminId int64) []ums_resource.UmsResource

	/**
	 * 设置缓存后台用户资源列表
	 */
	SetResourceList(ctx context.Context, adminId int64, resourceList []ums_resource.UmsResource)
}
