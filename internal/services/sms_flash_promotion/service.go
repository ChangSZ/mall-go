package sms_flash_promotion

import (
	"context"
	"time"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/sms_flash_promotion"
	"github.com/ChangSZ/mall-go/pkg/copy"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Create(ctx context.Context, param dto.SmsFlashPromotion) (int64, error) {
	data := sms_flash_promotion.NewModel()
	copy.AssignStruct(&param, data)
	data.Id = 0
	data.CreateTime = time.Now()
	return data.Create(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) Update(ctx context.Context, id int64, param dto.SmsFlashPromotion) (int64, error) {
	data := map[string]interface{}{
		"title":      param.Title,
		"start_date": param.StartDate,
		"end_date":   param.EndDate,
		"status":     param.Status,
	}
	qb := sms_flash_promotion.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) Delete(ctx context.Context, id int64) (int64, error) {
	qb := sms_flash_promotion.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Delete(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) UpdateStatus(ctx context.Context, id int64, status int32) (int64, error) {
	data := map[string]interface{}{
		"status": status,
	}
	qb := sms_flash_promotion.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) GetItem(ctx context.Context, id int64) (*dto.SmsFlashPromotion, error) {
	qb := sms_flash_promotion.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	item, err := qb.First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	res := &dto.SmsFlashPromotion{}
	copy.AssignStruct(item, res)
	return res, nil
}

func (s *service) List(ctx context.Context, keyword string, pageSize, pageNum int) (
	[]dto.SmsFlashPromotion, int64, error) {
	qb := sms_flash_promotion.NewQueryBuilder()
	if keyword != "" {
		qb = qb.WhereTitle(mysql.LikePredicate, "%"+keyword+"%")
	}
	count, err := qb.Count(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, 0, err
	}

	offset := (pageNum - 1) * pageSize
	list, err := qb.
		Limit(pageSize).
		Offset(offset).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, 0, err
	}

	listData := make([]dto.SmsFlashPromotion, 0, len(list))
	for _, v := range list {
		tmp := dto.SmsFlashPromotion{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, count, err
}
