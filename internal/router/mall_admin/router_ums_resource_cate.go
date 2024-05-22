package mall_admin

import (
	"github.com/ChangSZ/mall-go/internal/api/mall_admin/ums_resource_cate"
	"github.com/ChangSZ/mall-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 后台资源分类管理
func setUmsResourceCateRouter(eng *gin.Engine) {
	resourceCateHandler := ums_resource_cate.New()
	resourceCate := eng.Group("/resourceCategory", middleware.CheckToken(), middleware.DynamicAccess())
	{
		resourceCate.GET("/listAll", resourceCateHandler.ListAll)    // 查询所有后台资源分类
		resourceCate.POST("/create", resourceCateHandler.Create)     // 添加后台资源分类
		resourceCate.POST("/update/:id", resourceCateHandler.Update) // 修改后台资源分类
		resourceCate.POST("/delete/:id", resourceCateHandler.Delete) // 根据ID删除后台资源分类
	}
}
