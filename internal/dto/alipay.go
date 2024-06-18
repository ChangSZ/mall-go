package dto

// AliPayParam 支付宝支付请求参数
type AliPayParam struct {
	OutTradeNo  string  `form:"outTradeNo"`  // 商户订单号，商家自定义，保持唯一性
	Subject     string  `form:"subject"`     // 商品的标题/交易标题/订单标题/订单关键字等
	TotalAmount float64 `form:"totalAmount"` // 订单总金额，单位为元，精确到小数点后两位
}
