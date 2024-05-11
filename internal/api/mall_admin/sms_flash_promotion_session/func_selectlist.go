package sms_flash_promotion_session

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type selectListRequest struct {
	FlashPromotionId int64 `json:"flashPromotionId" binding:"omitempty"`
}

type selectListResponse struct {
	List []dto.SmsFlashPromotionSessionDetail `json:",inline"`
}

// SelectList 获取全部可选场次及其数量
// @Summary 获取全部可选场次及其数量
// @Description 获取全部可选场次及其数量
// @Tags SmsFlashPromotionSessionController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body selectListRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]dto.SmsFlashPromotionSessionDetail}
// @Failure 400 {object} code.Failure
// @Router /flashSession/selectList [get]
func (h *handler) SelectList(ctx *gin.Context) {
	req := new(selectListRequest)
	res := new(selectListResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	list, err := h.service.SelectList(ctx, req.FlashPromotionId)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = list
	api.Success(ctx, res.List)
}
