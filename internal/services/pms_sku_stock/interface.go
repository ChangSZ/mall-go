package pms_sku_stock

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 商品管理Service
type Service interface {
	i()

	Update(ctx context.Context, id int64, param dto.PmsSkuStock) (int64, error)
}
