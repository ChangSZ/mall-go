package sms_flash_promotion_product_relation

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type getItemRequest struct{}

type getItemResponse struct{}

// GetItem 获取关联商品促销信息
// @Summary 获取关联商品促销信息
// @Description 获取关联商品促销信息
// @Tags SmsFlashPromotionProductRelationController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getItemRequest true "请求信息"
// @Success 200 {object} code.Success{data=getItemResponse}
// @Failure 400 {object} code.Failure
// @Router /flashProductRelation/{id} [get]
func (h *handler) GetItem(ctx *gin.Context) {
	api.Success(ctx, nil)
}
