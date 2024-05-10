package sms_flash_promotion_product_relation

import (
	"github.com/ChangSZ/mall-go/internal/services/sms_flash_promotion_product_relation"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 批量选择商品添加关联
	// @Tags SmsFlashPromotionProductRelationController
	// @Router /flashProductRelation [post]
	Create(*gin.Context)

	// Update 修改关联信息
	// @Tags SmsFlashPromotionProductRelationController
	// @Router /flashProductRelation/update/{id} [post]
	Update(*gin.Context)

	// Delete 删除关联
	// @Tags SmsFlashPromotionProductRelationController
	// @Router /flashProductRelation/delete/{id} [post]
	Delete(*gin.Context)

	// List 分页查询不同场次关联及商品信息
	// @Tags SmsFlashPromotionProductRelationController
	// @Router /flashProductRelation/list [get]
	List(*gin.Context)

	// GetItem 获取关联商品促销信息
	// @Tags SmsFlashPromotionProductRelationController
	// @Router /flashProductRelation/{id} [get]
	GetItem(*gin.Context)
}

type handler struct {
	service sms_flash_promotion_product_relation.Service
}

func New() Handler {
	return &handler{
		service: sms_flash_promotion_product_relation.New(),
	}
}

func (h *handler) i() {}
