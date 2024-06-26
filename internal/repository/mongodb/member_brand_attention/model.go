package member_brand_attention

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MemberBrandAttention 会员品牌关注
type MemberBrandAttention struct {
	Id             primitive.ObjectID `bson:"_id,omitempty"`
	MemberId       int64              `bson:"memberId"`
	MemberNickname string             `bson:"memberNickname"`
	MemberIcon     string             `bson:"memberIcon"`
	BrandId        int64              `bson:"brandId"`
	BrandName      string             `bson:"brandName"`
	BrandLogo      string             `bson:"brandLogo"`
	BrandCity      string             `bson:"brandCity"`
	CreateTime     time.Time          `bson:"createTime"`
}
