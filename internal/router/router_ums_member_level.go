package router

import (
	"github.com/ChangSZ/mall-go/internal/api/ums_member_level"

	"github.com/gin-gonic/gin"
)

// 会员等级管理
func setUmsMemberLevelRouter(eng *gin.Engine) {
	memberLevelHandler := ums_member_level.New()
	memberLevel := eng.Group("/memberLevel")
	{
		memberLevel.GET("/list", memberLevelHandler.List) // 查询所有会员等级
	}
}
