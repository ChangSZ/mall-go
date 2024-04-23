package ums_member_level

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_member_level"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) List(ctx context.Context, defaultStatus int32) ([]*ums_member_level.UmsMemberLevel, error) {
	queryBuilder := ums_member_level.NewQueryBuilder()
	queryBuilder = queryBuilder.WhereDefaultStatus(mysql.EqualPredicate, int32(defaultStatus))
	return queryBuilder.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
}
