package oms_order

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type deliveryRequest struct{}

type deliveryResponse struct{}

// Delivery 批量发货
// @Summary 批量发货
// @Description 批量发货
// @Tags OmsOrderController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body deliveryRequest true "请求信息"
// @Success 200 {object} code.Success{data=deliveryResponse}
// @Failure 400 {object} code.Failure
// @Router /order/update/delivery [post]
func (h *handler) Delivery(ctx *gin.Context) {
	api.Success(ctx, nil)
}
