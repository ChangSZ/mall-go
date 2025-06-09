package pms_portal_brand

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 前台品牌管理Service
type Service interface {
	i()

	/**
	 * 分页获取推荐品牌
	 */
	RecommendList(ctx context.Context, pageNum, pageSize int) ([]dto.PmsBrand, error)

	/**
	 * 获取品牌详情
	 */
	Detail(ctx context.Context, brandId int64) (*dto.PmsBrand, error)

	/**
	 * 分页获取品牌关联商品
	 */
	ProductList(ctx context.Context, brandId int64, pageNum int, pageSize int) ([]dto.PmsProduct, int64, error)
}
