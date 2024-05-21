package oms_order_return_reason

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type deleteRequest struct {
	Ids []int64 `form:"ids"`
}

type deleteResponse struct {
	Count int64 `json:",inline"`
}

// Delete 批量删除退货原因
// @Summary 批量删除退货原因
// @Description 批量删除退货原因
// @Tags OmsOrderReturnReasonController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body deleteRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /returnReason/delete [post]
func (h *handler) Delete(ctx *gin.Context) {
	req := new(deleteRequest)
	res := new(deleteResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	cnt, err := h.service.Delete(ctx, req.Ids)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
