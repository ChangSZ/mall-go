package alipay

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type payRequest struct {
	dto.AliPayParam `json:",inline"`
}

type payResponse struct{}

// Pay 支付宝电脑网站支付
// @Summary 支付宝电脑网站支付
// @Description 支付宝电脑网站支付
// @Tags AlipayController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body payRequest true "请求信息"
// @Success 200 {object} code.Success{data=payResponse}
// @Failure 400 {object} code.Failure
// @Router /alipay/pay [get]
func (h *handler) Pay(ctx *gin.Context) {
	req := new(payRequest)
	_ = new(payResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	url, err := h.service.Pay(ctx, req.AliPayParam)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}
