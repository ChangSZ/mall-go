package member_read_history

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 会员浏览记录管理Service
type Service interface {
	i()

	/**
	 * 生成浏览记录
	 */
	Add(ctx context.Context, param dto.MemberReadHistory) (int64, error)

	/**
	 * 批量删除浏览记录
	 */
	Delete(ctx context.Context, ids []string) (int64, error)

	/**
	 * 分页获取用户浏览历史记录
	 */
	List(ctx context.Context, pageNum, pageSize int64) ([]dto.MemberReadHistory, int64, error)

	/**
	 * 清空浏览记录
	 */
	Clear(ctx context.Context) (int64, error)
}
