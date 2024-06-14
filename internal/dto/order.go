package dto

// OmsOrderReturnApplyParam 退货申请请求参数
type OmsOrderReturnApplyParam struct {
	OrderId          int64   `json:"orderId"`          // 订单id
	ProductId        int64   `json:"productId"`        // 退货商品id
	OrderSn          string  `json:"orderSn"`          // 订单编号
	MemberUsername   string  `json:"memberUsername"`   // 会员用户名
	ReturnName       string  `json:"returnName"`       // 退货人姓名
	ReturnPhone      string  `json:"returnPhone"`      // 退货人电话
	ProductPic       string  `json:"productPic"`       // 商品图片
	ProductName      string  `json:"productName"`      // 商品名称
	ProductBrand     string  `json:"productBrand"`     // 商品品牌
	ProductAttr      string  `json:"productAttr"`      // 商品销售属性：颜色：红色；尺码：xl
	ProductCount     int     `json:"productCount"`     // 退货数量
	ProductPrice     float64 `json:"productPrice"`     // 商品单价
	ProductRealPrice float64 `json:"productRealPrice"` // 商品实际支付单价
	Reason           string  `json:"reason"`           // 原因
	Description      string  `json:"description"`      // 描述
	ProofPics        string  `json:"proofPics"`        // 凭证图片，以逗号隔开
}

// ConfirmOrderResult 确认单信息
type ConfirmOrderResult struct {
	CartPromotionItemList     []CartPromotionItem          `json:"cartPromotionItemList"`     // 包含优惠信息的购物车信息
	MemberReceiveAddressList  []UmsMemberReceiveAddress    `json:"memberReceiveAddressList"`  // 用户收货地址列表
	CouponHistoryDetailList   []SmsCouponHistoryDetail     `json:"couponHistoryDetailList"`   // 用户可用优惠券列表
	IntegrationConsumeSetting UmsIntegrationConsumeSetting `json:"integrationConsumeSetting"` // 积分使用规则
	MemberIntegration         int32                        `json:"memberIntegration"`         // 会员持有的积分
	CalcAmount                CalcAmount                   `json:"calcAmount"`                // 计算的金额
}

type CalcAmount struct {
	TotalAmount     float64 `json:"totalAmount"`     // 订单商品总金额
	FreightAmount   float64 `json:"freightAmount"`   // 运费
	PromotionAmount float64 `json:"promotionAmount"` // 活动优惠
	PayAmount       float64 `json:"payAmount"`       // 应付金额
}

// OrderParam 生成订单时传入的参数
type OrderParam struct {
	MemberReceiveAddressId int64   `json:"memberReceiveAddressId"` // 收货地址ID
	CouponId               int64   `json:"couponId"`               // 优惠券ID
	UseIntegration         int32   `json:"useIntegration"`         // 使用的积分数
	PayType                int32   `json:"payType"`                // 支付方式
	CartIds                []int64 `json:"cartIds"`                // 被选中的购物车商品ID
}

// OrderDetail 包含商品信息的订单详情
type OrderDetail struct {
	OmsOrder      `json:",inline"`
	OrderItemList []OmsOrderItem `json:"orderItemList"   gorm:"foreignKey:OrderId"` // 订单商品列表
}

type Order struct {
	Order         OmsOrder       `json:"order"`
	OrderItemList []OmsOrderItem `json:"orderItemList"` // 订单商品列表
}
