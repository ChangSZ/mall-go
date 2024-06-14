package oms_order_return_apply

import "time"

// OmsOrderReturnApply 订单退货申请
//
//go:generate gormgen -structs OmsOrderReturnApply -input .
type OmsOrderReturnApply struct {
	Id               int64     //
	OrderId          int64     // 订单id
	CompanyAddressId int64     // 收货地址表id
	ProductId        int64     // 退货商品id
	OrderSn          string    // 订单编号
	CreateTime       time.Time `gorm:"autoCreateTime"` // 申请时间
	MemberUsername   string    // 会员用户名
	ReturnAmount     float64   // 退款金额
	ReturnName       string    // 退货人姓名
	ReturnPhone      string    // 退货人电话
	Status           int32     // 申请状态：0->待处理；1->退货中；2->已完成；3->已拒绝
	HandleTime       time.Time `gorm:"type:time;default:'1000-01-01 00:00:00'"` // 处理时间
	ProductPic       string    // 商品图片
	ProductName      string    // 商品名称
	ProductBrand     string    // 商品品牌
	ProductAttr      string    // 商品销售属性：颜色：红色；尺码：xl;
	ProductCount     int32     // 退货数量
	ProductPrice     float64   // 商品单价
	ProductRealPrice float64   // 商品实际支付单价
	Reason           string    // 原因
	Description      string    // 描述
	ProofPics        string    // 凭证图片，以逗号隔开
	HandleNote       string    // 处理备注
	HandleMan        string    // 处理人员
	ReceiveMan       string    // 收货人
	ReceiveTime      time.Time `gorm:"type:time;default:'1000-01-01 00:00:00'"` // 收货时间
	ReceiveNote      string    // 收货备注
}
