package sms_coupon_history

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 优惠券领取记录管理Service
type Service interface {
	i()

	/**
	 * 分页查询优惠券领取记录
	 * @param couponId 优惠券id
	 * @param useStatus 使用状态
	 * @param orderSn 使用订单号码
	 */
	List(ctx context.Context, couponId int64, useStatus int32, orderSn string, pageSize, pageNum int) (
		[]dto.SmsCouponHistory, int64, error)
}
