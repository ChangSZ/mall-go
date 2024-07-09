package ums_member_coupon

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type listHistoryRequest struct {
	UseStatus int32 `form:"useStatus" binding:"omitempty,oneof=0 1 2"` // 优惠券筛选类型:0->未使用；1->已使用；2->已过期
}

type listHistoryResponse struct {
	List []dto.SmsCouponHistory `json:",inline"`
}

// ListHistory 获取会员优惠券历史列表
// @Summary 获取会员优惠券历史列表
// @Description 获取会员优惠券历史列表
// @Tags UmsMemberCouponController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listHistoryRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]dto.SmsCouponHistory}
// @Failure 400 {object} code.Failure
// @Router /member/coupon/listHistory [get]
func (h *handler) ListHistory(ctx *gin.Context) {
	req := new(listHistoryRequest)
	res := new(listHistoryResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	list, err := h.service.ListHistory(ctx, req.UseStatus)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = list
	api.Success(ctx, res.List)
}
