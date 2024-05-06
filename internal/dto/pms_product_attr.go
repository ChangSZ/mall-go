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

type PmsProductAttrParam struct {
	ProductAttributeCategoryId int64  `json:"productAttributeCategoryId" binding:"required"` // 属性分类ID
	Name                       string `json:"name" binding:"required"`                       // 属性名称
	SelectType                 int32  `json:"selectType" binding:"omitempty,oneof=0 1 2"`    // 属性选择类型：0->唯一；1->单选；2->多选
	InputType                  int32  `json:"inputType" binding:"omitempty,oneof=0 1"`       // 属性录入方式：0->手工录入；1->从列表中选取
	InputList                  string `json:"inputList"`                                     // 可选值列表，以逗号隔开
	Sort                       int32  `json:"sort"`                                          // 排序
	FilterType                 int32  `json:"filterType" binding:"omitempty,oneof=0 1"`      // 分类筛选样式：0->普通；1->颜色
	SearchType                 int32  `json:"searchType" binding:"omitempty,oneof=0 1 2"`    // 检索类型；0->不需要进行检索；1->关键字检索；2->范围检索
	RelatedStatus              int32  `json:"relatedStatus" binding:"omitempty,oneof=0 1"`   // 相同属性商品是否关联；0->不关联；1->关联
	HandAddStatus              int32  `json:"handAddStatus" binding:"omitempty,oneof=0 1"`   // 是否支持手动新增；0->不支持；1->支持
	Type                       int32  `json:"type" binding:"omitempty,oneof=0 1"`            // 属性的类型；0->规格；1->参数
}

type PmsProductAttrInfo struct {
	AttributeId         int64 `json:"attributeId"`         // 商品属性ID
	AttributeCategoryId int64 `json:"attributeCategoryId"` // 商品属性分类ID
}

type PmsProductAttrUri struct {
	Cid int64 `uri:"cid" binding:"required"`
}

type PmsProductCateIdUri struct {
	ProductCategoryId int64 `uri:"productCategoryId" binding:"required"`
}
