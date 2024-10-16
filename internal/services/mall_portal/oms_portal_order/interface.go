package oms_portal_order

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/pagehelper"
)

var _ Service = (*service)(nil)

// 前台订单管理Service
type Service interface {
	i()

	/**
	 * 根据用户购物车信息生成确认单信息
	 */
	GenerateConfirmOrder(ctx context.Context, cartIds []int64) (*dto.ConfirmOrderResult, error)

	/**
	 * 根据提交信息生成订单
	 */
	GenerateOrder(ctx context.Context, orderParam dto.OrderParam) (*dto.Order, error)

	/**
	 * 支付成功后的回调
	 */
	PaySuccess(ctx context.Context, orderId int64, payType int32) (int64, error)

	/**
	 * 自动取消超时订单
	 */
	CancelTimeOutOrder(ctx context.Context) (int64, error)

	/**
	 * 取消单个超时订单
	 */
	CancelOrder(ctx context.Context, orderId int64) error

	/**
	 * 发送延迟消息取消订单
	 */
	SendDelayMessageCancelOrder(ctx context.Context, orderId int64) error

	/**
	 * 确认收货
	 */
	ConfirmReceiveOrder(ctx context.Context, orderId int64) error

	/**
	 * 分页获取用户订单
	 */
	List(ctx context.Context, status int32, pageNum, pageSize int) (
		*pagehelper.ListData[dto.OrderDetail], error)

	/**
	 * 根据订单ID获取订单详情
	 */
	Detail(ctx context.Context, orderId int64) (*dto.OrderDetail, error)

	/**
	 * 用户根据订单ID删除订单
	 */
	DeleteOrder(ctx context.Context, orderId int64) error

	/**
	 * 根据orderSn来实现的支付成功逻辑
	 */
	PaySuccessByOrderSn(ctx context.Context, orderSn string, payType int32) error
}
