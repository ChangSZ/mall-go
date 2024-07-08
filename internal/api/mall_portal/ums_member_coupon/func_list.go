package ums_member_coupon

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type listRequest struct {
	UseStatus int32 `form:"useStatus" binding:"omitempty,oneof=0 1 2"` // 优惠券筛选类型:0->未使用；1->已使用；2->已过期
}

type listResponse struct {
	List []dto.SmsCoupon `json:",inline"`
}

// List 获取会员优惠券列表
// @Summary 获取会员优惠券列表
// @Description 获取会员优惠券列表
// @Tags UmsMemberCouponController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]dto.SmsCoupon}
// @Failure 400 {object} code.Failure
// @Router /member/coupon/list [get]
func (h *handler) List(ctx *gin.Context) {
	req := new(listRequest)
	res := new(listResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	list, err := h.service.List(ctx, req.UseStatus)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = list
	api.Success(ctx, res.List)
}
