package cms_help

import "time"

// CmsHelp 帮助表
//
//go:generate gormgen -structs CmsHelp -input .
type CmsHelp struct {
	Id         int64     //
	CategoryId int64     //
	Icon       string    //
	Title      string    //
	ShowStatus int32     //
	CreateTime time.Time `gorm:"autoCreateTime"` //
	ReadCount  int32     //
	Content    string    //
}
