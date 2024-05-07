package sms_home_brand

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type updateRecommendStatusRequest struct{}

type updateRecommendStatusResponse struct{}

// UpdateRecommendStatus 批量修改推荐品牌状态
// @Summary 批量修改推荐品牌状态
// @Description 批量修改推荐品牌状态
// @Tags SmsHomeBrandController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRecommendStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateRecommendStatusResponse}
// @Failure 400 {object} code.Failure
// @Router /home/brand/update/recommendStatus [post]
func (h *handler) UpdateRecommendStatus(ctx *gin.Context) {
	api.Success(ctx, nil)
}
