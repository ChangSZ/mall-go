package ums_member_receive_address

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_member_receive_address"
	"github.com/ChangSZ/mall-go/internal/services/mall_portal/ums_member"
	"github.com/ChangSZ/mall-go/pkg/copy"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Add(ctx context.Context, param dto.UmsMemberReceiveAddress) (int64, error) {
	currentMember, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return 0, err
	}

	data := &ums_member_receive_address.UmsMemberReceiveAddress{}
	copy.AssignStruct(&param, data)
	data.MemberId = currentMember.Id
	return data.Create(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) Delete(ctx context.Context, id int64) (int64, error) {
	currentMember, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return 0, err
	}

	qb := ums_member_receive_address.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id).WhereMemberId(mysql.EqualPredicate, currentMember.Id)
	return qb.Delete(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) Update(ctx context.Context, id int64, param dto.UmsMemberReceiveAddress) (int64, error) {
	currentMember, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return 0, err
	}

	if param.DefaultStatus == 1 {
		// 先将原来的默认地址去除
		qb := ums_member_receive_address.NewQueryBuilder()
		qb = qb.WhereMemberId(mysql.EqualPredicate, currentMember.Id).WhereDefaultStatus(mysql.EqualPredicate, 1)
		_, err := qb.Updates(mysql.DB().GetDbW().WithContext(ctx), map[string]interface{}{"default_status": 0})
		if err != nil {
			return 0, err
		}
	}

	data := map[string]interface{}{
		"member_id":      param.MemberId,
		"name":           param.Name,
		"phone_number":   param.PhoneNumber,
		"default_status": param.DefaultStatus,
		"post_code":      param.PostCode,
		"province":       param.Province,
		"city":           param.City,
		"region":         param.Region,
		"detail_address": param.DetailAddress,
	}
	qb := ums_member_receive_address.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id).WhereMemberId(mysql.EqualPredicate, currentMember.Id)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) List(ctx context.Context) ([]dto.UmsMemberReceiveAddress, error) {
	currentMember, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return nil, err
	}

	qb := ums_member_receive_address.NewQueryBuilder()
	qb = qb.WhereMemberId(mysql.EqualPredicate, currentMember.Id)
	list, err := qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}

	listData := make([]dto.UmsMemberReceiveAddress, 0, len(list))
	for _, v := range list {
		tmp := dto.UmsMemberReceiveAddress{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, nil
}

func (s *service) GetItem(ctx context.Context, id int64) (*dto.UmsMemberReceiveAddress, error) {
	currentMember, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return nil, err
	}

	qb := ums_member_receive_address.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id).WhereMemberId(mysql.EqualPredicate, currentMember.Id)
	item, err := qb.First(mysql.DB().GetDbW().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	res := &dto.UmsMemberReceiveAddress{}
	copy.AssignStruct(item, res)
	return res, nil
}
