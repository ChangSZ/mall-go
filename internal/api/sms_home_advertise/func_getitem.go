package sms_home_advertise

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type getItemRequest struct{}

type getItemResponse struct{}

// GetItem 获取广告详情
// @Summary 获取广告详情
// @Description 获取广告详情
// @Tags SmsHomeAdvertiseController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getItemRequest true "请求信息"
// @Success 200 {object} code.Success{data=getItemResponse}
// @Failure 400 {object} code.Failure
// @Router /home/advertise/{id} [get]
func (h *handler) GetItem(ctx *gin.Context) {
	api.Success(ctx, nil)
}
