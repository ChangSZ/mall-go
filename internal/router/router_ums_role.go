package router

import (
	"github.com/ChangSZ/mall-go/internal/api/ums_role"

	"github.com/gin-gonic/gin"
)

// 后台用户角色管理
func setUmsRoleRouter(eng *gin.Engine) {
	roleHandler := ums_role.New()
	roles := eng.Group("/role")
	{
		roles.POST("/create", roleHandler.Create)                    // 添加角色
		roles.POST("/update/:id", roleHandler.Update)                // 修改角色
		roles.POST("/delete", roleHandler.Delete)                    // 批量删除角色
		roles.GET("/listAll", roleHandler.ListAll)                   // 获取所有角色
		roles.GET("/list", roleHandler.List)                         // 根据角色名称分页获取角色列表
		roles.POST("/updateStatus/:id", roleHandler.UpdateStatus)    // 修改角色状态
		roles.GET("/listMenu/:roleId", roleHandler.ListMenu)         // 获取角色相关菜单
		roles.GET("/listResource/:roleId", roleHandler.ListResource) // 获取角色相关资源
		roles.POST("/allocMenu", roleHandler.AllocMenu)              // 给角色分配菜单
		roles.POST("/allocResource", roleHandler.AllocResource)      // 给角色分配资源
	}
}
