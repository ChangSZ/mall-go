package pms_sku_stock

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type getListRequest struct{}

type getListResponse struct{}

// GetList 根据商品ID及sku编码模糊搜索sku库存
// @Summary 根据商品ID及sku编码模糊搜索sku库存
// @Description 根据商品ID及sku编码模糊搜索sku库存
// @Tags PmsSkuStockController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getListRequest true "请求信息"
// @Success 200 {object} code.Success{data=getListResponse}
// @Failure 400 {object} code.Failure
// @Router /sku/{pid} [get]
func (h *handler) GetList(ctx *gin.Context) {
	api.Success(ctx, nil)
}
