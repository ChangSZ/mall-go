package ums_resource

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_resource"
)

var _ Service = (*service)(nil)

// 后台资源管理Service
type Service interface {
	i()

	/**
	 * 添加资源
	 */
	Create(ctx context.Context, umsResource *ums_resource.UmsResource) (int64, error)

	/**
	 * 修改资源
	 */
	Update(ctx context.Context, id int64, umsResource *ums_resource.UmsResource) (int64, error)

	/**
	 * 获取资源详情
	 */
	GetItem(ctx context.Context, id int64) (*ums_resource.UmsResource, error)

	/**
	 * 删除资源
	 */
	Delete(ctx context.Context, id int64) (int64, error)

	/**
	 * 分页查询资源
	 */
	List(ctx context.Context, categoryId int64,
		nameKeyword, urlKeyword string, pageSize, pageNum int) ([]*ums_resource.UmsResource, int64, error)

	/**
	 * 查询全部资源
	 */
	ListAll(ctx context.Context) ([]*ums_resource.UmsResource, error)
}
