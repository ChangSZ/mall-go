package cms_topic_category

// CmsTopicCategory 话题分类表
//
//go:generate gormgen -structs CmsTopicCategory -input .
type CmsTopicCategory struct {
	Id           int64  //
	Name         string //
	Icon         string // 分类图标
	SubjectCount int32  // 专题数量
	ShowStatus   int32  //
	Sort         int32  //
}
