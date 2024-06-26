package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MemberProductCollection 会员商品收藏
type MemberProductCollection struct {
	Id              primitive.ObjectID `json:"_id"`
	MemberId        int64              `json:"memberId"`
	MemberNickname  string             `json:"memberNickname"`
	MemberIcon      string             `json:"memberIcon"`
	ProductId       int64              `json:"productId"`
	ProductName     string             `json:"productName"`
	ProductPic      string             `json:"productPic"`
	ProductSubTitle string             `json:"productSubTitle"`
	ProductPrice    float64            `json:"productPrice"`
	CreateTime      time.Time          `json:"create_time"`
}
