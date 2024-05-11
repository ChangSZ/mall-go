package mall_admin

import (
	"github.com/ChangSZ/mall-go/internal/api/mall_admin/sms_home_recommend_product"
	"github.com/ChangSZ/mall-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 首页人气推荐管理
func setSmsHomeRecommendProductRouter(eng *gin.Engine) {
	handler := sms_home_recommend_product.New()
	group := eng.Group("/home/recommendProduct", middleware.CheckToken())
	{
		group.POST("", handler.Create)                                              // 添加首页推荐
		group.POST("/update/sort/:id", handler.UpdateSort)                          // 修改推荐排序
		group.POST("/update/status/recommendStatus", handler.UpdateRecommendStatus) // 批量修改推荐状态
		group.POST("/delete", handler.Delete)                                       // 批量删除推荐
		group.GET("/list", handler.List)                                            // 分页查询推荐
	}
}
