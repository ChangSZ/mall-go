package mall_admin

import (
	"github.com/ChangSZ/mall-go/internal/api/mall_admin/sms_flash_promotion_session"
	"github.com/ChangSZ/mall-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 限时购场次管理
func setSmsFlashPromotionSessionRouter(eng *gin.Engine) {
	handler := sms_flash_promotion_session.New()
	group := eng.Group("/flashSession", middleware.CheckToken(), middleware.DynamicAccess())
	{
		group.POST("", handler.Create)                         // 添加活动
		group.POST("/update/:id", handler.Update)              // 编辑活动
		group.POST("/update/status/:id", handler.UpdateStatus) // 修改上下线状态
		group.POST("/delete/:id", handler.Delete)              // 删除活动
		group.GET("/list", handler.List)                       // 根据活动名称分页查询
		group.GET("/:id", handler.GetItem)                     // 获取活动详情
	}
}
