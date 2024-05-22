package mall_admin

import (
	"github.com/ChangSZ/mall-go/internal/api/mall_admin/sms_home_new_product"
	"github.com/ChangSZ/mall-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 首页新品管理
func setSmsHomeNewProductRouter(eng *gin.Engine) {
	handler := sms_home_new_product.New()
	group := eng.Group("/home/newProduct", middleware.CheckToken(), middleware.DynamicAccess())
	{
		group.POST("", handler.Create)                                              // 添加首页新品
		group.POST("/update/sort/:id", handler.UpdateSort)                          // 修改首页新品排序
		group.POST("/update/status/recommendStatus", handler.UpdateRecommendStatus) // 批量修改首页新品状态
		group.POST("/delete", handler.Delete)                                       // 批量删除首页新品
		group.GET("/list", handler.List)                                            // 分页查询首页新品
	}
}
