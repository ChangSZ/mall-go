package oms_portal_order_return_apply

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 前台订单退货管理Service
type Service interface {
	i()

	/**
	 * 提交申请
	 */
	Create(ctx context.Context, param dto.OmsOrderReturnApplyParam) (int64, error)
}
