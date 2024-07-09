package home

import (
	"context"
	"time"

	"github.com/ChangSZ/golib/copy"

	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/cms_subject"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product_category"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/sms_flash_promotion"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/sms_flash_promotion_session"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/sms_home_advertise"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Content(ctx context.Context) (*dto.HomeContentResult, error) {
	result := &dto.HomeContentResult{}
	{
		// 获取首页广告
		data, err := s.GetHomeAdvertiseList(ctx)
		if err != nil {
			return nil, err
		}
		result.AdvertiseList = data
	}

	{
		// 获取推荐品牌
		data, err := new(dao.HomeDao).GetRecommendBrandList(ctx, mysql.DB().GetDbR().WithContext(ctx), 1, 6)
		if err != nil {
			return nil, err
		}
		result.BrandList = data
	}

	{
		// 获取秒杀信息
		data, err := s.GetHomeFlashPromotion(ctx)
		if err != nil {
			return nil, err
		}
		result.HomeFlashPromotion = *data
	}

	{
		// 获取新品推荐
		data, err := new(dao.HomeDao).GetNewProductList(ctx, mysql.DB().GetDbR().WithContext(ctx), 1, 4)
		if err != nil {
			return nil, err
		}
		result.NewProductList = data
	}

	{
		// 获取人气推荐
		data, err := new(dao.HomeDao).GetHotProductList(ctx, mysql.DB().GetDbR().WithContext(ctx), 1, 4)
		if err != nil {
			return nil, err
		}
		result.HotProductList = data
	}

	{
		// 获取推荐专题
		data, err := new(dao.HomeDao).GetRecommendSubjectList(ctx, mysql.DB().GetDbR().WithContext(ctx), 1, 4)
		if err != nil {
			return nil, err
		}
		result.SubjectList = data
	}

	return result, nil
}

