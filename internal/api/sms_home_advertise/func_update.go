package sms_home_advertise

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type updateRequest struct{}

type updateResponse struct{}

// Update 修改广告
// @Summary 修改广告
// @Description 修改广告
// @Tags SmsHomeAdvertiseController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateResponse}
// @Failure 400 {object} code.Failure
// @Router /home/advertise/update/{id} [post]
func (h *handler) Update(ctx *gin.Context) {
	api.Success(ctx, nil)
}
