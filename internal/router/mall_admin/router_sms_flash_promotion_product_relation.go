package mall_admin

import (
	"github.com/ChangSZ/mall-go/internal/api/mall_admin/sms_flash_promotion_product_relation"
	"github.com/ChangSZ/mall-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 限时购和商品关系管理
func setSmsFlashPromotionProductRelationRouter(eng *gin.Engine) {
	handler := sms_flash_promotion_product_relation.New()
	group := eng.Group("/flashProductRelation", middleware.CheckToken())
	{
		group.POST("", handler.Create)            // 批量选择商品添加关联
		group.POST("/update/:id", handler.Update) // 修改关联信息
		group.POST("/delete/:id", handler.Delete) // 删除关联
		group.GET("/list", handler.List)          // 分页查询不同场次关联及商品信息
		group.GET("/:id", handler.GetItem)        // 获取关联商品促销信息
	}
}
