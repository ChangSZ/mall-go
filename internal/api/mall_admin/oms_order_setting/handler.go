package oms_order_setting

import (
	"github.com/ChangSZ/mall-go/internal/services/mall_admin/oms_order_setting"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// GetItem 获取指定订单设置
	// @Tags OmsOrderSettingController
	// @Router /orderSetting/{id} [get]
	GetItem(*gin.Context)

	// Update 修改指定订单设置
	// @Tags OmsOrderSettingController
	// @Router /orderSetting/update/{id} [post]
	Update(*gin.Context)
}

type handler struct {
	service oms_order_setting.Service
}

func New() Handler {
	return &handler{
		service: oms_order_setting.New(),
	}
}

func (h *handler) i() {}
