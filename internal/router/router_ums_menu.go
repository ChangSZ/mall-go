package router

import (
	"github.com/ChangSZ/mall-go/internal/api/ums_menu"
	"github.com/gin-gonic/gin"
)

// 后台菜单管理
func setUmsMenuRouter(eng *gin.Engine) {
	menuHandler := ums_menu.New()
	menus := eng.Group("/menu")
	{
		menus.POST("/create", menuHandler.Create)                 // 添加后台菜单
		menus.POST("/update/:id", menuHandler.Update)             // 修改后台菜单
		menus.GET("/:id", menuHandler.Get)                        // 根据ID获取菜单详情
		menus.POST("/delete/:id", menuHandler.Delete)             // 根据ID删除后台菜单
		menus.GET("/list/:parentId", menuHandler.List)            // 分页查询后台菜单
		menus.GET("/treeList", menuHandler.TreeList)              // 树形结构返回所有菜单列表
		menus.POST("/updateHidden/:id", menuHandler.UpdateHidden) // 修改菜单显示状态
	}
}
