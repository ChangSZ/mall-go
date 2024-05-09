package oms_order

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type closeRequest struct{}

type closeResponse struct{}

// Close 批量关闭订单
// @Summary 批量关闭订单
// @Description 批量关闭订单
// @Tags OmsOrderController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body closeRequest true "请求信息"
// @Success 200 {object} code.Success{data=closeResponse}
// @Failure 400 {object} code.Failure
// @Router /order/update/close [post]
func (h *handler) Close(ctx *gin.Context) {
	api.Success(ctx, nil)
}
