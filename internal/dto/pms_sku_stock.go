package dto

type PmsSkuStock struct {
	Id             int64   `json:"id"`             //
	ProductId      int64   `json:"productId"`      //
	SkuCode        string  `json:"skuCode"`        // sku编码
	Price          float64 `json:"price"`          //
	Stock          int32   `json:"stock"`          // 库存
	LowStock       int32   `json:"lowStock"`       // 预警库存
	Pic            string  `json:"pic"`            // 展示图片
	Sale           int32   `json:"sale"`           // 销量
	PromotionPrice float64 `json:"promotionPrice"` // 单品促销价格
	LockStock      int32   `json:"lockStock"`      // 锁定库存
	SpData         string  `json:"spData"`         // 商品销售属性，json格式
}

type PmsPidUri struct {
	Pid int64 `uri:"pid"`
}
