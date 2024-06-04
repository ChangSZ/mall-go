package dto

import "time"

// HomeContentResult 首页内容返回信息
type HomeContentResult struct {
	AdvertiseList      []SmsHomeAdvertise `json:"advertiseList"`      // 轮播广告
	BrandList          []PmsBrand         `json:"brandList"`          // 推荐品牌
	HomeFlashPromotion HomeFlashPromotion `json:"homeFlashPromotion"` // 当前秒杀场次
	NewProductList     []PmsProduct       `json:"newProductList"`     // 新品推荐
	HotProductList     []PmsProduct       `json:"hotProductList"`     // 人气推荐
	SubjectList        []CmsSubject       `json:"subjectList"`        // 推荐专题
}

// HomeFlashPromotion 首页秒杀场次
type HomeFlashPromotion struct {
	StartTime     time.Time               `json:"startTime"`     // 本场开始时间
	EndTime       time.Time               `json:"endTime"`       // 本场结束时间
	NextStartTime time.Time               `json:"nextStartTime"` // 下场开始时间
	NextEndTime   time.Time               `json:"nextEndTime"`   // 下场结束时间
	ProductList   []FlashPromotionProduct `json:"productList"`   // 属于该秒杀活动的商品
}

// FlashPromotionProduct 秒杀信息和商品
type FlashPromotionProduct struct {
	PmsProduct          `json:",inline"`
	FlashPromotionPrice float64 `json:"flashPromotionPrice"` // 秒杀价格
	FlashPromotionCount int32   `json:"flashPromotionCount"` // 用于秒杀到数量
	FlashPromotionLimit int32   `json:"flashPromotionLimit"` // 秒杀限购数量
}
