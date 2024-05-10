package sms_flash_promotion_session

import (
	"github.com/ChangSZ/mall-go/internal/services/sms_flash_promotion_session"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 添加场次
	// @Tags SmsFlashPromotionSessionController
	// @Router /flashSession [post]
	Create(*gin.Context)

	// Update 修改场次
	// @Tags SmsFlashPromotionSessionController
	// @Router /flashSession/update/{id} [post]
	Update(*gin.Context)

	// UpdateStatus 修改启用状态
	// @Tags SmsFlashPromotionSessionController
	// @Router /flashSession/update/status/{id} [post]
	UpdateStatus(*gin.Context)

	// Delete 删除场次
	// @Tags SmsFlashPromotionSessionController
	// @Router /flashSession/delete/{id} [post]
	Delete(*gin.Context)

	// List 获取全部场次
	// @Tags SmsFlashPromotionSessionController
	// @Router /flashSession/list [get]
	List(*gin.Context)

	// GetItem 获取场次详情
	// @Tags SmsFlashPromotionSessionController
	// @Router /flashSession/{id} [get]
	GetItem(*gin.Context)

	// SelectList 获取全部可选场次及其数量
	// @Tags SmsFlashPromotionSessionController
	// @Router /flashSession/selectList [get]
	SelectList(*gin.Context)
}

type handler struct {
	service sms_flash_promotion_session.Service
}

func New() Handler {
	return &handler{
		service: sms_flash_promotion_session.New(),
	}
}

func (h *handler) i() {}
