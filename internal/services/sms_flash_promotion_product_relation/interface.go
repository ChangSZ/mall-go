package sms_flash_promotion_product_relation

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 限时购商品关联管理Service
type Service interface {
	i()

	/**
	 * 批量添加关联
	 */
	Create(ctx context.Context, param []dto.SmsFlashPromotionProductRelation) (int64, error)

	/**
	 * 修改关联信息
	 */
	Update(ctx context.Context, id int64, param dto.SmsFlashPromotionProductRelation) (int64, error)

	/**
	 * 删除关联
	 */
	Delete(ctx context.Context, id int64) (int64, error)

	/**
	 * 获取关联详情
	 */
	GetItem(ctx context.Context, id int64) (*dto.SmsFlashPromotionProductRelation, error)

	/**
	 * 分页查询相关商品及限时购促销信息
	 *
	 * @param flashPromotionId        限时购id
	 * @param flashPromotionSessionId 限时购场次id
	 */
	List(ctx context.Context, flashPromotionId, flashPromotionSessionId int64, pageSize, pageNum int) (
		[]dto.SmsFlashPromotionProductRelation, int64, error)

	/**
	 * 根据活动和场次id获取商品关系数量
	 * @param flashPromotionId        限时购id
	 * @param flashPromotionSessionId 限时购场次id
	 */
	GetCount(ctx context.Context, flashPromotionId, flashPromotionSessionId int64) (int64, error)
}
