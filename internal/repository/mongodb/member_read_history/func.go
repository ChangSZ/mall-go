package member_read_history

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ChangSZ/mall-go/internal/repository/mongodb"
)

var Collection = mongodb.DB().Connection().Collection("memberReadHistory")

func NewModel() *MemberReadHistory {
	return new(MemberReadHistory)
}

func (m *MemberReadHistory) InsertOrUpdate(ctx context.Context) (*mongo.UpdateResult, error) {
	filter := bson.M{"product_id": m.ProductId, "member_id": m.MemberId}
	update := bson.M{"$set": m}
	opts := options.Update().SetUpsert(true)
	return Collection.UpdateOne(ctx, filter, update, opts)
}

// FindByMemberIDOrderByCreateTimeDesc 根据会员ID分页查找记录并按创建时间降序排序
func FindByMemberIDOrderByCreateTimeDesc(ctx context.Context, memberID int64, page, size int64) (
	[]*MemberReadHistory, int64, error) {
	filter := bson.M{"memberId": memberID}

	// 查询总记录数
	total, err := Collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	options := options.Find()
	options.SetSort(bson.D{{Key: "createTime", Value: -1}}) // 按创建时间降序排序
	options.SetSkip((page - 1) * size)
	options.SetLimit(size)

	cursor, err := Collection.Find(ctx, filter, options)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var results []*MemberReadHistory
	for cursor.Next(ctx) {
		var elem MemberReadHistory
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

// BatchDelete 批量删除
func BatchDelete(ctx context.Context, ids []string) (int64, error) {
	objectIDs := make([]primitive.ObjectID, len(ids))
	for i, id := range ids {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return 0, err
		}
		objectIDs[i] = objID
	}

	filter := bson.M{"_id": bson.M{"$in": objectIDs}}

	result, err := Collection.DeleteMany(ctx, filter)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}
