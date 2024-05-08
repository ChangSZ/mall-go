package dto

type SmsHomeBrand struct {
	Id              int64  `json:"id"`
	BrandId         int64  `json:"brandId"`
	BrandName       string `json:"brandName"`
	RecommendStatus int32  `json:"recommendStatus"`
	Sort            int32  `json:"sort"`
}
