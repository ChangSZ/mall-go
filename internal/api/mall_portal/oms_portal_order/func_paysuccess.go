package oms_portal_order

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type paySuccessRequest struct {
	OrderId int64 `form:"orderId"`
	PayType int32 `form:"payType"`
}

type paySuccessResponse struct {
	Count int64 `json:",inline"`
}

// PaySuccess 用户支付成功的回调
// @Summary 用户支付成功的回调
// @Description 用户支付成功的回调
// @Tags OmsPortalOrderController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body paySuccessRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /order/paySuccess [post]
func (h *handler) PaySuccess(ctx *gin.Context) {
	req := new(paySuccessRequest)
	res := new(paySuccessResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	cnt, err := h.service.PaySuccess(ctx, req.OrderId, req.PayType)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
