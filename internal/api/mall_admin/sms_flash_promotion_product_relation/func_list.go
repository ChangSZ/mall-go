package sms_flash_promotion_product_relation

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/pagehelper"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type listRequest struct {
	FlashPromotionId        int64  `form:"flashPromotionId" binding:"omitempty"`
	FlashPromotionSessionId int64  `form:"flashPromotionSessionId" binding:"omitempty"`
	Keyword                 string `form:"keyword" binding:"omitempty"`
	PageSize                int    `form:"pageSize,default=5" binding:"omitempty"`
	PageNum                 int    `form:"pageNum,default=1" binding:"omitempty"`
}

type listResponse struct {
	*pagehelper.ListData[dto.SmsFlashPromotionProductRelation] `json:",inline"`
}

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
	req := new(listRequest)
	res := new(listResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	list, err := h.service.List(
		ctx, req.FlashPromotionId, req.FlashPromotionSessionId, req.PageSize, req.PageNum)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.ListData = list
	api.Success(ctx, res)
}
