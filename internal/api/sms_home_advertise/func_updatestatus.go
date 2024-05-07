package sms_home_advertise

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type updateStatusRequest struct{}

type updateStatusResponse struct{}

// UpdateStatus 修改上下线状态
// @Summary 修改上下线状态
// @Description 修改上下线状态
// @Tags SmsHomeAdvertiseController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateStatusResponse}
// @Failure 400 {object} code.Failure
// @Router /home/advertise/update/status/{id} [post]
func (h *handler) UpdateStatus(ctx *gin.Context) {
	api.Success(ctx, nil)
}
