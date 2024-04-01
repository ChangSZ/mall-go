package router

import "github.com/ChangSZ/mall-go/internal/api/ums_member_level"

// 会员等级管理
func setUmsMemberLevelRouter(r *resource) {
	memberLevelHandler := ums_member_level.New(r.logger, r.db, r.cache)
	memberLevel := r.mux.Group("/memberLevel")
	{
		memberLevel.GET("/list", memberLevelHandler.List()) // 查询所有会员等级
	}
}
