package oms_cart_item

import (
	"context"
	"fmt"
	"time"

	"github.com/ChangSZ/golib/copy"

	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/oms_cart_item"
	"github.com/ChangSZ/mall-go/internal/services/mall_portal/oms_promotion"
	"github.com/ChangSZ/mall-go/internal/services/mall_portal/ums_member"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Add(ctx context.Context, param dto.OmsCartItem) (int64, error) {
	currentMember, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return 0, err
	}
	cartItem := &oms_cart_item.OmsCartItem{}
	copy.AssignStruct(&param, cartItem)
	cartItem.MemberId = currentMember.Id
	cartItem.MemberNickname = currentMember.Nickname
	cartItem.DeleteStatus = 0
	existCartItem, err := s.GetCartItem(ctx, cartItem)
	if err != nil {
		return 0, err
	}
	// 不存在时进行新建
	if existCartItem == nil {
		cartItem.CreateDate = time.Now()
		return cartItem.Create(mysql.DB().GetDbW().WithContext(ctx))
	}

	// 存在则进行更新
	data := map[string]interface{}{
		"modify_date": time.Now(),
		"quantity":    existCartItem.Quantity + cartItem.Quantity,
	}
	return oms_cart_item.NewQueryBuilder().WhereId(mysql.EqualPredicate, existCartItem.Id).
		Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

// GetCartItem 根据会员id,商品id和规格获取购物车中商品
func (s *service) GetCartItem(ctx context.Context, cartItem *oms_cart_item.OmsCartItem) (
	*oms_cart_item.OmsCartItem, error) {
	qb := oms_cart_item.NewQueryBuilder().
		WhereMemberId(mysql.EqualPredicate, cartItem.MemberId).
		WhereProductId(mysql.EqualPredicate, cartItem.ProductId).
		WhereDeleteStatus(mysql.EqualPredicate, 0)
	if cartItem.ProductSkuId != 0 {
		qb = qb.WhereProductSkuId(mysql.EqualPredicate, cartItem.ProductSkuId)
	}
	cartItemList, err := qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if len(cartItemList) > 0 {
		return cartItemList[0], nil
	}
	return nil, nil
}

func (s *service) List(ctx context.Context) ([]dto.OmsCartItem, error) {
	currentMember, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return nil, err
	}
	list, err := oms_cart_item.NewQueryBuilder().
		WhereDeleteStatus(mysql.EqualPredicate, 0).
		WhereMemberId(mysql.EqualPredicate, currentMember.Id).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}

	listData := make([]dto.OmsCartItem, 0, len(list))
	for _, v := range list {
		tmp := dto.OmsCartItem{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, nil
}

func (s *service) ListPromotion(ctx context.Context, cartIds []int64) ([]dto.CartPromotionItem, error) {
	cartItemList, err := s.List(ctx)
	if err != nil {
		return nil, err
	}
	if len(cartIds) != 0 {
		idsMap := make(map[int64]bool)
		for _, cartId := range cartIds {
			idsMap[cartId] = true
		}

		newCartItemList := make([]dto.OmsCartItem, 0)
		for _, cartItem := range cartItemList {
			if _, ok := idsMap[cartItem.Id]; ok {
				newCartItemList = append(newCartItemList, cartItem)
			}
		}
		cartItemList = newCartItemList
	}

	cartPromotionItemList := make([]dto.CartPromotionItem, 0)
	if len(cartItemList) != 0 {
		cartPromotionItemList, err = oms_promotion.New().CalcCartPromotion(ctx, cartItemList)
		if err != nil {
			return nil, err
		}
	}
	return cartPromotionItemList, nil
}

func (s *service) UpdateQuantity(ctx context.Context, id int64, quantity int32) (int64, error) {
	currentMember, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return 0, err
	}

	data := map[string]interface{}{
		"quantity": quantity,
	}
	return oms_cart_item.NewQueryBuilder().
		WhereDeleteStatus(mysql.EqualPredicate, 0).
		WhereId(mysql.EqualPredicate, id).
		WhereMemberId(mysql.EqualPredicate, currentMember.Id).
		Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) Delete(ctx context.Context, ids []int64) (int64, error) {
	currentMember, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return 0, err
	}

	data := map[string]interface{}{
		"delete_status": 1,
	}
	return oms_cart_item.NewQueryBuilder().
		WhereIdIn(ids).
		WhereMemberId(mysql.EqualPredicate, currentMember.Id).
		Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) GetCartProduct(ctx context.Context, productId int64) (*dto.CartProduct, error) {
	return new(dao.ProductDao).GetCartProduct(ctx, mysql.DB().GetDbR().WithContext(ctx), productId)
}

func (s *service) UpdateAttr(ctx context.Context, param dto.OmsCartItem) (int64, error) {
	data := map[string]interface{}{
		"modify_date":   time.Now(),
		"delete_status": 1,
	}
	cnt, err := oms_cart_item.NewQueryBuilder().
		WhereId(mysql.EqualPredicate, param.Id).
		Updates(mysql.DB().GetDbW().WithContext(ctx), data)
	if err != nil {
		return 0, err
	}
	if cnt != 0 {
		return 0, fmt.Errorf("更新失败")
	}
	param.Id = 0
	return s.Add(ctx, param)
}

func (s *service) Clear(ctx context.Context) (int64, error) {
	currentMember, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return 0, err
	}

	data := map[string]interface{}{
		"delete_status": 1,
	}
	return oms_cart_item.NewQueryBuilder().
		WhereMemberId(mysql.EqualPredicate, currentMember.Id).
		Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}
