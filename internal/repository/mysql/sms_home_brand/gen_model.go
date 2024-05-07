package sms_home_brand

// SmsHomeBrand 首页推荐品牌表
//
//go:generate gormgen -structs SmsHomeBrand -input .
type SmsHomeBrand struct {
	Id              int64  //
	BrandId         int64  //
	BrandName       string //
	RecommendStatus int32  //
	Sort            int32  //
}
