package oms_company_address

// OmsCompanyAddress 公司收发货地址表
//
//go:generate gormgen -structs OmsCompanyAddress -input .
type OmsCompanyAddress struct {
	Id            int64  //
	AddressName   string // 地址名称
	SendStatus    int32  // 默认发货地址：0->否；1->是
	ReceiveStatus int32  // 是否默认收货地址：0->否；1->是
	Name          string // 收发货人姓名
	Phone         string // 收货人电话
	Province      string // 省/直辖市
	City          string // 市
	Region        string // 区
	DetailAddress string // 详细地址
}
