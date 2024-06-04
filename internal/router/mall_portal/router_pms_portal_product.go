package mall_portal

import (
	"github.com/ChangSZ/mall-go/internal/api/mall_portal/pms_portal_product"
	"github.com/ChangSZ/mall-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 前台商品管理
func setPmsPortalProductRouter(eng *gin.Engine) {
	handler := pms_portal_product.New()
	group := eng.Group("/product", middleware.CheckMemberToken())
	{
		group.GET("/search", handler.Search)                     // 综合搜索、筛选、排序
		group.GET("/categoryTreeList", handler.CategoryTreeList) // 以树形结构获取所有商品分类
		group.GET("/detail/:id", handler.Detail)                 // 获取前台商品详情
	}
}
