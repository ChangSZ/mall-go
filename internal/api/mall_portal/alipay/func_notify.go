package alipay

import (
	"net/http"

	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/gin-gonic/gin"
)

type notifyRequest struct{}

type notifyResponse struct{}

// Notify 支付宝异步回调
// @Summary 支付宝异步回调
// @Description 支付宝异步回调
// @Tags AlipayController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body notifyRequest true "请求信息"
// @Success 200 {object} code.Success{data=notifyResponse}
// @Failure 400 {object} code.Failure
// @Router /alipay/notify [post]
func (h *handler) Notify(ctx *gin.Context) {
	_ = new(notifyRequest)
	_ = new(notifyResponse)
	ctx.Request.ParseForm()

	result, err := h.service.Notify(ctx, ctx.Request.Form)
	if err != nil {
		log.WithTrace(ctx).Error(err)
	}
	ctx.JSON(http.StatusOK, result)
}
