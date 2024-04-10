package authorized

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/authorized"
)

func (s *service) List(ctx context.Context, searchData *SearchData) (listData []*authorized.Authorized, err error) {

	qb := authorized.NewQueryBuilder()
	qb = qb.WhereIsDeleted(mysql.EqualPredicate, -1)

	if searchData.BusinessKey != "" {
		qb.WhereBusinessKey(mysql.EqualPredicate, searchData.BusinessKey)
	}

	if searchData.BusinessSecret != "" {
		qb.WhereBusinessSecret(mysql.EqualPredicate, searchData.BusinessSecret)
	}

	if searchData.BusinessDeveloper != "" {
		qb.WhereBusinessDeveloper(mysql.EqualPredicate, searchData.BusinessDeveloper)
	}

	listData, err = qb.
		OrderById(false).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
