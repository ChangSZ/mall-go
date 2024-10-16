package sms_home_brand

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/pagehelper"
)

var _ Service = (*service)(nil)

// 首页品牌管理Service
type Service interface {
	i()

	/**
	 * 添加首页品牌推荐
	 */
	Create(ctx context.Context, param []dto.SmsHomeBrand) (int64, error)

	/**
	 * 修改品牌推荐排序
	 */
	UpdateSort(ctx context.Context, id int64, sort int32) (int64, error)

	/**
	 * 批量删除品牌推荐
	 */
	Delete(ctx context.Context, ids []int64) (int64, error)

	/**
	 * 批量更新推荐状态
	 */
	UpdateRecommendStatus(ctx context.Context, ids []int64, recommendStatus int32) (int64, error)

	/**
	 * 分页查询品牌推荐
	 */
	List(ctx context.Context, brandName string, recommendStatus int32, pageSize, pageNum int) (
		*pagehelper.ListData[dto.SmsHomeBrand], error)
}
