package sms_home_advertise

import (
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/services/mall_admin/sms_home_advertise"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 添加广告
	// @Tags SmsHomeAdvertiseController
	// @Router /home/advertise [post]
	Create(*gin.Context)

	// Update 修改广告
	// @Tags SmsHomeAdvertiseController
	// @Router /home/advertise/update/{id} [post]
	Update(*gin.Context)

	// UpdateStatus 修改上下线状态
	// @Tags SmsHomeAdvertiseController
	// @Router /home/advertise/update/status/{id} [post]
	UpdateStatus(*gin.Context)

	// Delete 删除广告
	// @Tags SmsHomeAdvertiseController
	// @Router /home/advertise/delete [post]
	Delete(*gin.Context)

	// List 分页查询广告
	// @Tags SmsHomeAdvertiseController
	// @Router /home/advertise/list [get]
	List(*gin.Context)

	// GetItem 获取广告详情
	// @Tags SmsHomeAdvertiseController
	// @Router /home/advertise/{id} [get]
	GetItem(*gin.Context)
}

type handler struct {
	service sms_home_advertise.Service
}

func New() Handler {
	return &handler{
		service: sms_home_advertise.New(),
	}
}

func (h *handler) i() {}
