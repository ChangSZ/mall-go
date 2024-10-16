package ums_member_coupon

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/services/mall_portal/oms_cart_item"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type listCartRequest struct {
	Type int32 `uri:"type" binding:"required"` // 类型
}

type listCartResponse struct {
	List []dto.SmsCouponHistoryDetail `json:",inline"`
}

// ListCart 获取登录会员购物车的相关优惠券
// @Summary 获取登录会员购物车的相关优惠券
// @Description 获取登录会员购物车的相关优惠券
// @Tags UmsMemberCouponController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listCartRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]dto.SmsCouponHistoryDetail}
// @Failure 400 {object} code.Failure
// @Router /member/coupon/list/cart/{type} [get]
func (h *handler) ListCart(ctx *gin.Context) {
	req := new(listCartRequest)
	res := new(listCartResponse)
	if err := ctx.ShouldBindUri(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	cartItemList, err := oms_cart_item.New().ListPromotion(ctx, nil)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}

	list, err := h.service.ListCart(ctx, cartItemList, req.Type)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = list
	api.Success(ctx, res.List)
}
