package sms_coupon

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type getItemRequest struct{}

type getItemResponse struct{}

// GetItem 获取单个优惠券的详细信息
// @Summary 获取单个优惠券的详细信息
// @Description 获取单个优惠券的详细信息
// @Tags SmsCouponController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getItemRequest true "请求信息"
// @Success 200 {object} code.Success{data=getItemResponse}
// @Failure 400 {object} code.Failure
// @Router /coupon/{id} [get]
func (h *handler) GetItem(ctx *gin.Context) {
	api.Success(ctx, nil)
}
