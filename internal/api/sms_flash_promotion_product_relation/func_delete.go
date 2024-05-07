package sms_flash_promotion_product_relation

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type deleteRequest struct{}

type deleteResponse struct{}

// Delete 删除关联
// @Summary 删除关联
// @Description 删除关联
// @Tags SmsFlashPromotionProductRelationController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body deleteRequest true "请求信息"
// @Success 200 {object} code.Success{data=deleteResponse}
// @Failure 400 {object} code.Failure
// @Router /flashProductRelation/delete/{id} [post]
func (h *handler) Delete(ctx *gin.Context) {
	api.Success(ctx, nil)
}
