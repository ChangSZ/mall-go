package oms_order_return_reason

import (
	"github.com/ChangSZ/mall-go/internal/services/oms_order_return_reason"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 添加退货原因
	// @Tags OmsOrderReturnReasonController
	// @Router /returnReason/create [post]
	Create(*gin.Context)

	// Update 修改退货原因
	// @Tags OmsOrderReturnReasonController
	// @Router /returnReason/update/{id} [post]
	Update(*gin.Context)

	// List 分页查询退货原因
	// @Tags OmsOrderReturnReasonController
	// @Router /returnReason/list [get]
	List(*gin.Context)

	// Delete 批量删除退货原因
	// @Tags OmsOrderReturnReasonController
	// @Router /returnReason/delete [post]
	Delete(*gin.Context)

	// GetItem 获取单个退货原因详情信息
	// @Tags OmsOrderReturnReasonController
	// @Router /returnReason/{id} [get]
	GetItem(*gin.Context)

	// UpdateStatus 修改退货原因启用状态
	// @Tags OmsOrderReturnReasonController
	// @Router /returnReason/update/status [post]
	UpdateStatus(*gin.Context)
}

type handler struct {
	service oms_order_return_reason.Service
}

func New() Handler {
	return &handler{
		service: oms_order_return_reason.New(),
	}
}

func (h *handler) i() {}
