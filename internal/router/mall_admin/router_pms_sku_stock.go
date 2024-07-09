package mall_admin

import (
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api/mall_admin/pms_sku_stock"
	"github.com/ChangSZ/mall-go/internal/middleware"
)

// 商品SKU库存管理
func setPmsSkuStockRouter(eng *gin.Engine) {
	skuHandler := pms_sku_stock.New()
	skus := eng.Group("/sku", middleware.CheckToken(), middleware.DynamicAccess())
	{
		skus.GET("/:pid", skuHandler.GetList)        // 根据商品ID及sku编码模糊搜索sku库存
		skus.POST("/update/:pid", skuHandler.Update) // 批量更新sku库存信息
	}
}
