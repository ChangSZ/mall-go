package oms_cart_item

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
)

type clearRequest struct{}

type clearResponse struct {
	Count int64 `json:",inline"`
}

// Clear 清空当前会员的购物车
// @Summary 清空当前会员的购物车
// @Description 清空当前会员的购物车
// @Tags OmsCartItemController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body clearRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /cart/clear [post]
func (h *handler) Clear(ctx *gin.Context) {
	_ = new(clearRequest)
	res := new(clearResponse)

	count, err := h.service.Clear(ctx)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = count
	api.Success(ctx, res.Count)
}
