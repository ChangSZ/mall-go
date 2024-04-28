package cms_subject_product_relation

// CmsSubjectProductRelation 专题商品关系表
//
//go:generate gormgen -structs CmsSubjectProductRelation -input .
type CmsSubjectProductRelation struct {
	Id        int64 //
	SubjectId int64 //
	ProductId int64 //
}
