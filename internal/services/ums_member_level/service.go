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
	qb := ums_member_level.NewQueryBuilder()
	qb = qb.WhereDefaultStatus(mysql.EqualPredicate, int32(defaultStatus))
	return qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
}
