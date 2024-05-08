package sms_flash_promotion

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type getItemRequest struct{}

type getItemResponse struct {
	dto.SmsFlashPromotion `json:",inline"`
}

// GetItem 获取活动详情
// @Summary 获取活动详情
// @Description 获取活动详情
// @Tags SmsFlashPromotionController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getItemRequest true "请求信息"
// @Success 200 {object} code.Success{data=getItemResponse}
// @Failure 400 {object} code.Failure
// @Router /flash/{id} [get]
func (h *handler) GetItem(ctx *gin.Context) {
	_ = new(getItemRequest)
	res := new(getItemResponse)
	uri := new(dto.UriID)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	item, err := h.smsFlashPromotionService.GetItem(ctx, uri.Id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.SmsFlashPromotion = *item
	api.Success(ctx, res)
}
