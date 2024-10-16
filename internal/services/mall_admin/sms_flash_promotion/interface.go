package sms_flash_promotion

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/pagehelper"
)

var _ Service = (*service)(nil)

// 限时购活动管理Service
type Service interface {
	i()

	/**
	 * 添加活动
	 */
	Create(ctx context.Context, param dto.SmsFlashPromotion) (int64, error)

	/**
	 * 修改指定活动
	 */
	Update(ctx context.Context, id int64, param dto.SmsFlashPromotion) (int64, error)

	/**
	 * 删除单个活动
	 */
	Delete(ctx context.Context, id int64) (int64, error)

	/**
	 * 修改上下线状态
	 */
	UpdateStatus(ctx context.Context, id int64, status int32) (int64, error)

	/**
	 * 获取活动详情
	 */
	GetItem(ctx context.Context, id int64) (*dto.SmsFlashPromotion, error)

	/**
	 * 分页查询活动
	 */
	List(ctx context.Context, keyword string, pageSize, pageNum int) (
		*pagehelper.ListData[dto.SmsFlashPromotion], error)
}
