package pms_sku_stock

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type updateRequest struct{}

type updateResponse struct{}

// Update 批量更新sku库存信息
// @Summary 批量更新sku库存信息
// @Description 批量更新sku库存信息
// @Tags PmsSkuStockController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateResponse}
// @Failure 400 {object} code.Failure
// @Router /sku/update/{pid} [post]
func (h *handler) Update(ctx *gin.Context) {
	api.Success(ctx, nil)
}
