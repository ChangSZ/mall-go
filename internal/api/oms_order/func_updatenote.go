package oms_order

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type updateNoteRequest struct{}

type updateNoteResponse struct{}

// UpdateNote 备注订单
// @Summary 备注订单
// @Description 备注订单
// @Tags OmsOrderController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateNoteRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateNoteResponse}
// @Failure 400 {object} code.Failure
// @Router /order/update/note [post]
func (h *handler) UpdateNote(ctx *gin.Context) {
	api.Success(ctx, nil)
}
