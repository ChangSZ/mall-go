package oms_order_setting

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/oms_order_setting"
	"github.com/ChangSZ/mall-go/pkg/copy"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) GetItem(ctx context.Context, id int64) (*dto.OmsOrderSetting, error) {
	qb := oms_order_setting.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	data, err := qb.First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil
	}
	res := &dto.OmsOrderSetting{}
	copy.AssignStruct(data, res)
	return res, nil
}

func (s *service) Update(ctx context.Context, id int64, param dto.OmsOrderSetting) (int64, error) {
	data := map[string]interface{}{
		"flash_order_overtime":  param.FlashOrderOvertime,
		"normal_order_overtime": param.NormalOrderOvertime,
		"confirm_overtime":      param.ConfirmOvertime,
		"finish_overtime":       param.FinishOvertime,
		"comment_overtime":      param.CommentOvertime,
	}
	qb := oms_order_setting.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}
