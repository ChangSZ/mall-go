package oms_order

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type deleteRequest struct{}

type deleteResponse struct{}

// Delete 批量删除订单
// @Summary 批量删除订单
// @Description 批量删除订单
// @Tags OmsOrderController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body deleteRequest true "请求信息"
// @Success 200 {object} code.Success{data=deleteResponse}
// @Failure 400 {object} code.Failure
// @Router /order/delete [post]
func (h *handler) Delete(ctx *gin.Context) {
	api.Success(ctx, nil)
}
