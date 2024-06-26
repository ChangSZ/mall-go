package member_collection

import (
	"context"
	"fmt"
	"time"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mongodb/member_product_collection"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product"
	"github.com/ChangSZ/mall-go/internal/services/mall_portal/ums_member"
	"github.com/ChangSZ/mall-go/pkg/copy"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Add(ctx context.Context, param dto.MemberProductCollection) (int64, error) {
	if param.ProductId == 0 {
		return 0, nil
	}
	member, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return 0, err
	}

	data := member_product_collection.NewModel()
	copy.AssignStruct(&param, data)
	data.MemberId = member.Id
	data.MemberNickname = member.Nickname
	data.MemberIcon = member.Icon
	data.CreateTime = time.Now()
	findCollection, err := member_product_collection.FindByMemberIDAndProductID(ctx, data.MemberId, data.ProductId)
	if err != nil {
		return 0, err
	}

	if findCollection == nil {
		product, err := pms_product.NewQueryBuilder().
			WhereId(mysql.EqualPredicate, data.ProductId).
			First(mysql.DB().GetDbR().WithContext(ctx))
		if err != nil {
			return 0, err
		}
		if product == nil {
			return 0, fmt.Errorf("数据库未找到该品牌信息")
		}
		if product.DeleteStatus == 1 {
			return 0, fmt.Errorf("该商品已被删除")
		}

		data.ProductName = product.Name
		data.ProductSubTitle = product.SubTitle
		data.ProductPrice = product.Price
		data.ProductPic = product.Pic
		if _, err := data.Insert(ctx); err != nil {
			return 0, err
		}
		return 1, nil
	}
	return 0, nil
}

func (s *service) Delete(ctx context.Context, productId int64) (int64, error) {
	member, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return 0, err
	}
	return member_product_collection.DeleteByMemberIDAndProductID(ctx, member.Id, productId)
}

func (s *service) List(ctx context.Context, pageNum, pageSize int64) (
	[]dto.MemberProductCollection, int64, error) {
	member, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return nil, 0, err
	}

	list, total, err := member_product_collection.FindByMemberIDWithPagination(ctx, member.Id, pageNum, pageSize)
	if err != nil {
		return nil, 0, err
	}
	listData := make([]dto.MemberProductCollection, 0, len(list))
	for _, v := range list {
		tmp := dto.MemberProductCollection{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, total, nil
}

func (s *service) Detail(ctx context.Context, productId int64) (*dto.MemberProductCollection, error) {
	member, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return nil, err
	}
	data, err := member_product_collection.FindByMemberIDAndProductID(ctx, member.Id, productId)
	if err != nil || data == nil {
		return nil, err
	}

	res := &dto.MemberProductCollection{}
	copy.AssignStruct(data, res)
	return res, nil
}

func (s *service) Clear(ctx context.Context) (int64, error) {
	member, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return 0, err
	}
	return member_product_collection.DeleteAllByMemberID(ctx, member.Id)
}
