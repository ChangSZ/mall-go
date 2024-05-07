package sms_coupon

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type listRequest struct{}

type listResponse struct{}

// List 根据优惠券名称和类型分页获取优惠券列表
// @Summary 根据优惠券名称和类型分页获取优惠券列表
// @Description 根据优惠券名称和类型分页获取优惠券列表
// @Tags SmsCouponController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /coupon/list [get]
func (h *handler) List(ctx *gin.Context) {
	api.Success(ctx, nil)
}
