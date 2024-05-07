package sms_flash_promotion_session

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type updateStatusRequest struct{}

type updateStatusResponse struct{}

// UpdateStatus 修改启用状态
// @Summary 修改启用状态
// @Description 修改启用状态
// @Tags SmsFlashPromotionSessionController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateStatusResponse}
// @Failure 400 {object} code.Failure
// @Router /flashSession/update/status/{id} [post]
func (h *handler) UpdateStatus(ctx *gin.Context) {
	api.Success(ctx, nil)
}
