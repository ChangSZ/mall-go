package sms_home_advertise

import (
	"context"
	"time"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/sms_home_advertise"
	"github.com/ChangSZ/mall-go/pkg/copy"

	"github.com/ChangSZ/golib/log"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Create(ctx context.Context, param dto.SmsHomeAdvertise) (int64, error) {
	data := sms_home_advertise.NewModel()
	copy.AssignStruct(&param, data)
	data.Id = 0
	data.ClickCount = 0
	data.OrderCount = 0
	return data.Create(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) Delete(ctx context.Context, ids []int64) (int64, error) {
	qb := sms_home_advertise.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	return qb.Delete(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) UpdateStatus(ctx context.Context, id int64, status int32) (int64, error) {
	data := map[string]interface{}{
		"status": status,
	}
	qb := sms_home_advertise.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) GetItem(ctx context.Context, id int64) (*dto.SmsHomeAdvertise, error) {
	qb := sms_home_advertise.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	data, err := qb.First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	res := &dto.SmsHomeAdvertise{}
	copy.AssignStruct(data, res)
	return res, nil
}

func (s *service) Update(ctx context.Context, id int64, param dto.SmsHomeAdvertise) (int64, error) {
	data := map[string]interface{}{
		"name":        param.Name,
		"type":        param.Type,
		"pic":         param.Pic,
		"start_time":  param.StartTime,
		"end_time":    param.EndTime,
		"status":      param.Status,
		"click_count": param.ClickCount,
		"order_count": param.OrderCount,
		"url":         param.Url,
		"note":        param.Note,
		"sort":        param.Sort,
	}
	qb := sms_home_advertise.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) List(ctx context.Context, name string, adType int32, endTime string, pageSize, pageNum int) (
	[]dto.SmsHomeAdvertise, int64, error) {
	qb := sms_home_advertise.NewQueryBuilder()
	if name != "" {
		qb = qb.WhereName(mysql.LikePredicate, "%"+name+"%")
	}
	if adType != 0 {
		qb = qb.WhereType(mysql.EqualPredicate, adType)
	}
	if endTime != "" {
		startStr := endTime + " 00:00:00"
		endStr := endTime + " 23:59:59"
		start, err := time.Parse("2006-01-02 15:04:05", startStr)
		if err != nil {
			log.WithTrace(ctx).Errorf("解析时间出错: %s, err: %v", startStr, err)
			return nil, 0, err
		}
		end, err := time.Parse("2006-01-02 15:04:05", endStr)
		if err != nil {
			log.WithTrace(ctx).Errorf("解析时间出错: %s, err: %v", endStr, err)
			return nil, 0, err
		}
		qb = qb.WhereEndTime(mysql.GreaterThanOrEqualPredicate, start)
		qb = qb.WhereEndTime(mysql.SmallerThanOrEqualPredicate, end)
	}
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

	listData := make([]dto.SmsHomeAdvertise, 0, len(list))
	for _, v := range list {
		tmp := dto.SmsHomeAdvertise{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, count, err
}
