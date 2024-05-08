package dto

type SmsCouponProductRelation struct {
	Id          int64  `json:"id"`
	CouponId    int64  `json:"couponId"`
	ProductId   int64  `json:"productId"`
	ProductName string `json:"productName"`
	ProductSn   string `json:"productSn"`
}
