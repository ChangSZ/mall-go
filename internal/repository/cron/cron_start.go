package cron

import (
	"math"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/cron_task"
	"github.com/ChangSZ/mall-go/pkg/log"
)

func (s *server) Start() {
	s.cron.Start()
	go s.taskCount.Wait()

	qb := cron_task.NewQueryBuilder()
	qb.WhereIsUsed(mysql.EqualPredicate, cron_task.IsUsedYES)
	totalNum, err := qb.Count(mysql.DB().GetDbR())
	if err != nil {
		log.Fatal("cron initialize tasks count err: ", err)
	}

	pageSize := 50
	maxPage := int(math.Ceil(float64(totalNum) / float64(pageSize)))

	taskNum := 0
	log.Info("开始初始化后台任务")

	for page := 1; page <= maxPage; page++ {
		qb = cron_task.NewQueryBuilder()
		qb.WhereIsUsed(mysql.EqualPredicate, cron_task.IsUsedYES)
		listData, err := qb.
			Limit(pageSize).
			Offset((page - 1) * pageSize).
			OrderById(false).
			QueryAll(mysql.DB().GetDbR())
		if err != nil {
			log.Fatal("cron initialize tasks list err: ", err)
		}

		for _, item := range listData {
			s.AddTask(item)
			taskNum++
		}
	}

	log.Infof("后台任务初始化完成, 总数量: %d", taskNum)
}
