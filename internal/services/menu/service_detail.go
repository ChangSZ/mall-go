package menu

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/menu"
)

type SearchOneData struct {
	Id     int32 // 用户ID
	IsUsed int32 // 是否启用 1:是  -1:否
}

func (s *service) Detail(ctx context.Context, searchOneData *SearchOneData) (info *menu.Menu, err error) {

	qb := menu.NewQueryBuilder()
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)

	if searchOneData.Id != 0 {
		qb.WhereId(mysql.EqualPredicate, searchOneData.Id)
	}

	if searchOneData.IsUsed != 0 {
		qb.WhereIsUsed(mysql.EqualPredicate, searchOneData.IsUsed)
	}

	info, err = qb.QueryOne(mysql.DB().GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
