package oms_cart_item

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type listPromotionRequest struct {
	CartIds []int64 `form:"cartIds"`
}

type listPromotionResponse struct{}

// ListPromotion 获取当前会员的购物车列表,包括促销信息
// @Summary 获取当前会员的购物车列表,包括促销信息
// @Description 获取当前会员的购物车列表,包括促销信息
// @Tags OmsCartItemController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listPromotionRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]dto.CartPromotionItem}
// @Failure 400 {object} code.Failure
// @Router /cart/list/promotion [get]
func (h *handler) ListPromotion(ctx *gin.Context) {
	req := new(listPromotionRequest)
	_ = new(listPromotionResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	list, err := h.service.ListPromotion(ctx, req.CartIds)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	api.Success(ctx, list)
}
