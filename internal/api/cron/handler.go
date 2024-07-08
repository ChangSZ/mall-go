package cron

import (
	"github.com/ChangSZ/mall-go/configs"
	cronRepo "github.com/ChangSZ/mall-go/internal/repository/cron"
	"github.com/ChangSZ/mall-go/internal/services/cron"

	"github.com/ChangSZ/golib/hash"
	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 创建任务
	// @Tags API.cron
	// @Router /api/cron [post]
	Create(*gin.Context)

	// Modify 编辑任务
	// @Tags API.cron
	// @Router /api/cron/{id} [post]
	Modify(*gin.Context)

	// List 任务列表
	// @Tags API.cron
	// @Router /api/cron [get]
	List(*gin.Context)

	// UpdateUsed 更新任务为启用/禁用
	// @Tags API.cron
	// @Router /api/cron/used [patch]
	UpdateUsed(*gin.Context)

	// Detail 获取单条任务详情
	// @Tags API.cron
	// @Router /api/cron/{id} [get]
	Detail(*gin.Context)

	// Execute 手动执行任务
	// @Tags API.cron
	// @Router /api/cron/exec/{id} [patch]
	Execute(*gin.Context)
}

type handler struct {
	hashids hash.Hash
	service cron.Service
}

func New(cronServer cronRepo.Server) Handler {
	return &handler{
		hashids: hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		service: cron.New(cronServer),
	}
}

func (h *handler) i() {}
