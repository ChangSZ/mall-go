package pms_sku_stock

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 商品SKU库存管理Service
type Service interface {
	i()

	/**
	 * 根据商品id和skuCode关键字模糊搜索
	 */
	ListAll(ctx context.Context, pid int64, keyword string) ([]dto.PmsSkuStock, error)

	/**
	 * 批量更新商品库存信息
	 */
	Update(ctx context.Context, pid int64, param []dto.PmsSkuStockUpdateParam) (int64, error)
}
