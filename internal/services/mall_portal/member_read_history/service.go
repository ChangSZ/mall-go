package member_read_history

import (
	"context"
	"fmt"
	"time"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mongodb/member_read_history"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product"
	"github.com/ChangSZ/mall-go/internal/services/mall_portal/ums_member"

	"github.com/ChangSZ/golib/copy"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Add(ctx context.Context, param dto.MemberReadHistory) (int64, error) {
	if param.ProductId == 0 {
		return 0, nil
	}
	member, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return 0, err
	}

	data := member_read_history.NewModel()
	copy.AssignStruct(&param, data)
	data.MemberId = member.Id
	data.MemberNickname = member.Nickname
	data.MemberIcon = member.Icon
	data.CreateTime = time.Now()

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
	if _, err := data.InsertOrUpdate(ctx); err != nil {
		return 0, err
	}
	return 1, nil
}

func (s *service) Delete(ctx context.Context, ids []string) (int64, error) {
	return member_read_history.BatchDelete(ctx, ids)
}

func (s *service) List(ctx context.Context, pageNum, pageSize int64) (
	[]dto.MemberReadHistory, int64, error) {
	member, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return nil, 0, err
	}

	list, total, err := member_read_history.FindByMemberIDOrderByCreateTimeDesc(ctx, member.Id, pageNum, pageSize)
	if err != nil {
		return nil, 0, err
	}
	listData := make([]dto.MemberReadHistory, 0, len(list))
	for _, v := range list {
		tmp := dto.MemberReadHistory{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, total, nil
}

func (s *service) Clear(ctx context.Context) (int64, error) {
	member, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return 0, err
	}
	return member_read_history.DeleteAllByMemberID(ctx, member.Id)
}
