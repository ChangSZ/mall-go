package ums_resource_cate

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_resource_category"
)

var _ Service = (*service)(nil)

// 后台资源分类管理Service
type Service interface {
	i()

	/**
	 * 获取所有资源分类
	 */
	ListAll(ctx context.Context) ([]*ums_resource_category.UmsResourceCategory, error)

	/**
	 * 创建资源分类
	 */
	Create(ctx context.Context, umsResourceCate *ums_resource_category.UmsResourceCategory) (int64, error)

	/**
	 * 修改资源分类
	 */
	Update(ctx context.Context, id int64, umsResourceCate *ums_resource_category.UmsResourceCategory) (int64, error)

	/**
	 * 删除资源分类
	 */
	Delete(ctx context.Context, id int64) (int64, error)
}
