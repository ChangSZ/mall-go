package oms_promotion

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 促销管理Service
type Service interface {
	i()

	/**
	 * 计算购物车中的促销活动信息
	 * @param cartItemList 购物车
	 */
	CalcCartPromotion(ctx context.Context, cartItemList []dto.OmsCartItem) ([]dto.CartPromotionItem, error)
}
