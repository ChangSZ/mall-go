package cms_subject_category

// CmsSubjectCategory 专题分类表
//
//go:generate gormgen -structs CmsSubjectCategory -input .
type CmsSubjectCategory struct {
	Id           int64  //
	Name         string //
	Icon         string // 分类图标
	SubjectCount int32  // 专题数量
	ShowStatus   int32  //
	Sort         int32  //
}
