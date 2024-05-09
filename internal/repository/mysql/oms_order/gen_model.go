package oms_order

import "time"

// OmsOrder 订单表
//
//go:generate gormgen -structs OmsOrder -input .
type OmsOrder struct {
	Id                    int64     // 订单id
	MemberId              int64     //
	CouponId              int64     //
	OrderSn               string    // 订单编号
	CreateTime            time.Time `gorm:"autoCreateTime"` // 提交时间
	MemberUsername        string    // 用户帐号
	TotalAmount           float64   // 订单总金额
	PayAmount             float64   // 应付金额（实际支付金额）
	FreightAmount         float64   // 运费金额
	PromotionAmount       float64   // 促销优化金额（促销价、满减、阶梯价）
	IntegrationAmount     float64   // 积分抵扣金额
	CouponAmount          float64   // 优惠券抵扣金额
	DiscountAmount        float64   // 管理员后台调整订单使用的折扣金额
	PayType               int32     // 支付方式：0->未支付；1->支付宝；2->微信
	SourceType            int32     // 订单来源：0->PC订单；1->app订单
	Status                int32     // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	OrderType             int32     // 订单类型：0->正常订单；1->秒杀订单
	DeliveryCompany       string    // 物流公司(配送方式)
	DeliverySn            string    // 物流单号
	AutoConfirmDay        int32     // 自动确认时间（天）
	Integration           int32     // 可以获得的积分
	Growth                int32     // 可以活动的成长值
	PromotionInfo         string    // 活动信息
	BillType              int32     // 发票类型：0->不开发票；1->电子发票；2->纸质发票
	BillHeader            string    // 发票抬头
	BillContent           string    // 发票内容
	BillReceiverPhone     string    // 收票人电话
	BillReceiverEmail     string    // 收票人邮箱
	ReceiverName          string    // 收货人姓名
	ReceiverPhone         string    // 收货人电话
	ReceiverPostCode      string    // 收货人邮编
	ReceiverProvince      string    // 省份/直辖市
	ReceiverCity          string    // 城市
	ReceiverRegion        string    // 区
	ReceiverDetailAddress string    // 详细地址
	Note                  string    // 订单备注
	ConfirmStatus         int32     // 确认收货状态：0->未确认；1->已确认
	DeleteStatus          int32     // 删除状态：0->未删除；1->已删除
	UseIntegration        int32     // 下单时使用的积分
	PaymentTime           time.Time `gorm:"time"` // 支付时间
	DeliveryTime          time.Time `gorm:"time"` // 发货时间
	ReceiveTime           time.Time `gorm:"time"` // 确认收货时间
	CommentTime           time.Time `gorm:"time"` // 评价时间
	ModifyTime            time.Time `gorm:"time"` // 修改时间
}
