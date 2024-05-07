package sms_home_recommend_product

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type updateSortRequest struct{}

type updateSortResponse struct{}

// UpdateSort 修改推荐排序
// @Summary 修改推荐排序
// @Description 修改推荐排序
// @Tags SmsHomeRecommendProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateSortRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateSortResponse}
// @Failure 400 {object} code.Failure
// @Router /home/recommendProduct/update/sort/{id} [post]
func (h *handler) UpdateSort(ctx *gin.Context) {
	api.Success(ctx, nil)
}
