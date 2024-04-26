package pms_product_ladder

// PmsProductLadder 产品阶梯价格表(只针对同商品)
//
//go:generate gormgen -structs PmsProductLadder -input .
type PmsProductLadder struct {
	Id        int64   //
	ProductId int64   //
	Count     int32   // 满足的商品数量
	Discount  float64 // 折扣
	Price     float64 // 折后价格
}
