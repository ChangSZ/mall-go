package dto

type SmsFlashPromotionProductRelation struct {
	Id                      int64   `json:"id"`
	FlashPromotionId        int64   `json:"flashPromotionId"`
	FlashPromotionSessionId int64   `json:"flashPromotionSessionId"`
	ProductId               int64   `json:"productId"`
	FlashPromotionPrice     float64 `json:"flashPromotionPrice"`
	FlashPromotionCount     int32   `json:"flashPromotionCount"`
	FlashPromotionLimit     int32   `json:"flashPromotionLimit"`
	Sort                    int32   `json:"sort"`
}
