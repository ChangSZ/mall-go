package pms_sku_stock

import (
	"github.com/ChangSZ/mall-go/internal/services/pms_sku_stock"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// GetList 根据商品ID及sku编码模糊搜索sku库存
	// @Tags PmsSkuStockController
	// @Router /sku/{pid} [get]
	GetList(*gin.Context)

	// Update 批量更新sku库存信息
	// @Tags PmsSkuStockController
	// @Router /sku/update/{pid} [post]
	Update(*gin.Context)
}

type handler struct {
	service pms_sku_stock.Service
}

func New() Handler {
	return &handler{
		service: pms_sku_stock.New(),
	}
}

func (h *handler) i() {}
