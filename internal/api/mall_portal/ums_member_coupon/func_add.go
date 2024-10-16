package ums_member_coupon

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type addRequest struct {
	CouponId int64 `uri:"couponId" binding:"required"` // 优惠券ID
}

type addResponse struct{}

// Add 领取指定优惠券
// @Summary 领取指定优惠券
// @Description 领取指定优惠券
// @Tags UmsMemberCouponController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body addRequest true "请求信息"
// @Success 200 {object} code.Success{data=string}
// @Failure 400 {object} code.Failure
// @Router /member/coupon/add/{couponId} [post]
func (h *handler) Add(ctx *gin.Context) {
	req := new(addRequest)
	_ = new(addResponse)
	if err := ctx.ShouldBindUri(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	err := h.service.Add(ctx, req.CouponId)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	api.Success(ctx, "领取成功")
}
