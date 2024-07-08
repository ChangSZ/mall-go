package ums_member_coupon

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/sms_coupon"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/sms_coupon_history"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/sms_coupon_product_category_relation"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/sms_coupon_product_relation"
	"github.com/ChangSZ/mall-go/internal/services/mall_portal/ums_member"

	"github.com/ChangSZ/golib/copy"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Add(ctx context.Context, couponId int64) error {
	currentMember, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return err
	}
	// 获取优惠券信息，判断数量
	qb := sms_coupon.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, couponId)
	coupon, err := qb.First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return err
	}
	if coupon.Id == 0 {
		return fmt.Errorf("优惠券不存在")
	}
	if coupon.Count <= 0 {
		return fmt.Errorf("优惠券已经领完了")
	}
	if time.Now().Before(coupon.EnableTime) {
		return fmt.Errorf("优惠券还没到领取时间")
	}

	// 判断用户领取的优惠券数量是否超过限制
	{
		qb := sms_coupon_history.NewQueryBuilder().
			WhereCouponId(mysql.EqualPredicate, couponId).
			WhereMemberId(mysql.EqualPredicate, currentMember.Id)
		count, err := qb.Count(mysql.DB().GetDbR().WithContext(ctx))
		if err != nil {
			return err
		}
		if count >= int64(coupon.PerLimit) {
			return fmt.Errorf("您已经领取过该优惠券")
		}
	}

	// 生成领取优惠券历史
	{
		couponHistory := sms_coupon_history.NewModel()
		couponHistory.CouponId = couponId
		couponHistory.CouponCode = s.generateCouponCode(currentMember.Id)
		couponHistory.MemberId = currentMember.Id
		couponHistory.MemberNickname = currentMember.Nickname
		couponHistory.GetType = 1   // 主动领取
		couponHistory.UseStatus = 0 // 未使用
		if _, err = couponHistory.Create(mysql.DB().GetDbW().WithContext(ctx)); err != nil {
			return err
		}
	}

	// 更新优惠券记录
	{
		couponUpdateData := map[string]interface{}{
			"count":         coupon.Count - 1,
			"receive_count": coupon.ReceiveCount + 1,
		}
		qb := sms_coupon.NewQueryBuilder().WhereId(mysql.EqualPredicate, couponId)
		if _, err := qb.Updates(mysql.DB().GetDbW().WithContext(ctx), couponUpdateData); err != nil {
			return err
		}
	}
	return nil
}

// generateCouponCode 16位优惠码生成：时间戳后8位+4位随机数+用户id后4位
func (s *service) generateCouponCode(memberId int64) string {
	var sb string
	currentTimeMillis := time.Now().UnixNano() / int64(time.Millisecond)
	timeMillisStr := fmt.Sprintf("%d", currentTimeMillis)
	sb += timeMillisStr[len(timeMillisStr)-8:]

	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 4; i++ {
		sb += fmt.Sprintf("%d", rand.Intn(10))
	}

	memberIdStr := fmt.Sprintf("%d", memberId)
	if len(memberIdStr) <= 4 {
		sb += fmt.Sprintf("%04d", memberId)
	} else {
		sb += memberIdStr[len(memberIdStr)-4:]
	}
	return sb
}

