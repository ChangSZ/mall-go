package cms_member_report

import "time"

// CmsMemberReport 用户举报表
//
//go:generate gormgen -structs CmsMemberReport -input .
type CmsMemberReport struct {
	Id               int64     //
	ReportType       int32     // 举报类型：0->商品评价；1->话题内容；2->用户评论
	ReportMemberName string    // 举报人
	CreateTime       time.Time `gorm:"autoCreateTime"` //
	ReportObject     string    //
	ReportStatus     int32     // 举报状态：0->未处理；1->已处理
	HandleStatus     int32     // 处理结果：0->无效；1->有效；2->恶意
	Note             string    //
}
