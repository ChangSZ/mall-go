package pms_product

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type updateRecommendStatusRequest struct {
	Ids             []int64 `form:"ids"`
	RecommendStatus int32   `form:"recommendStatus"`
}

type updateRecommendStatusResponse struct {
	Count int64 `json:",inline"`
}

// UpdateRecommendStatus 批量推荐商品
// @Summary 批量推荐商品
// @Description 批量推荐商品
// @Tags PmsProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRecommendStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /product/update/recommendStatus [post]
func (h *handler) UpdateRecommendStatus(ctx *gin.Context) {
	req := new(updateRecommendStatusRequest)
	res := new(updateRecommendStatusResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	cnt, err := h.service.UpdateRecommendStatus(ctx, req.Ids, req.RecommendStatus)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	if cnt == 0 {
		api.Failed(ctx, "更新RecommendStatus个数为0")
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
