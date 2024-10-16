package oms_order

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type updateNoteRequest struct {
	Id     int64  `form:"id"`
	Note   string `form:"note"`
	Status int32  `form:"status"`
}

type updateNoteResponse struct {
	Count int64 `json:",inline"`
}

// UpdateNote 备注订单
// @Summary 备注订单
// @Description 备注订单
// @Tags OmsOrderController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateNoteRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /order/update/note [post]
func (h *handler) UpdateNote(ctx *gin.Context) {
	req := new(updateNoteRequest)
	res := new(updateNoteResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	cnt, err := h.service.UpdateNote(ctx, req.Id, req.Note, req.Status)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
