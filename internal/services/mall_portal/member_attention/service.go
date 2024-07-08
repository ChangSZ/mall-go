package member_attention

import (
	"context"
	"fmt"
	"time"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mongodb/member_brand_attention"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_brand"
	"github.com/ChangSZ/mall-go/internal/services/mall_portal/ums_member"

	"github.com/ChangSZ/golib/copy"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Add(ctx context.Context, param dto.MemberBrandAttention) (int64, error) {
	if param.BrandId == 0 {
		return 0, nil
	}
	member, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return 0, err
	}

	data := member_brand_attention.NewModel()
	copy.AssignStruct(&param, data)
	data.MemberId = member.Id
	data.MemberNickname = member.Nickname
	data.MemberIcon = member.Icon
	data.CreateTime = time.Now()
	findAttention, err := member_brand_attention.FindByMemberIDAndBrandID(ctx, data.MemberId, data.BrandId)
	if err != nil {
		return 0, err
	}

	if findAttention == nil {
		brand, err := pms_brand.NewQueryBuilder().
			WhereId(mysql.EqualPredicate, data.BrandId).
			First(mysql.DB().GetDbR().WithContext(ctx))
		if err != nil {
			return 0, err
		}
		if brand == nil {
			return 0, fmt.Errorf("数据库未找到该品牌信息")
		}

		data.BrandCity = ""
		data.BrandName = brand.Name
		data.BrandLogo = brand.Logo
		if _, err := data.Insert(ctx); err != nil {
			return 0, err
		}
		return 1, nil
	}
	return 0, nil
}

func (s *service) Delete(ctx context.Context, brandId int64) (int64, error) {
	member, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return 0, err
	}
	return member_brand_attention.DeleteByMemberIDAndBrandID(ctx, member.Id, brandId)
}

func (s *service) List(ctx context.Context, pageNum, pageSize int64) (
	[]dto.MemberBrandAttention, int64, error) {
	member, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return nil, 0, err
	}

	list, total, err := member_brand_attention.FindByMemberIDWithPagination(ctx, member.Id, pageNum, pageSize)
	if err != nil {
		return nil, 0, err
	}
	listData := make([]dto.MemberBrandAttention, 0, len(list))
	for _, v := range list {
		tmp := dto.MemberBrandAttention{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, total, nil
}

func (s *service) Detail(ctx context.Context, brandId int64) (*dto.MemberBrandAttention, error) {
	member, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return nil, err
	}
	data, err := member_brand_attention.FindByMemberIDAndBrandID(ctx, member.Id, brandId)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil
	}
	res := &dto.MemberBrandAttention{}
	copy.AssignStruct(data, res)
	return res, nil
}

func (s *service) Clear(ctx context.Context) (int64, error) {
	member, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return 0, err
	}
	return member_brand_attention.DeleteAllByMemberID(ctx, member.Id)
}
