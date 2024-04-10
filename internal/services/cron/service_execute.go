package cron

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/cron_task"
)

func (s *service) Execute(ctx context.Context, id int32) (err error) {
	qb := cron_task.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	info, err := qb.QueryOne(mysql.DB().GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return err
	}

	info.Spec = "手动执行"
	go s.cronServer.AddJob(info)()

	return nil
}
