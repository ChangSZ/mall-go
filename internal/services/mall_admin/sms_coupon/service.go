package sms_coupon

import (
	"context"

	"github.com/ChangSZ/golib/copy"

	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/sms_coupon"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/sms_coupon_product_category_relation"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/sms_coupon_product_relation"
	"github.com/ChangSZ/mall-go/pkg/pagehelper"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Create(ctx context.Context, param dto.SmsCouponParam) (int64, error) {
	data := sms_coupon.NewModel()
	copy.AssignStruct(&param, data)
	data.Count = param.PublishCount
	data.UseCount = 0
	data.ReceiveCount = 0
	// 插入优惠券表
	id, err := data.Create(mysql.DB().GetDbW().WithContext(ctx))
	if err != nil {
		return 0, err
	}

	switch param.UseType {
	case 2:
		// 插入优惠券和商品关系表
		for i := range param.ProductRelationList {
			param.ProductRelationList[i].CouponId = id
		}
		if err := new(dao.SmsCouponProductRelationDao).InsertList(
			ctx, mysql.DB().GetDbW().WithContext(ctx), param.ProductRelationList); err != nil {
			return id, err
		}
	case 1:
		// 插入优惠券和商品分类关系表
		for i := range param.ProductCategoryRelationList {
			param.ProductCategoryRelationList[i].CouponId = id
		}
		if err := new(dao.SmsCouponProductCategoryRelationDao).InsertList(
			ctx, mysql.DB().GetDbW().WithContext(ctx), param.ProductCategoryRelationList); err != nil {
			return id, err
		}
	}

	return id, nil
}

func (s *service) Delete(ctx context.Context, id int64) (int64, error) {
	// 删除优惠券
	qb := sms_coupon.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	cnt, err := qb.Delete(mysql.DB().GetDbW().WithContext(ctx))
	if err != nil {
		return 0, err
	}

	if err := s.deleteProductRelation(ctx, id); err != nil {
		return 0, err
	}
	if err := s.deleteProductCategoryRelation(ctx, id); err != nil {
		return 0, err
	}
	return cnt, nil
}

// deleteProductRelation 删除商品关联
func (s *service) deleteProductRelation(ctx context.Context, couponId int64) error {
	qb := sms_coupon_product_relation.NewQueryBuilder()
	qb = qb.WhereCouponId(mysql.EqualPredicate, couponId)
	if _, err := qb.Delete(mysql.DB().GetDbW().WithContext(ctx)); err != nil {
		return err
	}
	return nil
}

// deleteProductCategoryRelation 删除商品分类关联
func (s *service) deleteProductCategoryRelation(ctx context.Context, couponId int64) error {
	qb := sms_coupon_product_category_relation.NewQueryBuilder()
	qb = qb.WhereCouponId(mysql.EqualPredicate, couponId)
	if _, err := qb.Delete(mysql.DB().GetDbW().WithContext(ctx)); err != nil {
		return err
	}
	return nil
}

func (s *service) Update(ctx context.Context, id int64, param dto.SmsCouponParam) (int64, error) {
	data := map[string]interface{}{
		"type":          param.Type,
		"name":          param.Name,
		"platform":      param.Platform,
		"count":         param.Count,
		"amount":        param.Amount,
		"per_limit":     param.PerLimit,
		"min_point":     param.MinPoint,
		"start_time":    param.StartTime,
		"end_time":      param.EndTime,
		"use_type":      param.UseType,
		"note":          param.Note,
		"publish_count": param.PublishCount,
		"use_count":     param.UseCount,
		"receive_count": param.ReceiveCount,
		"enable_time":   param.EnableTime,
		"code":          param.Code,
		"member_level":  param.MemberLevel,
	}
	qb := sms_coupon.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	cnt, err := qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
	if err != nil {
		return 0, err
	}

	switch param.UseType {
	case 2:
		// 删除后插入优惠券和商品关系表
		for i := range param.ProductRelationList {
			param.ProductRelationList[i].CouponId = id
		}
		if err := s.deleteProductRelation(ctx, id); err != nil {
			return cnt, err
		}
		if err := new(dao.SmsCouponProductRelationDao).InsertList(
			ctx, mysql.DB().GetDbW().WithContext(ctx), param.ProductRelationList); err != nil {
			return cnt, err
		}
	case 1:
		// 删除后插入优惠券和商品分类关系表
		for i := range param.ProductCategoryRelationList {
			param.ProductCategoryRelationList[i].CouponId = id
		}
		if err := s.deleteProductCategoryRelation(ctx, id); err != nil {
			return cnt, err
		}
		if err := new(dao.SmsCouponProductCategoryRelationDao).InsertList(
			ctx, mysql.DB().GetDbW().WithContext(ctx), param.ProductCategoryRelationList); err != nil {
			return cnt, err
		}
	}
	return cnt, nil
}

func (s *service) List(ctx context.Context, name string, couponType int32, pageSize, pageNum int) (
	*pagehelper.ListData[dto.SmsCoupon], error) {
	res := pagehelper.New[dto.SmsCoupon]()
	qb := sms_coupon.NewQueryBuilder()
	if name != "" {
		qb = qb.WhereName(mysql.LikePredicate, "%"+name+"%")
	}
	if couponType != 0 {
		qb = qb.WhereType(mysql.EqualPredicate, couponType)
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

	listData := make([]dto.SmsCoupon, 0, len(list))
	for _, v := range list {
		tmp := dto.SmsCoupon{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	res.Set(pageNum, pageSize, count, listData)
	return res, err
}

func (s *service) GetItem(ctx context.Context, id int64) (*dto.SmsCouponParam, error) {
	return new(dao.SmsCouponDao).GetItem(ctx, mysql.DB().GetDbR().WithContext(ctx), id)
}
