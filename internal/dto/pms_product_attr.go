package dto

type PmsProductAttributeValue struct {
	Id                 int64  `json:"id"`                 //
	ProductId          int64  `json:"productId"`          //
	ProductAttributeId int64  `json:"productAttributeId"` //
	Value              string `json:"value"`              // 手动添加规格或参数的值，参数单值，规格有多个时以逗号隔开
}
