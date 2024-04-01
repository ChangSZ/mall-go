package ums_member_task

// UmsMemberTask 会员任务表
//
//go:generate gormgen -structs UmsMemberTask -input .
type UmsMemberTask struct {
	Id           int64  //
	Name         string //
	Growth       int32  // 赠送成长值
	Intergration int32  // 赠送积分
	Type         int32  // 任务类型：0->新手任务；1->日常任务
}
