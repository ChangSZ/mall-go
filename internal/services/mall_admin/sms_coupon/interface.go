package sms_coupon

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/pagehelper"
)

var _ Service = (*service)(nil)

// 优惠券管理Service
type Service interface {
	i()

	/**
	 * 添加优惠券
	 */
	Create(ctx context.Context, param dto.SmsCouponParam) (int64, error)

	/**
	 * 根据优惠券id删除优惠券
	 */
	Delete(ctx context.Context, id int64) (int64, error)

	/**
	 * 根据优惠券id更新优惠券信息
	 */
	Update(ctx context.Context, id int64, param dto.SmsCouponParam) (int64, error)

	/**
	 * 分页获取优惠券列表
	 */
	List(ctx context.Context, name string, couponType int32, pageSize, pageNum int) (
		*pagehelper.ListData[dto.SmsCoupon], error)

	/**
	 * 获取优惠券详情
	 * @param id 优惠券表id
	 */
	GetItem(ctx context.Context, id int64) (*dto.SmsCouponParam, error)
}
