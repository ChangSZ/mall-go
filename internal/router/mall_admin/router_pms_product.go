package mall_admin

import (
	"github.com/ChangSZ/mall-go/internal/api/mall_admin/pms_product"
	"github.com/ChangSZ/mall-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 商品管理
func setPmsProducteRouter(eng *gin.Engine) {
	productHandler := pms_product.New()
	products := eng.Group("/product", middleware.CheckToken())
	{
		products.POST("/create", productHandler.Create)                                // 创建商品
		products.GET("/updateInfo/:id", productHandler.GetUpdateInfo)                  // 根据商品id获取商品编辑信息
		products.POST("/update/:id", productHandler.Update)                            // 更新商品
		products.GET("/list", productHandler.List)                                     // 查询商品
		products.GET("/simpleList", productHandler.SimpleList)                         // 根据商品名称或货号模糊查询
		products.POST("/update/verifyStatus", productHandler.UpdateVerifyStatus)       // 批量修改审核状态
		products.POST("/update/publishStatus", productHandler.UpdatePublishStatus)     // 批量上下架商品
		products.POST("/update/recommendStatus", productHandler.UpdateRecommendStatus) // 批量推荐商品
		products.POST("/update/newStatus", productHandler.UpdateNewStatus)             // 批量设为新品
		products.POST("/update/deleteStatus", productHandler.UpdateDeleteStatus)       // 批量修改删除状态
	}
}
