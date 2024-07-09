package mall_admin

import (
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api/mall_admin/ums_admin"
	"github.com/ChangSZ/mall-go/internal/middleware"
)

// 后台用户管理
func setUmsAdminRouter(eng *gin.Engine) {
	adminHandler := ums_admin.New()
	admins := eng.Group("/admin")
	{
		admins.POST("/register", adminHandler.Register) // 用户注册
		admins.POST("/login", adminHandler.Login)       // 登录以后返回token
	}

	adminsM := eng.Group("/admin", middleware.CheckToken(), middleware.DynamicAccess())
	{
		adminsM.GET("/refreshToken", adminHandler.RefreshToken)      // 刷新token
		adminsM.GET("/info", adminHandler.Info)                      // 获取当前登录用户信息
		adminsM.POST("/logout", adminHandler.Logout)                 // 登出功能
		adminsM.GET("/list", adminHandler.List)                      // 根据用户名或姓名分页获取用户列表
		adminsM.GET("/:id", adminHandler.Get)                        // 获取指定用户信息
		adminsM.POST("/update/:id", adminHandler.Update)             // 修改指定用户信息
		adminsM.POST("/updatePassword", adminHandler.UpdatePassword) // 修改指定用户密码
		adminsM.POST("/delete/:id", adminHandler.Delete)             // 删除指定用户信息
		adminsM.POST("/updateStatus/:id", adminHandler.UpdateStatus) // 修改帐号状态
		adminsM.POST("/role/update", adminHandler.UpdateRole)        // 给用户分配角色
		adminsM.GET("/role/:adminId", adminHandler.GetRole)          // 获取指定用户的角色
	}
}
