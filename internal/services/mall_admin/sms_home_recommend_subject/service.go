package sms_home_recommend_subject

import (
	"context"

	"github.com/ChangSZ/golib/copy"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/sms_home_recommend_subject"
	"github.com/ChangSZ/mall-go/pkg/pagehelper"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Create(ctx context.Context, param []dto.SmsHomeRecommendSubject) (int64, error) {
	datas := make([]sms_home_recommend_subject.SmsHomeRecommendSubject, 0, len(param))
	for _, v := range param {
		data := sms_home_recommend_subject.SmsHomeRecommendSubject{}
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
	qb := sms_home_recommend_subject.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) Delete(ctx context.Context, ids []int64) (int64, error) {
	qb := sms_home_recommend_subject.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	return qb.Delete(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) UpdateRecommendStatus(ctx context.Context, ids []int64, recommendStatus int32) (int64, error) {
	data := map[string]interface{}{
		"recommend_status": recommendStatus,
	}
	qb := sms_home_recommend_subject.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) List(ctx context.Context, subjectName string, recommendStatus int32, pageSize, pageNum int) (
	*pagehelper.ListData[dto.SmsHomeRecommendSubject], error) {
	res := pagehelper.New[dto.SmsHomeRecommendSubject]()
	qb := sms_home_recommend_subject.NewQueryBuilder()
	if subjectName != "" {
		qb = qb.WhereSubjectName(mysql.LikePredicate, "%"+subjectName+"%")
	}
	if recommendStatus != 0 {
		qb = qb.WhereRecommendStatus(mysql.EqualPredicate, recommendStatus)
	}
	count, err := qb.Count(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return res, err
	}

	offset := (pageNum - 1) * pageSize
	list, err := qb.
		Limit(pageSize).
		Offset(offset).
		OrderBySort(false).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return res, err
	}

	listData := make([]dto.SmsHomeRecommendSubject, 0, len(list))
	for _, v := range list {
		tmp := dto.SmsHomeRecommendSubject{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	res.Set(pageNum, pageSize, count, listData)
	return res, err
}
