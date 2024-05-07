package sms_home_recommend_product

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type updateRecommendStatusRequest struct{}

type updateRecommendStatusResponse struct{}

// UpdateRecommendStatus 批量修改推荐状态
// @Summary 批量修改推荐状态
// @Description 批量修改推荐状态
// @Tags SmsHomeRecommendProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRecommendStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateRecommendStatusResponse}
// @Failure 400 {object} code.Failure
// @Router /home/recommendProduct/update/recommendStatus [post]
func (h *handler) UpdateRecommendStatus(ctx *gin.Context) {
	api.Success(ctx, nil)
}
