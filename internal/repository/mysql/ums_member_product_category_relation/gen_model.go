package ums_member_product_category_relation

// UmsMemberProductCategoryRelation 会员与产品分类关系表（用户喜欢的分类）
//
//go:generate gormgen -structs UmsMemberProductCategoryRelation -input .
type UmsMemberProductCategoryRelation struct {
	Id                int64 //
	MemberId          int64 //
	ProductCategoryId int64 //
}
