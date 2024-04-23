package ums_menu

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_menu"
)

var _ Service = (*service)(nil)

// 会员等级管理Service
type Service interface {
	i()

	/**
	 * 创建后台菜单
	 */
	Create(ctx context.Context, umsMenu *ums_menu.UmsMenu) (int64, error)

	/**
	 * 修改后台菜单
	 */
	Update(ctx context.Context, id int64, umsMenu *ums_menu.UmsMenu) (int64, error)

	/**
	 * 根据ID获取菜单详情
	 */
	GetItem(ctx context.Context, id int64) (*ums_menu.UmsMenu, error)

	/**
	 * 根据ID删除菜单
	 */
	Delete(ctx context.Context, id int64) (int64, error)

	/**
	 * 分页查询后台菜单
	 */
	List(ctx context.Context, parentId int64, pageSize, pageNum int) ([]*ums_menu.UmsMenu, int64, error)

	/**
	 * 树形结构返回所有菜单列表
	 */
	TreeList(ctx context.Context) ([]UmsMenuNode, error)

	/**
	 * 修改菜单显示状态
	 */
	UpdateHidden(ctx context.Context, id int64, hidden int32) (int64, error)
}
