package pms_portal_brand

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_brand"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product"
	"github.com/ChangSZ/mall-go/pkg/copy"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) RecommendList(ctx context.Context, pageNum, pageSize int) ([]dto.PmsBrand, error) {
	return new(dao.HomeDao).GetRecommendBrandList(ctx, mysql.DB().GetDbR().WithContext(ctx), pageNum, pageSize)
}

func (s *service) Detail(ctx context.Context, brandId int64) (*dto.PmsBrand, error) {
	qb := pms_brand.NewQueryBuilder().WhereId(mysql.EqualPredicate, brandId)
	item, err := qb.First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	res := &dto.PmsBrand{}
	copy.AssignStruct(item, res)
	return res, nil
}

func (s *service) ProductList(ctx context.Context,
	brandId int64, pageNum int, pageSize int) ([]dto.PmsProduct, int64, error) {
	qb := pms_product.NewQueryBuilder().
		WhereDeleteStatus(mysql.EqualPredicate, 0).
		WherePublishStatus(mysql.EqualPredicate, 1).
		WhereBrandId(mysql.EqualPredicate, brandId)

	count, err := qb.Count(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, 0, err
	}
	offset := (pageNum - 1) * pageSize
	list, err := qb.
		Limit(pageSize).
		Offset(offset).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, 0, err
	}

	listData := make([]dto.PmsProduct, 0, len(list))
	for _, v := range list {
		tmp := dto.PmsProduct{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, count, nil
}
