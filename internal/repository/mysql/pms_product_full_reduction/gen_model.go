package pms_product_full_reduction

// PmsProductFullReduction 产品满减表(只针对同商品)
//
//go:generate gormgen -structs PmsProductFullReduction -input .
type PmsProductFullReduction struct {
	Id          int64   //
	ProductId   int64   //
	FullPrice   float64 //
	ReducePrice float64 //
}
