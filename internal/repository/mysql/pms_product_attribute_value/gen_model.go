package pms_product_attribute_value

// PmsProductAttributeValue 存储产品参数信息的表
//
//go:generate gormgen -structs PmsProductAttributeValue -input .
type PmsProductAttributeValue struct {
	Id                 int64  //
	ProductId          int64  //
	ProductAttributeId int64  //
	Value              string // 手动添加规格或参数的值，参数单值，规格有多个时以逗号隔开
}
