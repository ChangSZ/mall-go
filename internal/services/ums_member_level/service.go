package ums_member_level

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_member_level"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) List(ctx context.Context, defaultStatus int32) ([]dto.UmsMemberLevel, error) {

	qb := ums_member_level.NewQueryBuilder()
	qb = qb.WhereDefaultStatus(mysql.EqualPredicate, int32(defaultStatus))
	list, err := qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	listData := make([]dto.UmsMemberLevel, 0, len(list))
	for _, v := range list {
		listData = append(listData, dto.UmsMemberLevel{
			Id:                    v.Id,
			Name:                  v.Name,
			GrowthPoint:           v.GrowthPoint,
			DefaultStatus:         v.DefaultStatus,
			FreeFreightPoint:      v.FreeFreightPoint,
			CommentGrowthPoint:    v.CommentGrowthPoint,
			PriviledgeFreeFreight: v.PriviledgeFreeFreight,
			PriviledgeSignIn:      v.PriviledgeSignIn,
			PriviledgeComment:     v.PriviledgeComment,
			PriviledgePromotion:   v.PriviledgePromotion,
			PriviledgeMemberPrice: v.PriviledgeMemberPrice,
			PriviledgeBirthday:    v.PriviledgeBirthday,
			Note:                  v.Note,
		})
	}
	return listData, nil
}
