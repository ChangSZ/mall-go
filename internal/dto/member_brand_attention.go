package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MemberBrandAttention 会员品牌关注
type MemberBrandAttention struct {
	Id             primitive.ObjectID `json:"_id"`
	MemberId       int64              `json:"memberId"`
	MemberNickname string             `json:"memberNickname"`
	MemberIcon     string             `json:"memberIcon"`
	BrandId        int64              `json:"brandId"`
	BrandName      string             `json:"brandName"`
	BrandLogo      string             `json:"brandLogo"`
	BrandCity      string             `json:"brandCity"`
	CreateTime     time.Time          `json:"createTime"`
}
