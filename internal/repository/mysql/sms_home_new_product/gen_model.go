package sms_home_new_product

// SmsHomeNewProduct 新鲜好物表
//
//go:generate gormgen -structs SmsHomeNewProduct -input .
type SmsHomeNewProduct struct {
	Id              int64  //
	ProductId       int64  //
	ProductName     string //
	RecommendStatus int32  //
	Sort            int32  //
}
