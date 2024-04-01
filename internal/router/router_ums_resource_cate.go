package router

import (
	"github.com/ChangSZ/mall-go/internal/api/ums_resource_cate"
)

// 后台资源分类管理
func setUmsResourceCateRouter(r *resource) {
	resourceCateHandler := ums_resource_cate.New(r.logger, r.db, r.cache)
	resourceCate := r.mux.Group("/resourceCategory")
	{
		resourceCate.POST("/listAll", resourceCateHandler.ListAll())   // 查询所有后台资源分类
		resourceCate.POST("/create", resourceCateHandler.Create())     // 添加后台资源分类
		resourceCate.POST("/update/:id", resourceCateHandler.Update()) // 修改后台资源分类
		resourceCate.GET("/delete/:id", resourceCateHandler.Delete())  // 根据ID删除后台资源分类
	}
}
