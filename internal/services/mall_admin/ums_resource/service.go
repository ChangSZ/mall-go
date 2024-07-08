package ums_resource

import (
	"context"
	"strings"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_resource"
	"github.com/ChangSZ/mall-go/internal/services/mall_admin/ums_admin"

	"github.com/ChangSZ/golib/copy"
)

type service struct {
	adminCacheService *ums_admin.UmsAdminCacheService
}

func New() Service {
	return &service{ums_admin.NewCacheService()}
}

func (s *service) i() {}

func (s *service) Create(ctx context.Context, param dto.UmsResourceParam) (int64, error) {
	data := &ums_resource.UmsResource{}
	copy.AssignStruct(&param, data)
	return data.Create(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) Update(ctx context.Context, id int64, param dto.UmsResourceParam) (int64, error) {
	data := map[string]interface{}{
		"name":        param.Name,
		"url":         param.Url,
		"description": param.Description,
		"category_id": param.CategoryId,
	}
	qb := ums_resource.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	cnt, err := qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
	if err != nil {
		return 0, err
	}
	s.adminCacheService.DelResourceListByResource(ctx, id)
	return cnt, nil
}

func (s *service) GetItem(ctx context.Context, id int64) (*dto.UmsResource, error) {
	qb := ums_resource.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	data, err := qb.First(mysql.DB().GetDbW().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	res := &dto.UmsResource{}
	copy.AssignStruct(data, res)
	return res, nil
}

func (s *service) Delete(ctx context.Context, id int64) (int64, error) {
	qb := ums_resource.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	cnt, err := qb.Delete(mysql.DB().GetDbW().WithContext(ctx))
	if err != nil {
		return 0, err
	}
	s.adminCacheService.DelResourceListByResource(ctx, id)
	return cnt, nil
}

func (s *service) List(ctx context.Context, categoryId int64,
	nameKeyword, urlKeyword string, pageSize, pageNum int) ([]dto.UmsResource, int64, error) {
	qb := ums_resource.NewQueryBuilder()
	if categoryId != 0 {
		qb = qb.WhereCategoryId(mysql.EqualPredicate, categoryId)
	}
	if strings.TrimSpace(nameKeyword) != "" {
		qb = qb.WhereName(mysql.LikePredicate, "%"+nameKeyword+"%")
	}
	if strings.TrimSpace(urlKeyword) != "" {
		qb = qb.WhereUrl(mysql.LikePredicate, "%"+urlKeyword+"%")
	}
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

	listData := make([]dto.UmsResource, 0, len(list))
	for _, v := range list {
		tmp := dto.UmsResource{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, count, nil
}

func (s *service) ListAll(ctx context.Context) ([]dto.UmsResource, error) {
	qb := ums_resource.NewQueryBuilder()
	list, err := qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	listData := make([]dto.UmsResource, 0, len(list))
	for _, v := range list {
		tmp := dto.UmsResource{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, nil
}
