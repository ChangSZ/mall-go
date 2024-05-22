package mall_admin

import (
	"github.com/ChangSZ/mall-go/internal/api/mall_admin/pms_product_cate"
	"github.com/ChangSZ/mall-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 商品分类管理
func setPmsProducteCateRouter(eng *gin.Engine) {
	productCateHandler := pms_product_cate.New()
	cates := eng.Group("/productCategory", middleware.CheckToken(), middleware.DynamicAccess())
	{
		cates.POST("/create", productCateHandler.Create)                      // 添加商品分类
		cates.POST("/update/:id", productCateHandler.Update)                  // 修改商品分类
		cates.GET("/list/:parentId", productCateHandler.List)                 // 查询商品
		cates.GET("/:id", productCateHandler.GetItem)                         // 根据id获取商品分类
		cates.POST("/delete/:id", productCateHandler.Delete)                  // 删除商品分类
		cates.POST("/update/navStatus", productCateHandler.UpdateNavStatus)   // 修改导航栏显示状态
		cates.POST("/update/showStatus", productCateHandler.UpdateShowStatus) // 修改显示状态
		cates.GET("/list/withChildren", productCateHandler.ListWithChildren)  // 查询所有一级分类及子分类
	}
}
