package router

import (
	"github.com/ChangSZ/mall-go/internal/api/sms_flash_promotion"
	"github.com/ChangSZ/mall-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 限时购活动管理
func setSmsFlashPromotionRouter(eng *gin.Engine) {
	handler := sms_flash_promotion.New()
	group := eng.Group("/flash", middleware.CheckToken())
	{
		group.POST("", handler.Create)                         // 添加活动
		group.POST("/update/:id", handler.Update)              // 编辑活动
		group.POST("/update/status/:id", handler.UpdateStatus) // 修改上下线状态
		group.POST("/delete/:id", handler.Delete)              // 删除活动
		group.GET("/list", handler.List)                       // 根据活动名称分页查询
		group.GET("/:id", handler.GetItem)                     // 获取活动详情
	}
}
