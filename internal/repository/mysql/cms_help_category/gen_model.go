package cms_help_category

// CmsHelpCategory 帮助分类表
//
//go:generate gormgen -structs CmsHelpCategory -input .
type CmsHelpCategory struct {
	Id         int64  //
	Name       string //
	Icon       string // 分类图标
	HelpCount  int32  // 专题数量
	ShowStatus int32  //
	Sort       int32  //
}