func (s *service) RecommendProductList(ctx context.Context, pageNum, pageSize int) ([]dto.PmsProduct, error) {
	// 暂时默认推荐所有商品
	offset := (pageNum - 1) * pageSize
	list, err := pms_product.NewQueryBuilder().
		WhereDeleteStatus(mysql.EqualPredicate, 0).
		WherePublishStatus(mysql.EqualPredicate, 1).
		Offset(offset).
		Limit(pageSize).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}

	listData := make([]dto.PmsProduct, 0, len(list))
	for _, v := range list {
		tmp := dto.PmsProduct{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, nil
}

func (s *service) GetProductCateList(ctx context.Context, parentId int64) ([]dto.PmsProductCategory, error) {
	list, err := pms_product_category.NewQueryBuilder().
		WhereShowStatus(mysql.EqualPredicate, 1).
		WhereParentId(mysql.EqualPredicate, parentId).
		OrderBySort(false).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}

	listData := make([]dto.PmsProductCategory, 0, len(list))
	for _, v := range list {
		tmp := dto.PmsProductCategory{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, nil
}

func (s *service) GetSubjectList(ctx context.Context, cateId int64, pageNum, pageSize int) ([]dto.CmsSubject, error) {
	offset := (pageNum - 1) * pageSize
	qb := cms_subject.NewQueryBuilder().
		WhereShowStatus(mysql.EqualPredicate, 0)
	if cateId != 0 {
		qb = qb.WhereCategoryId(mysql.EqualPredicate, cateId)
	}

	list, err := qb.Offset(offset).Limit(pageSize).QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}

	listData := make([]dto.CmsSubject, 0, len(list))
	for _, v := range list {
		tmp := dto.CmsSubject{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, nil
}

func (s *service) HotProductList(ctx context.Context, pageNum, pageSize int) ([]dto.PmsProduct, error) {
	return new(dao.HomeDao).GetHotProductList(ctx, mysql.DB().GetDbR().WithContext(ctx), pageNum, pageSize)
}

func (s *service) NewProductList(ctx context.Context, pageNum, pageSize int) ([]dto.PmsProduct, error) {
	return new(dao.HomeDao).GetNewProductList(ctx, mysql.DB().GetDbR().WithContext(ctx), pageNum, pageSize)
}

func (s *service) GetHomeFlashPromotion(ctx context.Context) (*dto.HomeFlashPromotion, error) {
	homeFlashPromotion := &dto.HomeFlashPromotion{}
	now := time.Now()
	// 获取当前秒杀活动
	flashPromotion, err := s.GetFlashPromotion(ctx, now)
	if err != nil || flashPromotion == nil {
		return homeFlashPromotion, err
	}

	// 获取当前秒杀场次
	flashPromotionSession, err := s.GetFlashPromotionSession(ctx, now)
	if err != nil || flashPromotionSession == nil {
		return homeFlashPromotion, err
	}

	homeFlashPromotion.StartTime = flashPromotionSession.StartTime
	homeFlashPromotion.EndTime = flashPromotionSession.EndTime

	// 获取下一个秒杀场次
	nextSession, err := s.GetNextFlashPromotionSession(ctx, homeFlashPromotion.StartTime)
	if err != nil || nextSession == nil {
		return homeFlashPromotion, err
	}
	homeFlashPromotion.NextStartTime = nextSession.StartTime
	homeFlashPromotion.NextEndTime = nextSession.EndTime

	// 获取秒杀商品
	flashProductList, err := new(dao.HomeDao).GetFlashProductList(ctx,
		mysql.DB().GetDbR().WithContext(ctx), flashPromotion.Id, flashPromotionSession.Id)
	if err != nil {
		return homeFlashPromotion, err
	}
	homeFlashPromotion.ProductList = flashProductList
	return homeFlashPromotion, nil
}

func (s *service) GetNextFlashPromotionSession(ctx context.Context, date time.Time) (
	*dto.SmsFlashPromotionSession, error) {
	promotionSessionList, err := sms_flash_promotion_session.NewQueryBuilder().
		WhereStatus(mysql.EqualPredicate, 1).
		WhereStartTime(mysql.GreaterThanPredicate, date).
		OrderByStartTime(true).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if len(promotionSessionList) > 0 {
		res := &dto.SmsFlashPromotionSession{}
		copy.AssignStruct(promotionSessionList[0], res)
		return res, nil
	}
	return nil, nil
}

func (s *service) GetHomeAdvertiseList(ctx context.Context) ([]dto.SmsHomeAdvertise, error) {
	list, err := sms_home_advertise.NewQueryBuilder().
		WhereType(mysql.EqualPredicate, 1).
		WhereStatus(mysql.EqualPredicate, 1).
		OrderBySort(false).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}

	listData := make([]dto.SmsHomeAdvertise, 0, len(list))
	for _, v := range list {
		tmp := dto.SmsHomeAdvertise{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, nil
}

// 根据时间获取秒杀活动
func (s *service) GetFlashPromotion(ctx context.Context, date time.Time) (*dto.SmsFlashPromotion, error) {
	flashPromotionList, err := sms_flash_promotion.NewQueryBuilder().
		WhereStatus(mysql.EqualPredicate, 1).
		WhereStartDate(mysql.SmallerThanPredicate, date).
		WhereEndDate(mysql.GreaterThanPredicate, date).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if len(flashPromotionList) > 0 {
		res := &dto.SmsFlashPromotion{}
		copy.AssignStruct(flashPromotionList[0], res)
		return res, nil
	}
	return nil, nil
}

// 根据时间获取秒杀场次
func (s *service) GetFlashPromotionSession(ctx context.Context, date time.Time) (
	*dto.SmsFlashPromotionSession, error) {
	promotionSessionList, err := sms_flash_promotion_session.NewQueryBuilder().
		WhereStatus(mysql.EqualPredicate, 1).
		WhereStartTime(mysql.SmallerThanOrEqualPredicate, date).
		WhereEndTime(mysql.GreaterThanOrEqualPredicate, date).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if len(promotionSessionList) > 0 {
		res := &dto.SmsFlashPromotionSession{}
		copy.AssignStruct(promotionSessionList[0], res)
		return res, nil
	}
	return nil, nil
}
