package oms_order_return_reason

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 退货原因管理Service
type Service interface {
	i()

	/**
	 * 添加退货原因
	 */
	Create(ctx context.Context, param dto.OmsOrderReturnReason) (int64, error)

	/**
	 * 修改退货原因
	 */
	Update(ctx context.Context, id int64, param dto.OmsOrderReturnReason) (int64, error)

	/**
	 * 批量删除退货原因
	 */
	Delete(ctx context.Context, ids []int64) (int64, error)

	/**
	 * 分页获取退货原因
	 */
	List(ctx context.Context, pageSize, pageNum int) ([]dto.OmsOrderReturnReason, int64, error)

	/**
	 * 批量修改退货原因状态
	 */
	UpdateStatus(ctx context.Context, ids []int64, status int32) (int64, error)

	/**
	 * 获取单个退货原因详情信息
	 */
	GetItem(ctx context.Context, id int64) (*dto.OmsOrderReturnReason, error)
}
