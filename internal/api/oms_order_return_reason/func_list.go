package oms_order_return_reason

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type listRequest struct{}

type listResponse struct{}

// List 分页查询退货原因
// @Summary 分页查询退货原因
// @Description 分页查询退货原因
// @Tags OmsOrderReturnReasonController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /returnReason/list [get]
func (h *handler) List(ctx *gin.Context) {
	api.Success(ctx, nil)
}
