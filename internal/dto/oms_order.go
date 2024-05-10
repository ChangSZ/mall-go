package dto

import (
	"time"
)

type OmsOrder struct {
	Id                    int64     `json:"id"`                    // 订单id
	MemberId              int64     `json:"memberId"`              //
	CouponId              int64     `json:"couponId"`              //
	OrderSn               string    `json:"orderSn"`               // 订单编号
	CreateTime            time.Time `json:"createTime"`            // 提交时间
	MemberUsername        string    `json:"memberUsername"`        // 用户帐号
	TotalAmount           float64   `json:"totalAmount"`           // 订单总金额
	PayAmount             float64   `json:"payAmount"`             // 应付金额（实际支付金额）
	FreightAmount         float64   `json:"freightAmount"`         // 运费金额
	PromotionAmount       float64   `json:"promotionAmount"`       // 促销优化金额（促销价、满减、阶梯价）
	IntegrationAmount     float64   `json:"integrationAmount"`     // 积分抵扣金额
	CouponAmount          float64   `json:"couponAmount"`          // 优惠券抵扣金额
	DiscountAmount        float64   `json:"discountAmount"`        // 管理员后台调整订单使用的折扣金额
	PayType               int32     `json:"payType"`               // 支付方式：0->未支付；1->支付宝；2->微信
	SourceType            int32     `json:"sourceType"`            // 订单来源：0->PC订单；1->app订单
	Status                int32     `json:"status"`                // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	OrderType             int32     `json:"orderType"`             // 订单类型：0->正常订单；1->秒杀订单
	DeliveryCompany       string    `json:"deliveryCompany"`       // 物流公司(配送方式)
	DeliverySn            string    `json:"deliverySn"`            // 物流单号
	AutoConfirmDay        int32     `json:"autoConfirmDay"`        // 自动确认时间（天）
	Integration           int32     `json:"integration"`           // 可以获得的积分
	Growth                int32     `json:"growth"`                // 可以活动的成长值
	PromotionInfo         string    `json:"promotionInfo"`         // 活动信息
	BillType              int32     `json:"billType"`              // 发票类型：0->不开发票；1->电子发票；2->纸质发票
	BillHeader            string    `json:"billHeader"`            // 发票抬头
	BillContent           string    `json:"billContent"`           // 发票内容
	BillReceiverPhone     string    `json:"billReceiverPhone"`     // 收票人电话
	BillReceiverEmail     string    `json:"billReceiverEmail"`     // 收票人邮箱
	ReceiverName          string    `json:"receiverName"`          // 收货人姓名
	ReceiverPhone         string    `json:"receiverPhone"`         // 收货人电话
	ReceiverPostCode      string    `json:"receiverPostCode"`      // 收货人邮编
	ReceiverProvince      string    `json:"receiverProvince"`      // 省份/直辖市
	ReceiverCity          string    `json:"receiverCity"`          // 城市
	ReceiverRegion        string    `json:"receiverRegion"`        // 区
	ReceiverDetailAddress string    `json:"receiverDetailAddress"` // 详细地址
	Note                  string    `json:"note"`                  // 订单备注
	ConfirmStatus         int32     `json:"confirmStatus"`         // 确认收货状态：0->未确认；1->已确认
	DeleteStatus          int32     `json:"deleteStatus"`          // 删除状态：0->未删除；1->已删除
	UseIntegration        int32     `json:"useIntegration"`        // 下单时使用的积分
	PaymentTime           time.Time `json:"paymentTime"`           // 支付时间
	DeliveryTime          time.Time `json:"deliveryTime"`          // 发货时间
	ReceiveTime           time.Time `json:"receiveTime"`           // 确认收货时间
	CommentTime           time.Time `json:"commentTime"`           // 评价时间
	ModifyTime            time.Time `json:"modifyTime"`            // 修改时间
}

type OmsOrderQueryParam struct {
	OrderSn         string `form:"orderSn"`
	ReceiverKeyword string `form:"receiverKeyword"`
	Status          *int32 `form:"status"`
	OrderType       *int32 `form:"orderType"`
	SourceType      *int32 `form:"sourceType"`
	CreateTime      string `form:"createTime"`
}

type OmsOrderDeliveryParam struct {
	OrderId         int64  `json:"orderId"`         // 订单ID
	DeliveryCompany string `json:"deliveryCompany"` // 物流公司
	DeliverySn      string `json:"deliverySn"`      // 物流ID
}

// OmsOrderDetail 订单详情信息
type OmsOrderDetail struct {
	OmsOrder      `json:",inline"`
	OrderItemList []OmsOrderItem           `json:"orderItemList"  gorm:"foreignKey:OrderId"` // 订单商品列表
	HistoryList   []OmsOrderOperateHistory `json:"historyList"  gorm:"foreignKey:OrderId"`   // 订单操作记录列表
}

// OmsReceiverInfoParam 订单修改收货人信息参数
type OmsReceiverInfoParam struct {
	OrderID               int64  `json:"orderId"`               // 订单ID
	ReceiverName          string `json:"receiverName"`          // 收货人姓名
	ReceiverPhone         string `json:"receiverPhone"`         // 收货人电话
	ReceiverPostCode      string `json:"receiverPostCode"`      // 收货人邮编
	ReceiverDetailAddress string `json:"receiverDetailAddress"` // 详细地址
	ReceiverProvince      string `json:"receiverProvince"`      // 省份/直辖市
	ReceiverCity          string `json:"receiverCity"`          // 城市
	ReceiverRegion        string `json:"receiverRegion"`        // 区
	Status                int32  `json:"status"`                // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
}

// OmsMoneyInfoParam 修改订单费用信息参数
type OmsMoneyInfoParam struct {
	OrderID        int64   `json:"orderId"`        // 订单ID
	FreightAmount  float64 `json:"freightAmount"`  // 运费金额
	DiscountAmount float64 `json:"discountAmount"` // 管理员后台调整订单所使用的折扣金额
	Status         int32   `json:"status"`         // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
}
