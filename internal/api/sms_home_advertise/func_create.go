package sms_home_advertise

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type createRequest struct{}

type createResponse struct{}

// Create 添加广告
// @Summary 添加广告
// @Description 添加广告
// @Tags SmsHomeAdvertiseController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} code.Success{data=createResponse}
// @Failure 400 {object} code.Failure
// @Router /home/advertise [post]
func (h *handler) Create(ctx *gin.Context) {
	api.Success(ctx, nil)
}
