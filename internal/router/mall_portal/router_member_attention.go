package mall_portal

import (
	"github.com/ChangSZ/mall-go/internal/api/mall_portal/member_attention"
	"github.com/ChangSZ/mall-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 会员关注品牌管理
func setMemberAttentionRouter(eng *gin.Engine) {
	handler := member_attention.New()
	group := eng.Group("/member/attention", middleware.CheckMemberToken())
	{
		group.POST("/add", handler.Add)       // 添加品牌关注
		group.POST("/delete", handler.Delete) // 取消品牌关注
		group.GET("/list", handler.List)      // 分页查询当前用户品牌关注列表
		group.GET("/detail", handler.Detail)  // 根据品牌ID获取品牌关注详情
		group.POST("/clear", handler.Clear)   // 清空当前用户品牌关注列表
	}
}
