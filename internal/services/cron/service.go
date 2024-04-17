package cron

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/repository/cron"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/cron_task"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Create(ctx context.Context, createData *CreateCronTaskData) (id int64, err error)
	Modify(ctx context.Context, id int64, modifyData *ModifyCronTaskData) (err error)
	PageList(ctx context.Context, searchData *SearchData) (listData []*cron_task.CronTask, err error)
	PageListCount(ctx context.Context, searchData *SearchData) (total int64, err error)
	UpdateUsed(ctx context.Context, id int64, used int32) (err error)
	Execute(ctx context.Context, id int64) (err error)
	Detail(ctx context.Context, searchOneData *SearchOneData) (info *cron_task.CronTask, err error)
}

type service struct {
	cronServer cron.Server
}

func New(cron cron.Server) Service {
	return &service{cronServer: cron}
}

func (s *service) i() {}
