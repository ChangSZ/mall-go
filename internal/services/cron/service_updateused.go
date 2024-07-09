package cron

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/cron_task"
)

func (s *service) UpdateUsed(ctx context.Context, id int64, used int32) (err error) {
	data := map[string]interface{}{
		"is_used":      used,
		"updated_user": core.SessionUserInfo(ctx).UserName,
	}

	qb := cron_task.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	_, err = qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
	if err != nil {
		return err
	}

	// region 操作定时任务 避免主从同步延迟，在这需要查询主库
	if used == cron_task.IsUsedNo {
		s.cronServer.RemoveTask(cast.ToInt(id))
	} else {
		qb = cron_task.NewQueryBuilder()
		qb.WhereId(mysql.EqualPredicate, id)
		info, err := qb.QueryOne(mysql.DB().GetDbW().WithContext(ctx))
		if err != nil {
			return err
		}

		s.cronServer.RemoveTask(cast.ToInt(id))
		s.cronServer.AddTask(info)

	}
	// endregion

	return
}
