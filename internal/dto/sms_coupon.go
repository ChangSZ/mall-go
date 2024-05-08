package dto

import "time"

// SmsCouponParam 优惠券信息封装，包括绑定商品和分类
type SmsCouponParam struct {
	SmsCoupon                   `json:",inline"`
	ProductRelationList         []SmsCouponProductRelation         `json:"productRelationList"  gorm:"foreignKey:CouponId"`         // 优惠券绑定的商品
	ProductCategoryRelationList []SmsCouponProductCategoryRelation `json:"productCategoryRelationList"  gorm:"foreignKey:CouponId"` // 优惠券绑定的商品分类

}

type SmsCoupon struct {
	Id           int64     `json:"id"`
	Type         int32     `json:"type"`
	Name         string    `json:"name"`
	Platform     int32     `json:"platform"`
	Count        int32     `json:"count"`
	Amount       float64   `json:"amount"`
	PerLimit     int32     `json:"perLimit"`
	MinPoint     float64   `json:"minPoint"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
	UseType      int32     `json:"useType"`
	Note         string    `json:"note"`
	PublishCount int32     `json:"publishCount"`
	UseCount     int32     `json:"useCount"`
	ReceiveCount int32     `json:"receiveCount"`
	EnableTime   time.Time `json:"enableTime"`
	Code         string    `json:"code"`
	MemberLevel  int32     `json:"memberLevel"`
}
