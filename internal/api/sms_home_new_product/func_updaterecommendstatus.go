package sms_home_new_product

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type updateRecommendStatusRequest struct{}

type updateRecommendStatusResponse struct{}

// UpdateRecommendStatus 批量修改首页新品状态
// @Summary 批量修改首页新品状态
// @Description 批量修改首页新品状态
// @Tags SmsHomeNewProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRecommendStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateRecommendStatusResponse}
// @Failure 400 {object} code.Failure
// @Router /home/newProduct/update/recommendStatus [post]
func (h *handler) UpdateRecommendStatus(ctx *gin.Context) {
	api.Success(ctx, nil)
}
