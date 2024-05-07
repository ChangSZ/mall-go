package sms_flash_promotion_session

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type selectListRequest struct{}

type selectListResponse struct{}

// SelectList 获取全部可选场次及其数量
// @Summary 获取全部可选场次及其数量
// @Description 获取全部可选场次及其数量
// @Tags SmsFlashPromotionSessionController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body selectListRequest true "请求信息"
// @Success 200 {object} code.Success{data=selectListResponse}
// @Failure 400 {object} code.Failure
// @Router /flashSession/selectList [get]
func (h *handler) SelectList(ctx *gin.Context) {
	api.Success(ctx, nil)
}