func (s *service) ListHistory(ctx context.Context, useStatus int32) ([]dto.SmsCouponHistory, error) {
	currentMember, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return nil, err
	}
	qb := sms_coupon_history.NewQueryBuilder().WhereMemberId(mysql.EqualPredicate, currentMember.Id)
	if useStatus != 0 {
		qb = qb.WhereUseStatus(mysql.EqualPredicate, useStatus)
	}
	list, err := qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}

	listData := make([]dto.SmsCouponHistory, 0, len(list))
	for _, v := range list {
		tmp := dto.SmsCouponHistory{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, nil
}

func (s *service) ListCart(ctx context.Context,
	cartItemList []dto.CartPromotionItem, enable int32) ([]dto.SmsCouponHistoryDetail, error) {
	currentMember, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return nil, err
	}
	// 获取该用户所有优惠券
	allList, err := new(dao.CouponDao).GetHistoryDetailList(ctx, mysql.DB().GetDbR().WithContext(ctx), currentMember.Id)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	// 根据优惠券使用类型来判断优惠券是否可用
	enableList := make([]dto.SmsCouponHistoryDetail, 0)
	disableList := make([]dto.SmsCouponHistoryDetail, 0)
	for _, v := range allList {
		useType := v.Coupon.UseType
		minPoint := v.Coupon.MinPoint
		endTime := v.Coupon.EndTime
		switch useType {
		case 0: // 全场通用
			// 判断是否满足优惠起点
			// 计算购物车商品的总价
			totalAmount := s.calcTotalAmount(cartItemList)
			if now.Before(endTime) && totalAmount-minPoint >= 0 {
				enableList = append(enableList, v)
			} else {
				disableList = append(disableList, v)
			}
		case 1: // 指定分类
			// 计算指定分类商品的总价
			productCategoryIds := make([]int64, 0)
			for _, item := range v.CategoryRelationList {
				productCategoryIds = append(productCategoryIds, item.ProductCategoryId)
			}
			totalAmount := s.calcTotalAmountByproductCategoryId(cartItemList, productCategoryIds)
			if now.Before(endTime) && totalAmount > 0 && totalAmount-minPoint >= 0 {
				enableList = append(enableList, v)
			} else {
				disableList = append(disableList, v)
			}
		case 2: // 指定商品
			// 计算指定商品的总价
			productIds := make([]int64, 0)
			for _, item := range v.ProductRelationList {
				productIds = append(productIds, item.ProductId)
			}
			totalAmount := s.calcTotalAmountByProductId(cartItemList, productIds)
			if now.Before(endTime) && totalAmount > 0 && totalAmount-minPoint >= 0 {
				enableList = append(enableList, v)
			} else {
				disableList = append(disableList, v)
			}
		}
	}
	if enable == 1 {
		return enableList, nil
	}
	return disableList, nil
}

func (s *service) ListByProduct(ctx context.Context, productId int64) ([]dto.SmsCoupon, error) {
	allCouponIds := make([]int64, 0)
	// 获取指定商品优惠券
	cprQb := sms_coupon_product_relation.NewQueryBuilder().WhereProductId(mysql.EqualPredicate, productId)
	cprList, err := cprQb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	for _, v := range cprList {
		allCouponIds = append(allCouponIds, v.CouponId)
	}

	// 获取指定分类优惠券
	pQb := pms_product.NewQueryBuilder().WhereId(mysql.EqualPredicate, productId)
	product, err := pQb.First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if product.Id != 0 {
		cpcrQb := sms_coupon_product_category_relation.NewQueryBuilder().
			WhereProductCategoryId(mysql.EqualPredicate, product.ProductCategoryId)
		cpcrList, err := cpcrQb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
		if err != nil {
			return nil, err
		}
		for _, v := range cpcrList {
			allCouponIds = append(allCouponIds, v.CouponId)
		}
	}

	// 所有优惠券
	return new(dao.CouponDao).ListByCouponIds(ctx, mysql.DB().GetDbR().WithContext(ctx), allCouponIds)
}

func (s *service) List(ctx context.Context, useStatus int32) ([]dto.SmsCoupon, error) {
	currentMember, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return nil, err
	}
	return new(dao.CouponDao).GetCouponList(ctx, mysql.DB().GetDbR().WithContext(ctx), currentMember.Id, useStatus)
}

func (s *service) calcTotalAmount(cartItemList []dto.CartPromotionItem) float64 {
	total := 0.0
	for _, item := range cartItemList {
		realPrice := item.Price - item.ReduceAmount
		total += realPrice * float64(item.Quantity)
	}
	return total
}

func (s *service) calcTotalAmountByproductCategoryId(
	cartItemList []dto.CartPromotionItem, productCategoryIds []int64) float64 {
	cateIdsMap := make(map[int64]bool)
	for _, v := range productCategoryIds {
		cateIdsMap[v] = true
	}

	total := 0.0
	for _, item := range cartItemList {
		if _, ok := cateIdsMap[item.ProductCategoryId]; !ok {
			continue
		}
		realPrice := item.Price - item.ReduceAmount
		total += realPrice * float64(item.Quantity)
	}
	return total
}

func (s *service) calcTotalAmountByProductId(
	cartItemList []dto.CartPromotionItem, productIds []int64) float64 {
	prodIdsMap := make(map[int64]bool)
	for _, v := range productIds {
		prodIdsMap[v] = true
	}

	total := 0.0
	for _, item := range cartItemList {
		if _, ok := prodIdsMap[item.ProductId]; !ok {
			continue
		}
		realPrice := item.Price - item.ReduceAmount
		total += realPrice * float64(item.Quantity)
	}
	return total
}
