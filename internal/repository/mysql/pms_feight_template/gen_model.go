package pms_feight_template

// PmsFeightTemplate 运费模版
//
//go:generate gormgen -structs PmsFeightTemplate -input .
type PmsFeightTemplate struct {
	Id             int64   //
	Name           string  //
	ChargeType     int32   // 计费类型:0->按重量；1->按件数
	FirstWeight    float64 // 首重kg
	FirstFee       float64 // 首费（元）
	ContinueWeight float64 //
	ContinmeFee    float64 //
	Dest           string  // 目的地（省、市）
}
