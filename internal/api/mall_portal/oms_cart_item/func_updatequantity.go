package oms_cart_item

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type updateQuantityRequest struct {
	Id       int64 `form:"id"`
	Quantity int32 `form:"quantity"`
}

type updateQuantityResponse struct {
	Count int64 `json:",inline"`
}

// UpdateQuantity 修改购物车中指定商品的数量
// @Summary 修改购物车中指定商品的数量
// @Description 修改购物车中指定商品的数量
// @Tags OmsCartItemController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateQuantityRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /cart/update/quantity [get]
func (h *handler) UpdateQuantity(ctx *gin.Context) {
	req := new(updateQuantityRequest)
	res := new(updateQuantityResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	count, err := h.service.UpdateQuantity(ctx, req.Id, req.Quantity)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = count
	api.Success(ctx, res.Count)
}
