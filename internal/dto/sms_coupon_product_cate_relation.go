package dto

type SmsCouponProductCategoryRelation struct {
	Id                  int64  `json:"id"`
	CouponId            int64  `json:"couponId"`
	ProductCategoryId   int64  `json:"productCategoryId"`
	ProductCategoryName string `json:"productCategoryName"`
	ParentCategoryName  string `json:"parentCategoryName"`
}
