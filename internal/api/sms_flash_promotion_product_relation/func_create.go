package sms_flash_promotion_product_relation

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type createRequest struct{}

type createResponse struct{}

// Create 批量选择商品添加关联
// @Summary 批量选择商品添加关联
// @Description 批量选择商品添加关联
// @Tags SmsFlashPromotionProductRelationController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} code.Success{data=createResponse}
// @Failure 400 {object} code.Failure
// @Router /flashProductRelation [post]
func (h *handler) Create(ctx *gin.Context) {
	api.Success(ctx, nil)
}
