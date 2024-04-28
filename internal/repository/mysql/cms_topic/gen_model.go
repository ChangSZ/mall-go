package cms_topic

import "time"

// CmsTopic 话题表
//
//go:generate gormgen -structs CmsTopic -input .
type CmsTopic struct {
	Id             int64     //
	CategoryId     int64     //
	Name           string    //
	CreateTime     time.Time `gorm:"autoCreateTime"` //
	StartTime      time.Time `gorm:"time"`           //
	EndTime        time.Time `gorm:"time"`           //
	AttendCount    int32     // 参与人数
	AttentionCount int32     // 关注人数
	ReadCount      int32     //
	AwardName      string    // 奖品名称
	AttendType     string    // 参与方式
	Content        string    // 话题内容
}
