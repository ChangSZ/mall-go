package sms_home_recommend_product

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type updateRecommendStatusRequest struct {
	Ids             []int64 `form:"ids" binding:"required"`
	RecommendStatus int32   `form:"recommendStatus" binding:"required"`
}

type updateRecommendStatusResponse struct {
	Count int64 `json:",inline"`
}

// UpdateRecommendStatus 批量修改推荐状态
// @Summary 批量修改推荐状态
// @Description 批量修改推荐状态
// @Tags SmsHomeRecommendProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRecommendStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /home/recommendProduct/update/recommendStatus [post]
func (h *handler) UpdateRecommendStatus(ctx *gin.Context) {
	req := new(updateRecommendStatusRequest)
	res := new(updateRecommendStatusResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	cnt, err := h.service.UpdateRecommendStatus(ctx, req.Ids, req.RecommendStatus)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
