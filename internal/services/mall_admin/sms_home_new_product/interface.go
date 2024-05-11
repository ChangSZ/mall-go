package sms_home_new_product

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 首页新品推荐管理Service
type Service interface {
	i()

	/**
	 * 添加首页推荐
	 */
	Create(ctx context.Context, param []dto.SmsHomeNewProduct) (int64, error)

	/**
	 * 修改推荐排序
	 */
	UpdateSort(ctx context.Context, id int64, sort int32) (int64, error)

	/**
	 * 批量删除推荐
	 */
	Delete(ctx context.Context, ids []int64) (int64, error)

	/**
	 * 批量更新推荐状态
	 */
	UpdateRecommendStatus(ctx context.Context, ids []int64, recommendStatus int32) (int64, error)

	/**
	 * 分页查询推荐
	 */
	List(ctx context.Context, productName string, recommendStatus int32, pageSize, pageNum int) (
		[]dto.SmsHomeNewProduct, int64, error)
}
