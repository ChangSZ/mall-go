package mall_admin

import (
	"github.com/ChangSZ/mall-go/internal/api/mall_admin/sms_home_advertise"
	"github.com/ChangSZ/mall-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 首页轮播广告管理
func setSmsHomeAdvertiseRouter(eng *gin.Engine) {
	handler := sms_home_advertise.New()
	group := eng.Group("/home/advertise", middleware.CheckToken(), middleware.DynamicAccess())
	{
		group.POST("", handler.Create)                         // 添加广告
		group.POST("/update/:id", handler.Update)              // 修改广告
		group.POST("/update/status/:id", handler.UpdateStatus) // 修改上下线状态
		group.POST("/delete", handler.Delete)                  // 删除广告
		group.GET("/list", handler.List)                       // 分页查询广告
		group.GET("/:id", handler.GetItem)                     // 获取广告详情
	}
}
