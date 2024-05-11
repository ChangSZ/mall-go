package oms_order_setting

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 订单设置管理Service
type Service interface {
	i()

	/**
	 * 获取指定订单设置
	 */
	GetItem(ctx context.Context, id int64) (*dto.OmsOrderSetting, error)

	/**
	 * 修改指定订单设置
	 */
	Update(ctx context.Context, id int64, param dto.OmsOrderSetting) (int64, error)
}
