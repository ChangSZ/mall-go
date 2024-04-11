package cron

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/cron_task"
)

func (s *service) PageListCount(ctx context.Context, searchData *SearchData) (total int64, err error) {
	qb := cron_task.NewQueryBuilder()

	if searchData.Name != "" {
		qb.WhereName(mysql.EqualPredicate, searchData.Name)
	}

	if searchData.Protocol != 0 {
		qb.WhereProtocol(mysql.EqualPredicate, searchData.Protocol)
	}

	if searchData.IsUsed != 0 {
		qb.WhereIsUsed(mysql.EqualPredicate, searchData.IsUsed)
	}

	total, err = qb.Count(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return 0, err
	}

	return
}
