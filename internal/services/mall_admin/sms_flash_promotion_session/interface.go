package sms_flash_promotion_session

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 限时购场次管理Service
type Service interface {
	i()

	/**
	 * 添加场次
	 */
	Create(ctx context.Context, param dto.SmsFlashPromotionSession) (int64, error)

	/**
	 * 修改场次
	 */
	Update(ctx context.Context, id int64, param dto.SmsFlashPromotionSession) (int64, error)

	/**
	 * 修改场次启用状态
	 */
	UpdateStatus(ctx context.Context, id int64, status int32) (int64, error)

	/**
	 * 删除场次
	 */
	Delete(ctx context.Context, id int64) (int64, error)

	/**
	 * 获取详情
	 */
	GetItem(ctx context.Context, id int64) (*dto.SmsFlashPromotionSession, error)

	/**
	 * 根据启用状态获取场次列表
	 */
	ListAll(ctx context.Context) ([]dto.SmsFlashPromotionSession, error)

	/**
	 * 获取全部可选场次及其数量
	 */
	SelectList(ctx context.Context, flashPromotionId int64) ([]dto.SmsFlashPromotionSessionDetail, error)
}
