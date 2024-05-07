package sms_flash_promotion_product_relation

// SmsFlashPromotionProductRelation 商品限时购与商品关系表
//
//go:generate gormgen -structs SmsFlashPromotionProductRelation -input .
type SmsFlashPromotionProductRelation struct {
	Id                      int64   // 编号
	FlashPromotionId        int64   //
	FlashPromotionSessionId int64   // 编号
	ProductId               int64   //
	FlashPromotionPrice     float64 // 限时购价格
	FlashPromotionCount     int32   // 限时购数量
	FlashPromotionLimit     int32   // 每人限购数量
	Sort                    int32   // 排序
}
