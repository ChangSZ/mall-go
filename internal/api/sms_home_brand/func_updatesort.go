package sms_home_brand

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type updateSortRequest struct{}

type updateSortResponse struct{}

// UpdateSort 修改推荐品牌排序
// @Summary 修改推荐品牌排序
// @Description 修改推荐品牌排序
// @Tags SmsHomeBrandController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateSortRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateSortResponse}
// @Failure 400 {object} code.Failure
// @Router /home/brand/update/sort/{id} [post]
func (h *handler) UpdateSort(ctx *gin.Context) {
	api.Success(ctx, nil)
}
