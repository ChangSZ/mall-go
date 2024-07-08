package sms_home_brand

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/sms_home_brand"

	"github.com/ChangSZ/golib/copy"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Create(ctx context.Context, param []dto.SmsHomeBrand) (int64, error) {
	datas := make([]sms_home_brand.SmsHomeBrand, 0, len(param))
	for _, v := range param {
		data := sms_home_brand.SmsHomeBrand{}
		copy.AssignStruct(&v, &data)
		data.Id = 0
		data.RecommendStatus = 1
		data.Sort = 0
		datas = append(datas, data)
	}

	if err := mysql.DB().GetDbW().WithContext(ctx).CreateInBatches(datas, len(datas)).Error; err != nil {
		return 0, err
	}
	return int64(len(datas)), nil
}

func (s *service) UpdateSort(ctx context.Context, id int64, sort int32) (int64, error) {
	data := map[string]interface{}{
		"sort": sort,
	}
	qb := sms_home_brand.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) Delete(ctx context.Context, ids []int64) (int64, error) {
	qb := sms_home_brand.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	return qb.Delete(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) UpdateRecommendStatus(ctx context.Context, ids []int64, recommendStatus int32) (int64, error) {
	data := map[string]interface{}{
		"recommend_status": recommendStatus,
	}
	qb := sms_home_brand.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) List(ctx context.Context, brandName string, recommendStatus int32, pageSize, pageNum int) (
	[]dto.SmsHomeBrand, int64, error) {
	qb := sms_home_brand.NewQueryBuilder()
	if brandName != "" {
		qb = qb.WhereBrandName(mysql.LikePredicate, "%"+brandName+"%")
	}
	if recommendStatus != 0 {
		qb = qb.WhereRecommendStatus(mysql.EqualPredicate, recommendStatus)
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

	listData := make([]dto.SmsHomeBrand, 0, len(list))
	for _, v := range list {
		tmp := dto.SmsHomeBrand{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, count, err
}
