package oms_order_return_apply

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type getItemRequest struct{}

type getItemResponse struct {
	dto.OmsOrderReturnApplyResult `json:",inline"`
}

// GetItem 获取退货申请详情
// @Summary 获取退货申请详情
// @Description 获取退货申请详情
// @Tags OmsOrderReturnApplyController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getItemRequest true "请求信息"
// @Success 200 {object} code.Success{data=getItemResponse}
// @Failure 400 {object} code.Failure
// @Router /returnApply/{id} [get]
func (h *handler) GetItem(ctx *gin.Context) {
	_ = new(getItemRequest)
	res := new(getItemResponse)
	uri := new(dto.UriID)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	item, err := h.service.GetItem(ctx, uri.Id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.OmsOrderReturnApplyResult = *item
	api.Success(ctx, res)
}
