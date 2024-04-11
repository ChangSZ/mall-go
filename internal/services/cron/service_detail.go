package cron

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/cron_task"
)

type SearchOneData struct {
	Id int32 // 任务ID
}

func (s *service) Detail(ctx context.Context, searchOneData *SearchOneData) (info *cron_task.CronTask, err error) {
	qb := cron_task.NewQueryBuilder()

	if searchOneData.Id != 0 {
		qb.WhereId(mysql.EqualPredicate, searchOneData.Id)
	}

	info, err = qb.QueryOne(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}

	return
}
