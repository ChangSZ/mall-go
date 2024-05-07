package sms_flash_promotion_product_relation

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type listRequest struct{}

type listResponse struct{}

// List 分页查询不同场次关联及商品信息
// @Summary 分页查询不同场次关联及商品信息
// @Description 分页查询不同场次关联及商品信息
// @Tags SmsFlashPromotionProductRelationController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /flashProductRelation/list [get]
func (h *handler) List(ctx *gin.Context) {
	api.Success(ctx, nil)
}
