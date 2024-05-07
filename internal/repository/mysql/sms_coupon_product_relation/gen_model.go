package sms_coupon_product_relation

// SmsCouponProductRelation 优惠券和产品的关系表
//
//go:generate gormgen -structs SmsCouponProductRelation -input .
type SmsCouponProductRelation struct {
	Id          int64  //
	CouponId    int64  //
	ProductId   int64  //
	ProductName string // 商品名称
	ProductSn   string // 商品编码
}
