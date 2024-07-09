package member_brand_attention

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ChangSZ/mall-go/internal/repository/mongodb"
)

var Collection = mongodb.DB().Connection().Collection("memberBrandAttention")

func NewModel() *MemberBrandAttention {
	return new(MemberBrandAttention)
}

func (m *MemberBrandAttention) Insert(ctx context.Context) (*mongo.InsertOneResult, error) {
	return Collection.InsertOne(ctx, m)
}

// FindByMemberIDAndBrandID 根据会员ID和品牌ID查找记录
func FindByMemberIDAndBrandID(ctx context.Context, memberID, brandID int64) (
	*MemberBrandAttention, error) {
	var result MemberBrandAttention
	filter := bson.M{"memberId": memberID, "brandId": brandID}
	err := Collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &result, err
}

// DeleteByMemberIDAndBrandID 根据会员ID和品牌ID删除记录
func DeleteByMemberIDAndBrandID(ctx context.Context, memberID, brandID int64) (int64, error) {
	filter := bson.M{"memberId": memberID, "brandId": brandID}
	result, err := Collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}

// FindByMemberIDWithPagination 根据会员ID分页查找记录
func FindByMemberIDWithPagination(ctx context.Context, memberID int64, page, size int64) (
	[]*MemberBrandAttention, int64, error) {
	filter := bson.M{"memberId": memberID}

	// 查询总记录数
	total, err := Collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	options := options.Find()
	options.SetSkip((page - 1) * size)
	options.SetLimit(size)

	cursor, err := Collection.Find(ctx, filter, options)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var results []*MemberBrandAttention
	for cursor.Next(ctx) {
		var elem MemberBrandAttention
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, 0, err
		}
		results = append(results, &elem)
	}
	if err := cursor.Err(); err != nil {
		return nil, 0, err
	}

	return results, total, nil
}

// DeleteAllByMemberID 根据会员ID删除记录
func DeleteAllByMemberID(ctx context.Context, memberID int64) (int64, error) {
	filter := bson.M{"memberId": memberID}
	result, err := Collection.DeleteMany(ctx, filter)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}
