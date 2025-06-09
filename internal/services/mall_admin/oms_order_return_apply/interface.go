package oms_order_return_apply

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 退货申请管理Service
type Service interface {
	i()

	/**
	 * 分页查询申请
	 */
	List(ctx context.Context, queryParam dto.OmsReturnApplyQueryParam, pageSize, pageNum int) (
		[]dto.OmsOrderReturnApply, int64, error)

	/**
	 * 批量删除申请
	 */
	Delete(ctx context.Context, ids []int64) (int64, error)

	/**
	 * 修改指定申请状态
	 */
	UpdateStatus(ctx context.Context, id int64, param dto.OmsUpdateStatusParam) (int64, error)

	/**
	 * 获取指定申请详情
	 */
	GetItem(ctx context.Context, id int64) (*dto.OmsOrderReturnApplyResult, error)
}
