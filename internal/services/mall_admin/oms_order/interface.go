package oms_order

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 订单管理Service
type Service interface {
	i()

	/**
	 * 分页查询订单
	 */
	List(ctx context.Context, queryParam dto.OmsOrderQueryParam, pageSize, pageNum int) ([]dto.OmsOrder, int64, error)

	/**
	 * 批量发货
	 */
	Delivery(ctx context.Context, deliveryParamList []dto.OmsOrderDeliveryParam) (int64, error)

	/**
	 * 批量关闭订单
	 */
	Close(ctx context.Context, ids []int64, note string) (int64, error)

	/**
	 * 批量删除订单
	 */
	Delete(ctx context.Context, ids []int64) (int64, error)

	/**
	 * 获取指定订单详情
	 */
	GetItem(ctx context.Context, id int64) (*dto.OmsOrderDetail, error)

	/**
	 * 修改订单收货人信息
	 */
	UpdateReceiverInfo(ctx context.Context, param dto.OmsReceiverInfoParam) (int64, error)

	/**
	 * 修改订单费用信息
	 */
	UpdateMoneyInfo(ctx context.Context, param dto.OmsMoneyInfoParam) (int64, error)

	/**
	 * 修改订单备注
	 */
	UpdateNote(ctx context.Context, id int64, note string, status int32) (int64, error)
}
