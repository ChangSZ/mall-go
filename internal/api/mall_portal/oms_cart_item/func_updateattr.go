package oms_cart_item

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type updateAttrRequest struct {
	dto.OmsCartItem `json:",inline"`
}

type updateAttrResponse struct {
	Count int64 `json:",inline"`
}

// UpdateAttr 修改购物车中商品的规格
// @Summary 修改购物车中商品的规格
// @Description 修改购物车中商品的规格
// @Tags OmsCartItemController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateAttrRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /cart/update/attr [post]
func (h *handler) UpdateAttr(ctx *gin.Context) {
	req := new(updateAttrRequest)
	res := new(updateAttrResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	count, err := h.service.UpdateAttr(ctx, req.OmsCartItem)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = count
	api.Success(ctx, res.Count)
}
