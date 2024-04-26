package pms_product

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type updateRecommendStatusRequest struct{}

type updateRecommendStatusResponse struct{}

// UpdateRecommendStatus 批量推荐商品
// @Summary 批量推荐商品
// @Description 批量推荐商品
// @Tags PmsProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRecommendStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateRecommendStatusResponse}
// @Failure 400 {object} code.Failure
// @Router /product/update/recommendStatus [post]
func (h *handler) UpdateRecommendStatus(ctx *gin.Context) {
	api.Success(ctx, nil)
}
