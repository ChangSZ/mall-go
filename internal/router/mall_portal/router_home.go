package mall_portal

import (
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api/mall_portal/home"
	"github.com/ChangSZ/mall-go/internal/middleware"
)

// 首页内容管理
func setHomeRouter(eng *gin.Engine) {
	handler := home.New()
	group := eng.Group("/home", middleware.CheckMemberToken())
	{
		group.GET("/content", handler.Content)                              // 首页内容信息展示
		group.GET("/recommendProductList", handler.RecommendProductList)    // 分页获取推荐商品
		group.GET("/productCateList/:parentId", handler.GetProductCateList) // 获取首页商品分类
		group.GET("/subjectList", handler.GetSubjectList)                   // 根据分类获取专题
		group.GET("/hotProductList", handler.HotProductList)                // 分页获取人气推荐商品
		group.GET("/newProductList", handler.NewProductList)                // 分页获取新品推荐商品
	}
}
