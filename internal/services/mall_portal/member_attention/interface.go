package member_attention

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 会员品牌关注管理Service
type Service interface {
	i()

	/**
	 * 添加关注
	 */
	Add(ctx context.Context, param dto.MemberBrandAttention) (int64, error)

	/**
	 * 取消关注
	 */
	Delete(ctx context.Context, brandId int64) (int64, error)

	/**
	 * 获取用户关注列表
	 */
	List(ctx context.Context, pageNum, pageSize int64) ([]dto.MemberBrandAttention, int64, error)

	/**
	 * 获取用户关注详情
	 */
	Detail(ctx context.Context, brandId int64) (*dto.MemberBrandAttention, error)

	/**
	 * 清空关注列表
	 */
	Clear(ctx context.Context) (int64, error)
}
