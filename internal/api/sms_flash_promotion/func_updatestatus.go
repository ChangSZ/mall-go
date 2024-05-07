package sms_flash_promotion

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type updateStatusRequest struct{}

type updateStatusResponse struct{}

// UpdateStatus 编辑活动
// @Summary 编辑活动
// @Description 编辑活动
// @Tags SmsFlashPromotionController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateStatusResponse}
// @Failure 400 {object} code.Failure
// @Router /flash/update/status/{id} [post]
func (h *handler) UpdateStatus(ctx *gin.Context) {
	api.Success(ctx, nil)
}
