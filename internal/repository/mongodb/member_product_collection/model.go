package member_product_collection

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MemberProductCollection 会员商品收藏
type MemberProductCollection struct {
	Id              primitive.ObjectID `bson:"_id,omitempty"`
	MemberId        int64              `bson:"memberId"`
	MemberNickname  string             `bson:"memberNickname"`
	MemberIcon      string             `bson:"memberIcon"`
	ProductId       int64              `bson:"productId"`
	ProductName     string             `bson:"productName"`
	ProductPic      string             `bson:"productPic"`
	ProductSubTitle string             `bson:"productSubTitle"`
	ProductPrice    float64            `bson:"productPrice"`
	CreateTime      time.Time          `bson:"create_time"`
}
