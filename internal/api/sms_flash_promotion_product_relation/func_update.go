package sms_flash_promotion_product_relation

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type updateRequest struct{}

type updateResponse struct{}

// Update 修改关联信息
// @Summary 修改关联信息
// @Description 修改关联信息
// @Tags SmsFlashPromotionProductRelationController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateResponse}
// @Failure 400 {object} code.Failure
// @Router /flashProductRelation/update/{id} [post]
func (h *handler) Update(ctx *gin.Context) {
	api.Success(ctx, nil)
}
