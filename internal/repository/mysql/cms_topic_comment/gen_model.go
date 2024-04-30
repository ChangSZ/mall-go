package cms_topic_comment

import "time"

// CmsTopicComment 专题评论表
//
//go:generate gormgen -structs CmsTopicComment -input .
type CmsTopicComment struct {
	Id             int64     //
	MemberNickName string    //
	TopicId        int64     //
	MemberIcon     string    //
	Content        string    //
	CreateTime     time.Time `gorm:"autoCreateTime"` //
	ShowStatus     int32     //
}
