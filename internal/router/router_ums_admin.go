package router

import (
	"github.com/ChangSZ/mall-go/internal/api/ums_admin"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

// 后台用户管理
func setUmsAdminRouter(r *resource) {
	ums_admin.InitUmsUserService(r.db, r.cache)
	adminHandler := ums_admin.New(r.logger, r.db, r.cache)
	admins := r.mux.Group("/admin")
	{
		admins.POST("/register", adminHandler.Register())                                                               // 用户注册
		admins.POST("/login", adminHandler.Login())                                                                     // 登录以后返回token
		admins.GET("/refreshToken", core.WrapTokenHandler(r.interceptors.CheckToken), adminHandler.RefreshToken())      // 刷新token
		admins.GET("/info", core.WrapTokenHandler(r.interceptors.CheckToken), adminHandler.Info())                      // 获取当前登录用户信息
		admins.GET("/logout", core.WrapTokenHandler(r.interceptors.CheckToken), adminHandler.Logout())                  // 登出功能
		admins.GET("/:id", core.WrapTokenHandler(r.interceptors.CheckToken), adminHandler.Get())                        // 获取指定用户信息
		admins.POST("/update/:id", core.WrapTokenHandler(r.interceptors.CheckToken), adminHandler.Update())             // 修改指定用户信息
		admins.POST("/updatePassword", core.WrapTokenHandler(r.interceptors.CheckToken), adminHandler.UpdatePassword()) // 修改指定用户密码
		admins.POST("/delete/:id", core.WrapTokenHandler(r.interceptors.CheckToken), adminHandler.Delete())             // 删除指定用户信息
		admins.POST("/updateStatus/:id", core.WrapTokenHandler(r.interceptors.CheckToken), adminHandler.UpdateStatus()) // 修改帐号状态
		admins.POST("/role/update", core.WrapTokenHandler(r.interceptors.CheckToken), adminHandler.UpdateRole())        // 给用户分配角色
		admins.GET("/role/:adminId", core.WrapTokenHandler(r.interceptors.CheckToken), adminHandler.GetRole())          // 获取指定用户的角色
	}
}
