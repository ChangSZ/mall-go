package ums_resource_cate

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_resource_category"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) ListAll(ctx context.Context) ([]dto.UmsResourceCate, error) {
	qb := ums_resource_category.NewQueryBuilder()
	qb = qb.OrderBySort(false)
	list, err := qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}

	listData := make([]dto.UmsResourceCate, 0, len(list))
	for _, v := range list {
		listData = append(listData, dto.UmsResourceCate{
			Id:         v.Id,
			CreateTime: v.CreateTime,
			Name:       v.Name,
			Sort:       v.Sort,
		})
	}
	return listData, nil
}

func (s *service) Create(ctx context.Context, param dto.UmsResourceCateParam) (int64, error) {
	data := &ums_resource_category.UmsResourceCategory{
		Name: param.Name,
		Sort: param.Sort,
	}
	return data.Create(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) Update(ctx context.Context, id int64, param dto.UmsResourceCateParam) (int64, error) {
	data := map[string]interface{}{
		"name": param.Name,
		"sort": param.Sort,
	}
	qb := ums_resource_category.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) Delete(ctx context.Context, id int64) (int64, error) {
	qb := ums_resource_category.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Delete(mysql.DB().GetDbW().WithContext(ctx))
}
