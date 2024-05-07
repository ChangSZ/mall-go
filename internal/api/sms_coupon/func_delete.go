package sms_coupon

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type deleteRequest struct{}

type deleteResponse struct{}

// Delete 删除优惠券
// @Summary 删除优惠券
// @Description 删除优惠券
// @Tags SmsCouponController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body deleteRequest true "请求信息"
// @Success 200 {object} code.Success{data=deleteResponse}
// @Failure 400 {object} code.Failure
// @Router /coupon/delete/{id} [post]
func (h *handler) Delete(ctx *gin.Context) {
	api.Success(ctx, nil)
}
