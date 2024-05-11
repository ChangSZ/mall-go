package oms_order_return_apply

import (
	"github.com/ChangSZ/mall-go/internal/services/mall_admin/oms_order_return_apply"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// List 分页查询退货申请
	// @Tags OmsOrderReturnApplyController
	// @Router /returnApply/list [get]
	List(*gin.Context)

	// Delete 批量删除退货申请
	// @Tags OmsOrderReturnApplyController
	// @Router /returnApply/delete [post]
	Delete(*gin.Context)

	// GetItem 获取退货申请详情
	// @Tags OmsOrderReturnApplyController
	// @Router /returnApply/{id} [get]
	GetItem(*gin.Context)

	// UpdateStatus 修改退货申请状态
	// @Tags OmsOrderReturnApplyController
	// @Router /returnApply/update/status/{id} [post]
	UpdateStatus(*gin.Context)
}

type handler struct {
	service oms_order_return_apply.Service
}

func New() Handler {
	return &handler{
		service: oms_order_return_apply.New(),
	}
}

func (h *handler) i() {}
