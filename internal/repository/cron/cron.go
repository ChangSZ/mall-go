package cron

import (
	"sync"

	"github.com/ChangSZ/mall-go/internal/repository/mysql/cron_task"

	"github.com/jakecoffman/cron"
)

var _ Server = (*server)(nil)

type taskCount struct {
	wg   sync.WaitGroup
	exit chan struct{}
}

func (tc *taskCount) i() {}

func (tc *taskCount) Add() {
	tc.wg.Add(1)
}

func (tc *taskCount) Done() {
	tc.wg.Done()
}

func (tc *taskCount) Exit() {
	tc.wg.Done()
	<-tc.exit
}

func (tc *taskCount) Wait() {
	tc.Add()
	tc.wg.Wait()
	close(tc.exit)
}

type server struct {
	cron      *cron.Cron
	taskCount *taskCount
}

type Server interface {
	i()

	// Start 启动 cron 服务
	Start()

	// Stop 停止 cron 服务
	Stop()

	// AddTask 增加定时任务
	AddTask(task *cron_task.CronTask)

	// RemoveTask 删除定时任务
	RemoveTask(taskId int)

	// AddJob 增加定时任务执行的工作内容
	AddJob(task *cron_task.CronTask) cron.FuncJob
}

func New() (Server, error) {
	return &server{
		cron: cron.New(),
		taskCount: &taskCount{
			wg:   sync.WaitGroup{},
			exit: make(chan struct{}),
		},
	}, nil
}

func (s *server) i() {}
