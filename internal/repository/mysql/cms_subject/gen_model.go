package cms_subject

import "time"

// CmsSubject 专题表
//
//go:generate gormgen -structs CmsSubject -input .
type CmsSubject struct {
	Id              int64     //
	CategoryId      int64     //
	Title           string    //
	Pic             string    // 专题主图
	ProductCount    int32     // 关联产品数量
	RecommendStatus int32     //
	CreateTime      time.Time `gorm:"autoCreateTime"` //
	CollectCount    int32     //
	ReadCount       int32     //
	CommentCount    int32     //
	AlbumPics       string    // 画册图片用逗号分割
	Description     string    //
	ShowStatus      int32     // 显示状态：0->不显示；1->显示
	Content         string    //
	ForwardCount    int32     // 转发数
	CategoryName    string    // 专题分类名称
}
