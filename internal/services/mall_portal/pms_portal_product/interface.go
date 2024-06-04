package pms_portal_product

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 前台商品管理Service
type Service interface {
	i()

	/**
	 * 综合搜索商品
	 */
	Search(ctx context.Context, keyword string, brandId, productCategoryId int64,
		pageNum, pageSize, sort int) ([]dto.PmsProduct, int64, error)

	/**
	 * 以树形结构获取所有商品分类
	 */
	CategoryTreeList(ctx context.Context) ([]dto.PmsProductCategoryNode, error)

	/**
	 * 获取前台商品详情
	 */
	Detail(ctx context.Context, id int64) (*dto.PmsPortalProductDetail, error)
}
