package dto

// PromotionProduct 促销商品信息，包括sku、打折优惠、满减优惠
type PromotionProduct struct {
	PmsProduct               `json:",inline"`
	SkuStockList             []PmsSkuStock             `json:"skuStockList" gorm:"foreignKey:ProductId"`             // 商品库存信息
	ProductLadderList        []PmsProductLadder        `json:"productLadderList" gorm:"foreignKey:ProductId"`        // 商品打折信息
	ProductFullReductionList []PmsProductFullReduction `json:"productFullReductionList" gorm:"foreignKey:ProductId"` // 商品满减信息
}
