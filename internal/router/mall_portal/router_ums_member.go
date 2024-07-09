package mall_portal

import (
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api/mall_portal/ums_member"
	"github.com/ChangSZ/mall-go/internal/middleware"
)

// 会员管理
func setUmsMemberRouter(eng *gin.Engine) {
	handler := ums_member.New()
	group := eng.Group("/sso")
	{
		group.POST("/register", handler.Register) // 会员注册
		group.POST("/login", handler.Login)       // 会员登录
	}

	groupM := eng.Group("/sso", middleware.CheckMemberToken())
	{
		groupM.GET("/info", handler.Info)                      // 获取会员信息
		groupM.GET("/getAuthCode", handler.GetAuthCode)        // 获取验证码
		groupM.POST("/updatePassword", handler.UpdatePassword) // 会员修改密码
		groupM.GET("/refreshToken", handler.RefreshToken)      // 刷新token
	}
}
