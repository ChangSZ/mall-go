package pms_comment

import "time"

// PmsComment 商品评价表
//
//go:generate gormgen -structs PmsComment -input .
type PmsComment struct {
	Id               int64     //
	ProductId        int64     //
	MemberNickName   string    //
	ProductName      string    //
	Star             int32     // 评价星数：0->5
	MemberIp         string    // 评价的ip
	CreateTime       time.Time `gorm:"autoCreateTime"` //
	ShowStatus       int32     //
	ProductAttribute string    // 购买时的商品属性
	CollectCouont    int32     //
	ReadCount        int32     //
	Content          string    //
	Pics             string    // 上传图片地址，以逗号隔开
	MemberIcon       string    // 评论用户头像
	ReplayCount      int32     //
}
