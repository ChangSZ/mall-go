package router

import (
	"github.com/ChangSZ/mall-go/internal/api/ums_resource"

	"github.com/gin-gonic/gin"
)

// 后台资源管理
func setUmsResourceRouter(eng *gin.Engine) {
	resourceHandler := ums_resource.New()
	resources := eng.Group("/resource")
	{
		resources.POST("/create", resourceHandler.Create)     // 添加后台资源
		resources.POST("/update/:id", resourceHandler.Update) // 修改后台资源
		resources.GET("/:id", resourceHandler.Get)            // 根据ID获取资源详情
		resources.POST("/delete/:id", resourceHandler.Delete) // 根据ID删除后台资源
		resources.GET("/list", resourceHandler.List)          // 分页模糊查询后台资源
		resources.GET("/listAll", resourceHandler.ListAll)    // 查询所有后台资源
	}
}
