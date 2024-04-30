package dto

type PmsProductAttributeValue struct {
	Id                 int64  `json:"id"`                 //
	ProductId          int64  `json:"productId"`          //
	ProductAttributeId int64  `json:"productAttributeId"` //
	Value              string `json:"value"`              // 手动添加规格或参数的值，参数单值，规格有多个时以逗号隔开
}

type PmsProductAttr struct {
	Id                         int64  `json:"id"`
	ProductAttributeCategoryId int64  `json:"productAttributeCategoryId"`
	Name                       string `json:"name"`
	SelectType                 int32  `json:"selectType"`
	InputType                  int32  `json:"inputType"`
	InputList                  string `json:"inputList"`
	Sort                       int32  `json:"sort"`
	FilterType                 int32  `json:"filterType"`
	SearchType                 int32  `json:"searchType"`
	RelatedStatus              int32  `json:"relatedStatus"`
	HandAddStatus              int32  `json:"handAddStatus"`
	Type                       int32  `json:"type"`
}
