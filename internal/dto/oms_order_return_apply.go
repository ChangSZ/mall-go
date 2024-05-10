package dto

import "time"

type OmsOrderReturnApply struct {
	Id               int64     `json:"id"`               //
	OrderId          int64     `json:"orderId"`          // 订单id
	CompanyAddressId int64     `json:"companyAddressId"` // 收货地址表id
	ProductId        int64     `json:"productId"`        // 退货商品id
	OrderSn          string    `json:"orderSn"`          // 订单编号
	CreateTime       time.Time `json:"createTime"`       // 申请时间
	MemberUsername   string    `json:"memberUsername"`   // 会员用户名
	ReturnAmount     float64   `json:"returnAmount"`     // 退款金额
	ReturnName       string    `json:"returnName"`       // 退货人姓名
	ReturnPhone      string    `json:"returnPhone"`      // 退货人电话
	Status           int32     `json:"status"`           // 申请状态：0->待处理；1->退货中；2->已完成；3->已拒绝
	HandleTime       time.Time `json:"handleTime"`       // 处理时间
	ProductPic       string    `json:"productPic"`       // 商品图片
	ProductName      string    `json:"productName"`      // 商品名称
	ProductBrand     string    `json:"productBrand"`     // 商品品牌
	ProductAttr      string    `json:"productAttr"`      // 商品销售属性：颜色：红色；尺码：xl;
	ProductCount     int32     `json:"productCount"`     // 退货数量
	ProductPrice     float64   `json:"productPrice"`     // 商品单价
	ProductRealPrice float64   `json:"productRealPrice"` // 商品实际支付单价
	Reason           string    `json:"reason"`           // 原因
	Description      string    `json:"description"`      // 描述
	ProofPics        string    `json:"proofPics"`        // 凭证图片，以逗号隔开
	HandleNote       string    `json:"handleNote"`       // 处理备注
	HandleMan        string    `json:"handleMan"`        // 处理人员
	ReceiveMan       string    `json:"receiveMan"`       // 收货人
	ReceiveTime      time.Time `json:"receiveTime"`      // 收货时间
	ReceiveNote      string    `json:"receiveNote"`      // 收货备注
}

// OmsReturnApplyQueryParam 订单退货申请查询参数
type OmsReturnApplyQueryParam struct {
	ID              int64  `json:"id"`              // 服务单号
	ReceiverKeyword string `json:"receiverKeyword"` // 收货人姓名/号码
	Status          *int32 `json:"status"`          // 申请状态：0->待处理；1->退货中；2->已完成；3->已拒绝
	CreateTime      string `json:"createTime"`      // 申请时间
	HandleMan       string `json:"handleMan"`       // 处理人员
	HandleTime      string `json:"handleTime"`      // 处理时间
}

// OmsUpdateStatusParam 确认收货请求参数
type OmsUpdateStatusParam struct {
	ID               int64   `json:"id"`               // 服务单号
	CompanyAddressID int64   `json:"companyAddressId"` // 收货地址关联id
	ReturnAmount     float64 `json:"returnAmount"`     // 确认退款金额
	HandleNote       string  `json:"handleNote"`       // 处理备注
	HandleMan        string  `json:"handleMan"`        // 处理人
	ReceiveNote      string  `json:"receiveNote"`      // 收货备注
	ReceiveMan       string  `json:"receiveMan"`       // 收货人
	Status           int32   `json:"status"`           // 申请状态：1->退货中；2->已完成；3->已拒绝
}

// OmsOrderReturnApplyResult 申请信息封装
type OmsOrderReturnApplyResult struct {
	OmsOrderReturnApply `json:",inline"`
	CompanyAddress      OmsCompanyAddress `json:"companyAddress" gorm:"embedded;embeddedPrefix:ca_"` // 公司收货地址
}
