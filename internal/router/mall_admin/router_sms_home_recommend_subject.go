package mall_admin

import (
	"github.com/ChangSZ/mall-go/internal/api/mall_admin/sms_home_recommend_subject"
	"github.com/ChangSZ/mall-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 首页专题推荐管理
func setSmsHomeRecommendSubjectRouter(eng *gin.Engine) {
	handler := sms_home_recommend_subject.New()
	group := eng.Group("/home/recommendSubject", middleware.CheckToken())
	{
		group.POST("", handler.Create)                                              // 添加首页推荐专题
		group.POST("/update/sort/:id", handler.UpdateSort)                          // 修改推荐专题排序
		group.POST("/update/status/recommendStatus", handler.UpdateRecommendStatus) // 批量修改推荐专题状态
		group.POST("/delete", handler.Delete)                                       // 批量删除推荐专题
		group.GET("/list", handler.List)                                            // 分页查询推荐专题
	}
}
