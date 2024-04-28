package cms_subject_comment

import "time"

// CmsSubjectComment 专题评论表
//
//go:generate gormgen -structs CmsSubjectComment -input .
type CmsSubjectComment struct {
	Id             int64     //
	SubjectId      int64     //
	MemberNickName string    //
	MemberIcon     string    //
	Content        string    //
	CreateTime     time.Time `gorm:"autoCreateTime"` //
	ShowStatus     int32     //
}
