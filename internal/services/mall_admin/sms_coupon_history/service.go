package sms_coupon_history

import (
	"context"

	"github.com/ChangSZ/golib/copy"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/sms_coupon_history"
	"github.com/ChangSZ/mall-go/pkg/pagehelper"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) List(ctx context.Context, couponId int64, useStatus int32, orderSn string, pageSize, pageNum int) (
	*pagehelper.ListData[dto.SmsCouponHistory], error) {
	res := pagehelper.New[dto.SmsCouponHistory]()
	qb := sms_coupon_history.NewQueryBuilder()
	if couponId != 0 {
		qb = qb.WhereCouponId(mysql.EqualPredicate, couponId)
	}
	if useStatus != 0 {
		qb = qb.WhereUseStatus(mysql.EqualPredicate, useStatus)
	}
	if orderSn != "" {
		qb = qb.WhereOrderSn(mysql.LikePredicate, "%"+orderSn+"%")
	}
	count, err := qb.Count(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return res, err
	}

	offset := (pageNum - 1) * pageSize
	list, err := qb.
		Limit(pageSize).
		Offset(offset).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return res, err
	}

	listData := make([]dto.SmsCouponHistory, 0, len(list))
	for _, v := range list {
		tmp := dto.SmsCouponHistory{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	res.Set(pageNum, pageSize, count, listData)
	return res, err
}
