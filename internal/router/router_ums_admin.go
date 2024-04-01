package router

import "github.com/ChangSZ/mall-go/internal/api/ums_admin"

// 后台用户管理
func setUmsAdminRouter(r *resource) {
	adminHandler := ums_admin.New(r.logger, r.db, r.cache)
	admins := r.mux.Group("/admin")
	{
		admins.POST("/register", adminHandler.Register())             // 用户注册
		admins.POST("/login", adminHandler.Login())                   // 登录以后返回token
		admins.GET("/refreshToken", adminHandler.RefreshToken())      // 刷新token
		admins.GET("/listAll", adminHandler.ListAll())                // 获取所有角色
		admins.GET("/info", adminHandler.Info())                      // 获取当前登录用户信息
		admins.GET("/logout", adminHandler.Logout())                  // 登出功能
		admins.GET("/:id", adminHandler.Get())                        // 获取指定用户信息
		admins.POST("/update/:id", adminHandler.Update())             // 修改指定用户信息
		admins.POST("/updatePassword", adminHandler.UpdatePassword()) // 修改指定用户密码
		admins.POST("/delete/:id", adminHandler.Delete())             // 删除指定用户信息
		admins.POST("/updateStatus/:id", adminHandler.UpdateStatus()) // 修改帐号状态
		admins.POST("/role/update", adminHandler.UpdateRole())        // 给用户分配角色
		admins.GET("/role/:adminId", adminHandler.GetRole())          // 获取指定用户的角色
	}
}
