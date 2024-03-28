package router

import (
	"github.com/ChangSZ/mall-go/internal/api/ums_resource"
)

// 后台资源管理
func setUmsResourceRouter(r *resource) {
	resourceHandler := ums_resource.New(r.logger, r.db, r.cache)
	resources := r.mux.Group("/resource")
	{
		resources.POST("/create", resourceHandler.Create())     // 添加后台资源
		resources.POST("/update/:id", resourceHandler.Update()) // 修改后台资源
		resources.GET("/:id", resourceHandler.Get())            // 根据ID获取资源详情
		resources.POST("/delete/:id", resourceHandler.Delete()) // 根据ID删除后台资源
		resources.GET("/list", resourceHandler.List())          // 分页模糊查询后台资源
		resources.GET("/listAll", resourceHandler.ListAll())    // 查询所有后台资源
	}
}
