package mall_admin

import (
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api/mall_admin/ums_member_level"
	"github.com/ChangSZ/mall-go/internal/middleware"
)

// 会员等级管理
func setUmsMemberLevelRouter(eng *gin.Engine) {
	memberLevelHandler := ums_member_level.New()
	memberLevel := eng.Group("/memberLevel", middleware.CheckToken(), middleware.DynamicAccess())
	{
		memberLevel.GET("/list", memberLevelHandler.List) // 查询所有会员等级
	}
}
