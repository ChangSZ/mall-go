package pms_sku_stock

// PmsSkuStock sku的库存
//
//go:generate gormgen -structs PmsSkuStock -input .
type PmsSkuStock struct {
	Id             int64   //
	ProductId      int64   //
	SkuCode        string  // sku编码
	Price          float64 //
	Stock          int32   // 库存
	LowStock       int32   // 预警库存
	Pic            string  // 展示图片
	Sale           int32   // 销量
	PromotionPrice float64 // 单品促销价格
	LockStock      int32   // 锁定库存
	SpData         string  // 商品销售属性，json格式
}
