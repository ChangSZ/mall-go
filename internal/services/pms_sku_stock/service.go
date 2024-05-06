package pms_sku_stock

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_sku_stock"
	"github.com/ChangSZ/mall-go/pkg/copy"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) ListAll(ctx context.Context, pid int64, keyword string) ([]dto.PmsSkuStock, error) {
	qb := pms_sku_stock.NewQueryBuilder()
	qb = qb.WhereProductId(mysql.EqualPredicate, pid)
	qb = qb.WhereSkuCode(mysql.LikePredicate, "%"+keyword+"%")
	list, err := qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	listData := make([]dto.PmsSkuStock, 0, len(list))
	for _, v := range list {
		tmp := dto.PmsSkuStock{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, nil
}

func (s *service) Update(ctx context.Context, pid int64, param []dto.PmsSkuStockUpdateParam) (int64, error) {
	list := make([]dto.PmsSkuStockUpdateParam, 0)
	for _, v := range param {
		if pid == v.ProductId {
			list = append(list, v)
		}
	}
	return new(dao.PmsSkuStockDao).ReplaceList(ctx, mysql.DB().GetDbW().WithContext(ctx), list)
}
