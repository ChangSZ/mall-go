package alipay

import (
	"github.com/ChangSZ/mall-go/internal/services/mall_portal/alipay"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Pay 支付宝电脑网站支付
	// @Tags AlipayController
	// @Router /alipay/pay [get]
	Pay(*gin.Context)

	// WebPay 支付宝手机网站支付
	// @Tags AlipayController
	// @Router /alipay/webPay [get]
	WebPay(*gin.Context)

	// Notify 支付宝异步回调
	// @Tags AlipayController
	// @Router /alipay/notify [post]
	Notify(*gin.Context)

	// Query 支付宝统一收单线下交易查询
	// @Tags AlipayController
	// @Router /alipay/query [get]
	Query(*gin.Context)
}

type handler struct {
	service alipay.Service
}

func New() Handler {
	return &handler{
		service: alipay.New(),
	}
}

func (h *handler) i() {}
