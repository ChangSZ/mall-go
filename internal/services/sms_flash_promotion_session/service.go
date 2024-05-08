package sms_flash_promotion_session

import (
	"context"
	"time"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/sms_flash_promotion_session"
	"github.com/ChangSZ/mall-go/internal/services/sms_flash_promotion_product_relation"
	"github.com/ChangSZ/mall-go/pkg/copy"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Create(ctx context.Context, param dto.SmsFlashPromotionSession) (int64, error) {
	data := sms_flash_promotion_session.SmsFlashPromotionSession{}
	copy.AssignStruct(&param, &data)
	data.Id = 0
	data.CreateTime = time.Now()
	return data.Create(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) Update(ctx context.Context, id int64, param dto.SmsFlashPromotionSession) (int64, error) {
	data := map[string]interface{}{
		"name":      param.Name,
		"startTime": param.StartTime,
		"endTime":   param.EndTime,
		"status":    param.Status,
	}
	qb := sms_flash_promotion_session.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) UpdateStatus(ctx context.Context, id int64, status int32) (int64, error) {
	data := map[string]interface{}{
		"status": status,
	}
	qb := sms_flash_promotion_session.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) Delete(ctx context.Context, id int64) (int64, error) {
	qb := sms_flash_promotion_session.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Delete(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) GetItem(ctx context.Context, id int64) (*dto.SmsFlashPromotionSession, error) {
	qb := sms_flash_promotion_session.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	item, err := qb.First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	res := &dto.SmsFlashPromotionSession{}
	copy.AssignStruct(item, res)
	return res, nil
}

func (s *service) ListAll(ctx context.Context) ([]dto.SmsFlashPromotionSession, error) {
	qb := sms_flash_promotion_session.NewQueryBuilder()
	list, err := qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	listData := make([]dto.SmsFlashPromotionSession, 0, len(list))
	for _, v := range list {
		tmp := dto.SmsFlashPromotionSession{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, nil
}

func (s *service) SelectList(ctx context.Context, flashPromotionId int64) (
	[]dto.SmsFlashPromotionSessionDetail, error) {
	qb := sms_flash_promotion_session.NewQueryBuilder()
	qb = qb.WhereStatus(mysql.EqualPredicate, 1)
	list, err := qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}

	relationService := sms_flash_promotion_product_relation.New()
	res := make([]dto.SmsFlashPromotionSessionDetail, 0, len(list))
	for _, v := range list {
		tmp := dto.SmsFlashPromotionSessionDetail{}
		copy.AssignStruct(v, &tmp)
		count, err := relationService.GetCount(ctx, flashPromotionId, v.Id)
		if err != nil {
			return nil, err
		}
		tmp.ProductCount = count
		res = append(res, tmp)
	}
	return res, nil
}
