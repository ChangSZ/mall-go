package sms_flash_promotion_session

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type getItemRequest struct{}

type getItemResponse struct{}

// GetItem 获取场次详情
// @Summary 获取场次详情
// @Description 获取场次详情
// @Tags SmsFlashPromotionSessionController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getItemRequest true "请求信息"
// @Success 200 {object} code.Success{data=getItemResponse}
// @Failure 400 {object} code.Failure
// @Router /flashSession/{id} [get]
func (h *handler) GetItem(ctx *gin.Context) {
	api.Success(ctx, nil)
}
