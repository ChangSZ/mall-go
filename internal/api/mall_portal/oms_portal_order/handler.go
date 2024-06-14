package oms_portal_order

import (
	"github.com/ChangSZ/mall-go/internal/services/mall_portal/oms_portal_order"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// GenerateConfirmOrder 根据购物车信息生成确认单
	// @Tags OmsPortalOrderController
	// @Router /order/generateConfirmOrder [post]
	GenerateConfirmOrder(*gin.Context)

	// GenerateOrder 根据购物车信息生成订单
	// @Tags OmsPortalOrderController
	// @Router /order/generateOrder [post]
	GenerateOrder(*gin.Context)

	// PaySuccess 用户支付成功的回调
	// @Tags OmsPortalOrderController
	// @Router /order/paySuccess [post]
	PaySuccess(*gin.Context)

	// CancelTimeOutOrder 自动取消超时订单
	// @Tags OmsPortalOrderController
	// @Router /order/cancelTimeOutOrder [post]
	CancelTimeOutOrder(*gin.Context)

	// CancelOrder 取消单个超时订单
	// @Tags OmsPortalOrderController
	// @Router /order/cancelOrder [post]
	CancelOrder(*gin.Context)

	// List 按状态分页获取用户订单列表
	// @Tags OmsPortalOrderController
	// @Router /order/list [get]
	List(*gin.Context)

	// Detail 根据ID获取订单详情
	// @Tags OmsPortalOrderController
	// @Router /order/detail/{orderId} [get]
	Detail(*gin.Context)

	// CancelUserOrder 用户取消订单
	// @Tags OmsPortalOrderController
	// @Router /order/cancelUserOrder [post]
	CancelUserOrder(*gin.Context)

	// ConfirmReceiveOrder 用户确认收货
	// @Tags OmsPortalOrderController
	// @Router /order/confirmReceiveOrder [post]
	ConfirmReceiveOrder(*gin.Context)

	// DeleteOrder 用户删除订单
	// @Tags OmsPortalOrderController
	// @Router /order/deleteOrder [post]
	DeleteOrder(*gin.Context)
}

type handler struct {
	service oms_portal_order.Service
}

func New() Handler {
	return &handler{
		service: oms_portal_order.New(),
	}
}

func (h *handler) i() {}
