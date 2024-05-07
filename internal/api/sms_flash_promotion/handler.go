package sms_flash_promotion

import (
	"github.com/ChangSZ/mall-go/internal/services/sms_flash_promotion"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 添加活动
	// @Tags SmsFlashPromotionController
	// @Router /flash [post]
	Create(*gin.Context)

	// Update 编辑活动
	// @Tags SmsFlashPromotionController
	// @Router /flash/update/{id} [post]
	Update(*gin.Context)

	// UpdateStatus 编辑活动
	// @Tags SmsFlashPromotionController
	// @Router /flash/update/status/{id} [post]
	UpdateStatus(*gin.Context)

	// Delete 删除活动
	// @Tags SmsFlashPromotionController
	// @Router /flash/delete/{id} [post]
	Delete(*gin.Context)

	// List 根据活动名称分页查询
	// @Tags SmsFlashPromotionController
	// @Router /flash/list [get]
	List(*gin.Context)

	// GetItem 获取活动详情
	// @Tags SmsFlashPromotionController
	// @Router /flash/{id} [get]
	GetItem(*gin.Context)
}

type handler struct {
	smsFlashPromotionService sms_flash_promotion.Service
}

func New() Handler {
	return &handler{
		smsFlashPromotionService: sms_flash_promotion.New(),
	}
}

func (h *handler) i() {}
