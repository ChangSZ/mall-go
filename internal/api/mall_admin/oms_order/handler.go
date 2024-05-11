package oms_order

import (
	"github.com/ChangSZ/mall-go/internal/services/mall_admin/oms_order"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// List 查询订单
	// @Tags OmsOrderController
	// @Router /order/list [get]
	List(*gin.Context)

	// Delivery 批量发货
	// @Tags OmsOrderController
	// @Router /order/update/delivery [post]
	Delivery(*gin.Context)

	// Close 批量关闭订单
	// @Tags OmsOrderController
	// @Router /order/update/close [post]
	Close(*gin.Context)

	// Delete 批量删除订单
	// @Tags OmsOrderController
	// @Router /order/delete [post]
	Delete(*gin.Context)

	// GetItem 获取订单详情：订单信息、商品信息、操作记录
	// @Tags OmsOrderController
	// @Router /order/{id} [get]
	GetItem(*gin.Context)

	// UpdateReceiverInfo 修改收货人信息
	// @Tags OmsOrderController
	// @Router /order/update/receiverInfo [post]
	UpdateReceiverInfo(*gin.Context)

	// UpdateMoneyInfo 修改订单费用信息
	// @Tags OmsOrderController
	// @Router /order/update/moneyInfo [post]
	UpdateMoneyInfo(*gin.Context)

	// UpdateNote 备注订单
	// @Tags OmsOrderController
	// @Router /order/update/note [post]
	UpdateNote(*gin.Context)
}

type handler struct {
	service oms_order.Service
}

func New() Handler {
	return &handler{
		service: oms_order.New(),
	}
}

func (h *handler) i() {}
