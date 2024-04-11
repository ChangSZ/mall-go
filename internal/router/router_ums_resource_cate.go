package router

import (
	"github.com/ChangSZ/mall-go/internal/api/ums_resource_cate"
	"github.com/gin-gonic/gin"
)

// 后台资源分类管理
func setUmsResourceCateRouter(eng *gin.Engine) {
	resourceCateHandler := ums_resource_cate.New()
	resourceCate := eng.Group("/resourceCategory")
	{
		resourceCate.POST("/listAll", resourceCateHandler.ListAll)   // 查询所有后台资源分类
		resourceCate.POST("/create", resourceCateHandler.Create)     // 添加后台资源分类
		resourceCate.POST("/update/:id", resourceCateHandler.Update) // 修改后台资源分类
		resourceCate.GET("/delete/:id", resourceCateHandler.Delete)  // 根据ID删除后台资源分类
	}
}
