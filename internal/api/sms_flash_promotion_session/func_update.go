package sms_flash_promotion_session

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type updateRequest struct{}

type updateResponse struct{}

// Update 修改场次
// @Summary 修改场次
// @Description 修改场次
// @Tags SmsFlashPromotionSessionController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateResponse}
// @Failure 400 {object} code.Failure
// @Router /flashSession/update/{id} [post]
func (h *handler) Update(ctx *gin.Context) {
	api.Success(ctx, nil)
}
