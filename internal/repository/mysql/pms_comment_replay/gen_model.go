package pms_comment_replay

import "time"

// PmsCommentReplay 产品评价回复表
//
//go:generate gormgen -structs PmsCommentReplay -input .
type PmsCommentReplay struct {
	Id             int64     //
	CommentId      int64     //
	MemberNickName string    //
	MemberIcon     string    //
	Content        string    //
	CreateTime     time.Time `gorm:"autoCreateTime"` //
	Type           int32     // 评论人员类型；0->会员；1->管理员
}
