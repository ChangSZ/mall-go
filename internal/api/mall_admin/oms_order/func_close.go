package oms_order

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type closeRequest struct {
	Ids  []int64 `form:"ids"`
	Note string  `form:"note"`
}

type closeResponse struct {
	Count int64 `json:",inline"`
}

// Close 批量关闭订单
// @Summary 批量关闭订单
// @Description 批量关闭订单
// @Tags OmsOrderController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body closeRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /order/update/close [post]
func (h *handler) Close(ctx *gin.Context) {
	req := new(closeRequest)
	res := new(closeResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	cnt, err := h.service.Close(ctx, req.Ids, req.Note)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
