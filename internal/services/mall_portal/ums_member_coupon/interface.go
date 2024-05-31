package ums_member_coupon

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 会员优惠券管理Service
type Service interface {
	i()

	/**
	 * 会员添加优惠券
	 */
	Add(ctx context.Context, couponId int64) error

	/**
	 * 获取优惠券历史列表
	 */
	ListHistory(ctx context.Context, useStatus int32) ([]dto.SmsCouponHistory, error)

	/**
	 * 根据购物车信息获取可用优惠券
	 */
	ListCart(ctx context.Context, enable int32) ([]dto.SmsCouponHistoryDetail, error)

	/**
	 * 获取当前商品相关优惠券
	 */
	ListByProduct(ctx context.Context, productId int64) ([]dto.SmsCoupon, error)

	/**
	 * 获取用户优惠券列表
	 */
	List(ctx context.Context, useStatus int32) ([]dto.SmsCoupon, error)
}
