package oms_order_return_apply

import (
	"context"
	"time"

	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/oms_order_return_apply"
	"github.com/ChangSZ/mall-go/pkg/pagehelper"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) List(ctx context.Context, queryParam dto.OmsReturnApplyQueryParam, pageSize, pageNum int) (
	*pagehelper.ListData[dto.OmsOrderReturnApply], error) {
	list, total, err := new(dao.OmsOrderReturnApplyDao).List(
		ctx, mysql.DB().GetDbR().WithContext(ctx), queryParam, pageSize, pageNum)
	res := pagehelper.New[dto.OmsOrderReturnApply]()
	res.Set(pageNum, pageSize, total, list)
	return res, err
}

func (s *service) Delete(ctx context.Context, ids []int64) (int64, error) {
	qb := oms_order_return_apply.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	return qb.Delete(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) UpdateStatus(ctx context.Context, id int64, param dto.OmsUpdateStatusParam) (int64, error) {
	data := map[string]interface{}{}
	switch param.Status {
	case 1:
		// 确认退货
		data["status"] = 1
		data["return_amount"] = param.ReturnAmount
		data["company_address_id"] = param.CompanyAddressID
		data["handle_time"] = time.Now()
		data["handle_man"] = param.HandleMan
		data["handle_note"] = param.HandleNote
	case 2:
		// 完成退货
		data["status"] = 2
		data["receive_time"] = time.Now()
		data["receive_man"] = param.ReceiveMan
		data["receive_note"] = param.ReceiveNote
	case 3:
		// 拒绝退货
		data["status"] = 3
		data["handle_time"] = time.Now()
		data["handle_man"] = param.HandleMan
		data["handle_note"] = param.HandleNote
	default:
		return 0, nil
	}
	qb := oms_order_return_apply.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) GetItem(ctx context.Context, id int64) (*dto.OmsOrderReturnApplyResult, error) {
	return new(dao.OmsOrderReturnApplyDao).GetDetail(ctx, mysql.DB().GetDbR().WithContext(ctx), id)
}
