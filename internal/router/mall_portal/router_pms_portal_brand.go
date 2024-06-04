package mall_portal

import (
	"github.com/ChangSZ/mall-go/internal/api/mall_portal/pms_portal_brand"
	"github.com/ChangSZ/mall-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 首页品牌推荐管理
func setPmsPortalBrandRouter(eng *gin.Engine) {
	handler := pms_portal_brand.New()
	group := eng.Group("/brand", middleware.CheckMemberToken())
	{
		group.GET("/recommendList", handler.RecommendList) // 分页获取推荐品牌
		group.GET("/detail/:brandId", handler.Detail)      // 获取品牌详情
		group.GET("/productList", handler.ProductList)     // 分页获取品牌相关商品
	}
}
