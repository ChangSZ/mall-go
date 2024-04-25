package ums_role

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_menu"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_resource"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_role"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	/**
	 * 添加角色
	 */
	Create(ctx context.Context, umsRole *ums_role.UmsRole) (int64, error)

	/**
	 * 修改角色信息
	 */
	Update(ctx context.Context, id int64, umsRole *ums_role.UmsRole) (int64, error)

	/**
	 * 批量删除角色
	 */
	Delete(ctx context.Context, ids []int64) (int64, error)

	/**
	 * 获取所有角色列表
	 */
	ListAll(ctx context.Context) ([]*ums_role.UmsRole, error)

	/**
	 * 分页获取角色列表
	 */
	List(ctx context.Context, keyword string, pageSize, pageNum int) ([]*ums_role.UmsRole, int64, error)

	/**
	 * 根据管理员ID获取对应菜单
	 */
	GetMenuList(ctx context.Context, adminId int64) ([]ums_menu.UmsMenu, error)

	/**
	 * 获取角色相关菜单
	 */
	ListMenu(ctx context.Context, roleId int64) ([]ums_menu.UmsMenu, error)

	/**
	 * 获取角色相关资源
	 */
	ListResource(ctx context.Context, roleId int64) ([]ums_resource.UmsResource, error)

	/**
	 * 给角色分配菜单
	 */
	AllocMenu(ctx context.Context, roleId int64, menuIds []int64) (int64, error)

	/**
	 * 给角色分配资源
	 */
	AllocResource(ctx context.Context, roleId int64, resourceIds []int64) (int64, error)
}
