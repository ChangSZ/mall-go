package dto

type CartPromotionItem struct {
	OmsCartItem      `json:",inline"`
	PromotionMessage string  `json:"promotionMessage"` // 促销活动信息
	ReduceAmount     float64 `json:"reduceAmount"`     // 促销活动减去的金额，针对每个商品
	RealStock        int32   `json:"realStock"`        // 剩余库存-锁定库存
	Integration      int32   `json:"integration"`      // 购买商品赠送积分
	Growth           int32   `json:"growth"`           // 购买商品赠送成长值
}

// CartProduct 购物车中带规格和SKU的商品信息
type CartProduct struct {
	PmsProduct           `json:",inline"`
	ProductAttributeList []PmsProductAttribute `json:"productAttributeList" gorm:"foreignKey:ProductAttributeCategoryId"` // 商品属性列表
	SkuStockList         []PmsSkuStock         `json:"skuStockList" gorm:"foreignKey:ProductId"`                          // 商品SKU库存列表
}
