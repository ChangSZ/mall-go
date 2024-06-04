package ums_growth_change_history

import "time"

// UmsGrowthChangeHistory 成长值变化历史记录表
//
//go:generate gormgen -structs UmsGrowthChangeHistory -input .
type UmsGrowthChangeHistory struct {
	Id          int64     //
	MemberId    int64     //
	CreateTime  time.Time `gorm:"autoCreateTime"` //
	ChangeType  int32     // 改变类型：0->增加；1->减少
	ChangeCount int32     // 积分改变数量
	OperateMan  string    // 操作人员
	OperateNote string    // 操作备注
	SourceType  int32     // 积分来源：0->购物；1->管理员修改
}
