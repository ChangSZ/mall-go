package ums_resource_cate

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_resource_category"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) ListAll(ctx context.Context) ([]*ums_resource_category.UmsResourceCategory, error) {
	qb := ums_resource_category.NewQueryBuilder()
	qb = qb.OrderBySort(false)
	return qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
}

func (s *service) Create(ctx context.Context,
	umsResourceCate *ums_resource_category.UmsResourceCategory) (int64, error) {
	return umsResourceCate.Create(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) Update(ctx context.Context,
	id int64, umsResourceCate *ums_resource_category.UmsResourceCategory) (int64, error) {
	umsResourceCate.Id = id
	qb := ums_resource_category.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	cnt, err := qb.Update(mysql.DB().GetDbW().WithContext(ctx), umsResourceCate)
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

func (s *service) Delete(ctx context.Context, id int64) (int64, error) {
	qb := ums_resource_category.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	cnt, err := qb.Count(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil || cnt == 0 {
		return 0, err
	}
	err = qb.Delete(mysql.DB().GetDbW().WithContext(ctx))
	if err != nil {
		return 0, err
	}
	return cnt, nil
}
