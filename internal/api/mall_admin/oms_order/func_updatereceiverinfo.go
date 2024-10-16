package oms_order

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type updateReceiverInfoRequest struct {
	dto.OmsReceiverInfoParam `json:",inline"`
}

type updateReceiverInfoResponse struct {
	Count int64 `json:",inline"`
}

// UpdateReceiverInfo 修改收货人信息
// @Summary 修改收货人信息
// @Description 修改收货人信息
// @Tags OmsOrderController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateReceiverInfoRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /order/update/receiverInfo [post]
func (h *handler) UpdateReceiverInfo(ctx *gin.Context) {
	req := new(updateReceiverInfoRequest)
	res := new(updateReceiverInfoResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	cnt, err := h.service.UpdateReceiverInfo(ctx, req.OmsReceiverInfoParam)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
