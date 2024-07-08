package oms_order_return_reason

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/oms_order_return_reason"

	"github.com/ChangSZ/golib/copy"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Create(ctx context.Context, param dto.OmsOrderReturnReason) (int64, error) {
	data := oms_order_return_reason.NewModel()
	copy.AssignStruct(&param, data)
	return data.Create(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) Update(ctx context.Context, id int64, param dto.OmsOrderReturnReason) (int64, error) {
	data := map[string]interface{}{
		"name":        param.Name,
		"sort":        param.Sort,
		"status":      param.Status,
		"create_time": param.CreateTime,
	}
	qb := oms_order_return_reason.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) Delete(ctx context.Context, ids []int64) (int64, error) {
	qb := oms_order_return_reason.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	return qb.Delete(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) List(ctx context.Context, pageSize, pageNum int) ([]dto.OmsOrderReturnReason, int64, error) {
	qb := oms_order_return_reason.NewQueryBuilder()
	count, err := qb.Count(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, 0, err
	}

	offset := (pageNum - 1) * pageSize
	list, err := qb.
		Limit(pageSize).
		Offset(offset).
		OrderBySort(false).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, 0, err
	}

	listData := make([]dto.OmsOrderReturnReason, 0, len(list))
	for _, v := range list {
		tmp := dto.OmsOrderReturnReason{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, count, err
}

func (s *service) UpdateStatus(ctx context.Context, ids []int64, status int32) (int64, error) {
	if status != 0 && status != 1 {
		return 0, nil
	}
	data := map[string]interface{}{
		"status": status,
	}
	qb := oms_order_return_reason.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) GetItem(ctx context.Context, id int64) (*dto.OmsOrderReturnReason, error) {
	qb := oms_order_return_reason.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	data, err := qb.First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	res := &dto.OmsOrderReturnReason{}
	copy.AssignStruct(data, res)
	return res, nil
}
