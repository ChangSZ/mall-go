package ums_resource

import (
	"context"
	"strings"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_resource"
	"github.com/ChangSZ/mall-go/internal/services/ums_admin"
)

type service struct {
	adminCacheService *ums_admin.UmsAdminCacheService
}

func New() Service {
	return &service{ums_admin.NewCacheService()}
}

func (s *service) i() {}

func (s *service) Create(ctx context.Context, umsResource *ums_resource.UmsResource) (int64, error) {
	return umsResource.Create(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) Update(ctx context.Context, id int64, umsResource *ums_resource.UmsResource) (int64, error) {
	umsResource.Id = id
	qb := ums_resource.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	cnt, err := qb.Update(mysql.DB().GetDbW().WithContext(ctx), umsResource)
	if err != nil {
		return 0, err
	}
	s.adminCacheService.DelResourceListByResource(ctx, id)
	return cnt, nil
}

func (s *service) GetItem(ctx context.Context, id int64) (*ums_resource.UmsResource, error) {
	qb := ums_resource.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.First(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) Delete(ctx context.Context, id int64) (int64, error) {
	qb := ums_resource.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	cnt, err := qb.Count(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil || cnt == 0 {
		return 0, err
	}
	err = qb.Delete(mysql.DB().GetDbW().WithContext(ctx))
	if err != nil {
		return 0, err
	}
	s.adminCacheService.DelResourceListByResource(ctx, id)
	return cnt, nil
}

func (s *service) List(ctx context.Context, categoryId int64,
	nameKeyword, urlKeyword string, pageSize, pageNum int) ([]*ums_resource.UmsResource, int64, error) {
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
	return list, count, err
}

func (s *service) ListAll(ctx context.Context) ([]*ums_resource.UmsResource, error) {
	qb := ums_resource.NewQueryBuilder()
	return qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
}
