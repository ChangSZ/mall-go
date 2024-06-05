package dto

type PmsProductAttributeCategory struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	AttributeCount int32  `json:"attributeCount"`
	ParamCount     int32  `json:"paramCount"`
}

type PmsProductAttrCateItem struct {
	PmsProductAttributeCategory `json:",inline"`
	ProductAttributeList        []PmsProductAttribute `json:"productAttributeList"`
}
