package sms_home_recommend_product

// SmsHomeRecommendProduct 人气推荐商品表
//
//go:generate gormgen -structs SmsHomeRecommendProduct -input .
type SmsHomeRecommendProduct struct {
	Id              int64  //
	ProductId       int64  //
	ProductName     string //
	RecommendStatus int32  //
	Sort            int32  //
}
