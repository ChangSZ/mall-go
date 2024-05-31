package dto

type SmsCouponHistoryDetail struct {
	SmsCouponHistory     `json:",inline"`
	Coupon               SmsCoupon                          `json:"coupon"  gorm:"embedded;embeddedPrefix:c_"`       // 相关优惠券信息
	ProductRelationList  []SmsCouponProductRelation         `json:"productRelationList" gorm:"foreignKey:CouponId"`  // 优惠券关联商品
	CategoryRelationList []SmsCouponProductCategoryRelation `json:"categoryRelationList" gorm:"foreignKey:CouponId"` // 优惠券关联商品分类
}
