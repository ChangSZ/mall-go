package router

import (
	"github.com/ChangSZ/mall-go/internal/api/pms_product_attr"
	"github.com/ChangSZ/mall-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 商品属性管理
func setPmsProducteAttrRouter(eng *gin.Engine) {
	productAttrHandler := pms_product_attr.New()
	attrs := eng.Group("/productAttribute", middleware.CheckToken())
	{
		attrs.POST("/create", productAttrHandler.Create)                          // 添加商品属性信息
		attrs.POST("/update/:id", productAttrHandler.Update)                      // 修改商品属性信息
		attrs.GET("/list/:cid", productAttrHandler.List)                          // 根据分类查询属性列表或参数列表
		attrs.GET("/:id", productAttrHandler.GetItem)                             // 查询单个商品属性
		attrs.POST("/delete", productAttrHandler.Delete)                          // 批量删除商品属性
		attrs.GET("/attrInfo/:productCategoryId", productAttrHandler.GetAttrInfo) // 根据商品分类的id获取商品属性及属性分类
	}
}
