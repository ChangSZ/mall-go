package pms_product_vertify_record

import "time"

// PmsProductVertifyRecord 商品审核记录
//
//go:generate gormgen -structs PmsProductVertifyRecord -input .
type PmsProductVertifyRecord struct {
	Id         int64     //
	ProductId  int64     //
	CreateTime time.Time `gorm:"autoCreateTime"` //
	VertifyMan string    // 审核人
	Status     int32     //
	Detail     string    // 反馈详情
}
