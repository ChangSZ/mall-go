package mall_portal

import (
	"github.com/ChangSZ/mall-go/internal/api/mall_portal/member_collection"
	"github.com/ChangSZ/mall-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 会员商品收藏管理
func setMemberCollectionRouter(eng *gin.Engine) {
	handler := member_collection.New()
	group := eng.Group("/member/productCollection", middleware.CheckMemberToken())
	{
		group.POST("/add", handler.Add)       // 添加商品收藏
		group.POST("/delete", handler.Delete) // 删除商品收藏
		group.GET("/list", handler.List)      // 显示当前用户商品收藏列表
		group.GET("/detail", handler.Detail)  // 显示商品收藏详情
		group.POST("/clear", handler.Clear)   // 清空当前用户商品收藏列表
	}
}
