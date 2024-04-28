package cms_prefrence_area_product_relation

// CmsPrefrenceAreaProductRelation 优选专区和产品关系表
//
//go:generate gormgen -structs CmsPrefrenceAreaProductRelation -input .
type CmsPrefrenceAreaProductRelation struct {
	Id              int64 //
	PrefrenceAreaId int64 //
	ProductId       int64 //
}
