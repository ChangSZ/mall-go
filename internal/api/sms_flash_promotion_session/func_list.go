package sms_flash_promotion_session

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type listRequest struct{}

type listResponse struct{}

// List 获取全部场次
// @Summary 获取全部场次
// @Description 获取全部场次
// @Tags SmsFlashPromotionSessionController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /flashSession/list [get]
func (h *handler) List(ctx *gin.Context) {
	api.Success(ctx, nil)
}
