package sms_flash_promotion_product_relation

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/sms_flash_promotion_product_relation"

	"github.com/ChangSZ/golib/copy"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Create(ctx context.Context, param []dto.SmsFlashPromotionProductRelation) (int64, error) {
	datas := make([]sms_flash_promotion_product_relation.SmsFlashPromotionProductRelation, 0, len(param))
	for _, v := range param {
		data := sms_flash_promotion_product_relation.SmsFlashPromotionProductRelation{}
		copy.AssignStruct(&v, &data)
		data.Id = 0
		datas = append(datas, data)
	}

	if err := mysql.DB().GetDbW().WithContext(ctx).CreateInBatches(datas, len(datas)).Error; err != nil {
		return 0, err
	}
	return int64(len(datas)), nil
}

func (s *service) Update(ctx context.Context, id int64, param dto.SmsFlashPromotionProductRelation) (int64, error) {
	data := map[string]interface{}{
		"id":                         id,
		"flash_promotion_id":         param.FlashPromotionId,
		"flash_promotion_session_id": param.FlashPromotionSessionId,
		"product_id":                 param.ProductId,
		"flash_promotion_price":      param.FlashPromotionPrice,
		"flash_promotion_count":      param.FlashPromotionCount,
		"flash_promotion_limit":      param.FlashPromotionLimit,
		"sort":                       param.Sort,
	}
	qb := sms_flash_promotion_product_relation.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) Delete(ctx context.Context, id int64) (int64, error) {
	qb := sms_flash_promotion_product_relation.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Delete(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) GetItem(ctx context.Context, id int64) (*dto.SmsFlashPromotionProductRelation, error) {
	qb := sms_flash_promotion_product_relation.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	data, err := qb.First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	res := &dto.SmsFlashPromotionProductRelation{}
	copy.AssignStruct(data, res)
	return res, nil
}

func (s *service) List(ctx context.Context, flashPromotionId, flashPromotionSessionId int64, pageSize, pageNum int) (
	[]dto.SmsFlashPromotionProductRelation, int64, error) {
	return new(dao.SmsFlashPromotionProductRelationDao).GetList(ctx,
		mysql.DB().GetDbR().WithContext(ctx), flashPromotionId, flashPromotionSessionId, pageSize, pageNum)
}

func (s *service) GetCount(ctx context.Context, flashPromotionId, flashPromotionSessionId int64) (int64, error) {
	qb := sms_flash_promotion_product_relation.NewQueryBuilder()
	qb = qb.WhereFlashPromotionId(mysql.EqualPredicate, flashPromotionId)
	qb = qb.WhereFlashPromotionSessionId(mysql.EqualPredicate, flashPromotionSessionId)
	return qb.Count(mysql.DB().GetDbR().WithContext(ctx))
}
