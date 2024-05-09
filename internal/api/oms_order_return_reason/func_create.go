package oms_order_return_reason

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type createRequest struct{}

type createResponse struct{}

// Create 添加退货原因
// @Summary 添加退货原因
// @Description 添加退货原因
// @Tags OmsOrderReturnReasonController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} code.Success{data=createResponse}
// @Failure 400 {object} code.Failure
// @Router /returnReason/create [post]
func (h *handler) Create(ctx *gin.Context) {
	api.Success(ctx, nil)
}
