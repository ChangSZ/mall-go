package mall_portal

import (
	"github.com/ChangSZ/mall-go/internal/api/mall_portal/member_read_history"
	"github.com/ChangSZ/mall-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 会员商品浏览记录管理
func setMemberReadHistoryRouter(eng *gin.Engine) {
	handler := member_read_history.New()
	group := eng.Group("/member/readHistory", middleware.CheckMemberToken())
	{
		group.POST("/create", handler.Create) // 创建浏览记录
		group.POST("/delete", handler.Delete) // 删除浏览记录
		group.GET("/list", handler.List)      // 分页获取浏览记录
		group.POST("/clear", handler.Clear)   // 清空浏览记录
	}
}
