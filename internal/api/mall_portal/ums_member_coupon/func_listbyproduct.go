package ums_member_coupon

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type listByProductRequest struct {
	ProductId int64 `uri:"productId" binding:"required"` // 商品ID
}

type listByProductResponse struct {
	List []dto.SmsCoupon `json:",inline"`
}

// ListByProduct 获取当前商品相关优惠券
// @Summary 获取当前商品相关优惠券
// @Description 获取当前商品相关优惠券
// @Tags UmsMemberCouponController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listByProductRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]dto.SmsCoupon}
// @Failure 400 {object} code.Failure
// @Router /member/coupon/listByProduct/{productId} [get]
func (h *handler) ListByProduct(ctx *gin.Context) {
	req := new(listByProductRequest)
	res := new(listByProductResponse)
	if err := ctx.ShouldBindUri(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	list, err := h.service.ListByProduct(ctx, req.ProductId)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = list
	api.Success(ctx, res.List)
}
