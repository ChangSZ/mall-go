package oms_order_return_apply

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type updateStatusRequest struct {
	dto.OmsUpdateStatusParam `json:",inline"`
}

type updateStatusResponse struct {
	Count int64 `json:",inline"`
}

// UpdateStatus 修改退货申请状态
// @Summary 修改退货申请状态
// @Description 修改退货申请状态
// @Tags OmsOrderReturnApplyController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /returnApply/update/status/{id} [post]
func (h *handler) UpdateStatus(ctx *gin.Context) {
	req := new(updateStatusRequest)
	res := new(updateStatusResponse)
	uri := new(dto.UriID)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	cnt, err := h.service.UpdateStatus(ctx, uri.Id, req.OmsUpdateStatusParam)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	if cnt == 0 {
		api.Failed(ctx, "更新Status个数为0")
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
