package router

import (
	"github.com/ChangSZ/mall-go/internal/api/sms_home_brand"
	"github.com/ChangSZ/mall-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 首页品牌管理
func setSmsHomeBrandRouter(eng *gin.Engine) {
	handler := sms_home_brand.New()
	group := eng.Group("/home/brand", middleware.CheckToken())
	{
		group.POST("", handler.Create)                                              // 添加首页推荐品牌
		group.POST("/update/sort/:id", handler.UpdateSort)                          // 修改推荐品牌排序
		group.POST("/update/status/recommendStatus", handler.UpdateRecommendStatus) // 批量修改推荐品牌状态
		group.POST("/delete", handler.Delete)                                       // 批量删除推荐品牌
		group.GET("/list", handler.List)                                            // 分页查询推荐品牌
	}
}
