package oms_order

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type updateReceiverInfoRequest struct{}

type updateReceiverInfoResponse struct{}

// UpdateReceiverInfo 修改收货人信息
// @Summary 修改收货人信息
// @Description 修改收货人信息
// @Tags OmsOrderController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateReceiverInfoRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateReceiverInfoResponse}
// @Failure 400 {object} code.Failure
// @Router /order/update/receiverInfo [post]
func (h *handler) UpdateReceiverInfo(ctx *gin.Context) {
	api.Success(ctx, nil)
}
