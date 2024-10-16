package alipay

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type queryRequest struct {
	OutTradeNo string `form:"outTradeNo"`
	TradeNo    string `form:"tradeNo"`
}

type queryResponse struct{}

// Query 支付宝统一收单线下交易查询
// @Summary 支付宝统一收单线下交易查询
// @Description 支付宝统一收单线下交易查询
// @Tags AlipayController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body queryRequest true "请求信息"
// @Success 200 {object} code.Success{data=queryResponse}
// @Failure 400 {object} code.Failure
// @Router /alipay/query [get]
func (h *handler) Query(ctx *gin.Context) {
	req := new(queryRequest)
	_ = new(queryResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	status, err := h.service.Query(ctx, req.OutTradeNo, req.TradeNo)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	api.Success(ctx, status)
}
