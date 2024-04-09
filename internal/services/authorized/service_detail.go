package authorized

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/authorized"
)

func (s *service) Detail(ctx core.Context, id int32) (info *authorized.Authorized, err error) {
	qb := authorized.NewQueryBuilder()
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)
	qb.WhereId(mysql.EqualPredicate, id)

	info, err = qb.First(mysql.DB().GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
