package sms_flash_promotion

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type updateRequest struct{}

type updateResponse struct{}

// Update 编辑活动
// @Summary 编辑活动
// @Description 编辑活动
// @Tags SmsFlashPromotionController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateResponse}
// @Failure 400 {object} code.Failure
// @Router /flash/update/{id} [post]
func (h *handler) Update(ctx *gin.Context) {
	api.Success(ctx, nil)
}
