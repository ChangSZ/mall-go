package ums_member_receive_address

// UmsMemberReceiveAddress 会员收货地址表
//
//go:generate gormgen -structs UmsMemberReceiveAddress -input .
type UmsMemberReceiveAddress struct {
	Id            int64  //
	MemberId      int64  //
	Name          string // 收货人名称
	PhoneNumber   string //
	DefaultStatus int32  // 是否为默认
	PostCode      string // 邮政编码
	Province      string // 省份/直辖市
	City          string // 城市
	Region        string // 区
	DetailAddress string // 详细地址(街道)
}
