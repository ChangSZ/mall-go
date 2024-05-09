package oms_order_return_reason

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type updateStatusRequest struct{}

type updateStatusResponse struct{}

// UpdateStatus 修改退货原因启用状态
// @Summary 修改退货原因启用状态
// @Description 修改退货原因启用状态
// @Tags OmsOrderReturnReasonController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateStatusResponse}
// @Failure 400 {object} code.Failure
// @Router /returnReason/update/status [post]
func (h *handler) UpdateStatus(ctx *gin.Context) {
	api.Success(ctx, nil)
}
