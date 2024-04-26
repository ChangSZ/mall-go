package pms_product_attribute

// PmsProductAttribute 商品属性参数表
//
//go:generate gormgen -structs PmsProductAttribute -input .
type PmsProductAttribute struct {
	Id                         int64  //
	ProductAttributeCategoryId int64  //
	Name                       string //
	SelectType                 int32  // 属性选择类型：0->唯一；1->单选；2->多选
	InputType                  int32  // 属性录入方式：0->手工录入；1->从列表中选取
	InputList                  string // 可选值列表，以逗号隔开
	Sort                       int32  // 排序字段：最高的可以单独上传图片
	FilterType                 int32  // 分类筛选样式：1->普通；1->颜色
	SearchType                 int32  // 检索类型；0->不需要进行检索；1->关键字检索；2->范围检索
	RelatedStatus              int32  // 相同属性产品是否关联；0->不关联；1->关联
	HandAddStatus              int32  // 是否支持手动新增；0->不支持；1->支持
	Type                       int32  // 属性的类型；0->规格；1->参数
}
