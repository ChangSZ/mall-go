package sms_flash_promotion_product_relation

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type createRequest struct{}

type createResponse struct {
	Count int64 `json:",inline"`
}

// Create 批量选择商品添加关联
// @Summary 批量选择商品添加关联
// @Description 批量选择商品添加关联
// @Tags SmsFlashPromotionProductRelationController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body []SmsFlashPromotionProductRelation true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /flashProductRelation [post]
func (h *handler) Create(ctx *gin.Context) {
	_ = new(createRequest)
	req := make([]dto.SmsFlashPromotionProductRelation, 0)
	res := new(createResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	cnt, err := h.service.Create(ctx, req)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
