package dto

type SmsHomeRecommendProduct struct {
	Id              int64  `json:"id"`
	ProductId       int64  `json:"productId"`
	ProductName     string `json:"productName"`
	RecommendStatus int32  `json:"recommendStatus"`
	Sort            int32  `json:"sort"`
}
