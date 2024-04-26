package pms_product_category_attribute_relation

// PmsProductCategoryAttributeRelation 产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）
//
//go:generate gormgen -structs PmsProductCategoryAttributeRelation -input .
type PmsProductCategoryAttributeRelation struct {
	Id                 int64 //
	ProductCategoryId  int64 //
	ProductAttributeId int64 //
}
