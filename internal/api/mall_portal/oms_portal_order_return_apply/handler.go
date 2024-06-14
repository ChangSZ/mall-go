package oms_portal_order_return_apply

import (
	"github.com/ChangSZ/mall-go/internal/services/mall_portal/oms_portal_order_return_apply"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 申请退货
	// @Tags OmsPortalOrderReturnApplyController
	// @Router /returnApply/create [post]
	Create(*gin.Context)
}

type handler struct {
	service oms_portal_order_return_apply.Service
}

func New() Handler {
	return &handler{
		service: oms_portal_order_return_apply.New(),
	}
}

func (h *handler) i() {}
