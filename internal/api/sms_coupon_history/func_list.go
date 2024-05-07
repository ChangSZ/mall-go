package sms_coupon_history

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type listRequest struct{}

type listResponse struct{}

// List 根据优惠券id，使用状态，订单编号分页获取领取记录
// @Summary 根据优惠券id，使用状态，订单编号分页获取领取记录
// @Description 根据优惠券id，使用状态，订单编号分页获取领取记录
// @Tags SmsCouponHistoryController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /couponHistory/list [get]
func (h *handler) List(ctx *gin.Context) {
	api.Success(ctx, nil)
}
