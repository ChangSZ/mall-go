package ums_role

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/pagehelper"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	/**
	 * 添加角色
	 */
	Create(ctx context.Context, param dto.UmsRoleParam) (int64, error)

	/**
	 * 修改角色信息
	 */
	Update(ctx context.Context, id int64, param dto.UmsRoleParam) (int64, error)

	/**
	 * 修改角色状态
	 */
	UpdateStatus(ctx context.Context, id int64, status int32) (int64, error)

	/**
	 * 批量删除角色
	 */
	Delete(ctx context.Context, ids []int64) (int64, error)

	/**
	 * 获取所有角色列表
	 */
	ListAll(ctx context.Context) ([]dto.UmsRole, error)

	/**
	 * 分页获取角色列表
	 */
	List(ctx context.Context, keyword string, pageSize, pageNum int) (*pagehelper.ListData[dto.UmsRole], error)

	/**
	 * 根据管理员ID获取对应菜单
	 */
	GetMenuList(ctx context.Context, adminId int64) ([]dto.UmsMenu, error)

	/**
	 * 获取角色相关菜单
	 */
	ListMenu(ctx context.Context, roleId int64) ([]dto.UmsMenu, error)

	/**
	 * 获取角色相关资源
	 */
	ListResource(ctx context.Context, roleId int64) ([]dto.UmsResource, error)

	/**
	 * 给角色分配菜单
	 */
	AllocMenu(ctx context.Context, roleId int64, menuIds string) (int64, error)

	/**
	 * 给角色分配资源
	 */
	AllocResource(ctx context.Context, roleId int64, resourceIds []int64) (int64, error)
}
