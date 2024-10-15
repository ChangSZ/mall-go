package oms_portal_order_return_apply

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type createRequest struct {
	dto.OmsOrderReturnApplyParam `json:",inline"`
}

type createResponse struct {
	Count int64 `json:",inline"`
}

// Create 申请退货
// @Summary 申请退货
// @Description 申请退货
// @Tags OmsPortalOrderReturnApplyController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /returnApply/create [post]
func (h *handler) Create(ctx *gin.Context) {
	req := new(createRequest)
	res := new(createResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	count, err := h.service.Create(ctx, req.OmsOrderReturnApplyParam)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	if count == 0 {
		api.Failed(ctx, "操作失败")
		return
	}
	res.Count = count
	api.Success(ctx, res.Count)
}
