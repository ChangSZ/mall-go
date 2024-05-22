package mall_admin

import (
	"github.com/ChangSZ/mall-go/internal/api/mall_admin/pms_product_attr_cate"
	"github.com/ChangSZ/mall-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 商品属性分类管理
func setPmsProducteAttrCateRouter(eng *gin.Engine) {
	productAttrCateHandler := pms_product_attr_cate.New()
	attrCates := eng.Group("/productAttribute/category", middleware.CheckToken(), middleware.DynamicAccess())
	{
		attrCates.POST("/create", productAttrCateHandler.Create)                // 添加商品属性分类
		attrCates.POST("/update/:id", productAttrCateHandler.Update)            // 修改商品属性分类
		attrCates.POST("/delete/:id", productAttrCateHandler.Delete)            // 删除单个商品属性分类
		attrCates.GET("/:id", productAttrCateHandler.GetItem)                   // 获取单个商品属性分类信息
		attrCates.GET("/list", productAttrCateHandler.List)                     // 分页获取所有商品属性分类
		attrCates.GET("/list/withAttr", productAttrCateHandler.GetListWithAttr) // 获取所有商品属性分类及其下属性
	}
}
