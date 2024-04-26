package pms_member_price

// PmsMemberPrice 商品会员价格表
//
//go:generate gormgen -structs PmsMemberPrice -input .
type PmsMemberPrice struct {
	Id              int64   //
	ProductId       int64   //
	MemberLevelId   int64   //
	MemberPrice     float64 // 会员价格
	MemberLevelName string  //
}
