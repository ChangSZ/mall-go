package sms_coupon_product_category_relation

// SmsCouponProductCategoryRelation 优惠券和产品分类关系表
//
//go:generate gormgen -structs SmsCouponProductCategoryRelation -input .
type SmsCouponProductCategoryRelation struct {
	Id                  int64  //
	CouponId            int64  //
	ProductCategoryId   int64  //
	ProductCategoryName string // 产品分类名称
	ParentCategoryName  string // 父分类名称
}
