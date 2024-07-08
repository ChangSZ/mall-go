package alipay

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type webPayRequest struct {
	dto.AliPayParam `json:",inline"`
}

type webPayResponse struct{}

// WebPay 支付宝手机网站支付
// @Summary 支付宝手机网站支付
// @Description 支付宝手机网站支付
// @Tags AlipayController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body webPayRequest true "请求信息"
// @Success 200 {object} code.Success{data=webPayResponse}
// @Failure 400 {object} code.Failure
// @Router /alipay/webPay [get]
func (h *handler) WebPay(ctx *gin.Context) {
	req := new(webPayRequest)
	_ = new(webPayResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	url, err := h.service.PhoneWebPay(ctx, req.AliPayParam)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}
