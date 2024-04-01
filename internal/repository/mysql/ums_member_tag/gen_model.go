package ums_member_tag

// UmsMemberTag 用户标签表
//
//go:generate gormgen -structs UmsMemberTag -input .
type UmsMemberTag struct {
	Id                int64   //
	Name              string  //
	FinishOrderCount  int32   // 自动打标签完成订单数量
	FinishOrderAmount float64 // 自动打标签完成订单金额
}
