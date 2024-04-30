package pms_sku_stock

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_sku_stock"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Update(ctx context.Context, id int64, param dto.PmsSkuStock) (int64, error) {
	data := map[string]interface{}{
		"product_id":      param.ProductId,
		"sku_code":        param.SkuCode,
		"price":           param.Price,
		"stock":           param.Stock,
		"low_stock":       param.LowStock,
		"pic":             param.Pic,
		"sale":            param.Sale,
		"promotion_price": param.PromotionPrice,
		"lock_stock":      param.LockStock,
		"sp_data":         param.SpData,
	}
	qb := pms_sku_stock.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}
