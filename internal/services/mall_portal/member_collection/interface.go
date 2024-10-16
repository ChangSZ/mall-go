package member_collection

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/pagehelper"
)

var _ Service = (*service)(nil)

// 会员商品收藏管理Service
type Service interface {
	i()

	/**
	 * 添加收藏
	 */
	Add(ctx context.Context, param dto.MemberProductCollection) (int64, error)

	/**
	 * 删除收藏
	 */
	Delete(ctx context.Context, productId int64) (int64, error)

	/**
	 * 分页查询收藏
	 */
	List(ctx context.Context, pageNum, pageSize int64) (
		*pagehelper.ListData[dto.MemberProductCollection], error)

	/**
	 * 查看收藏详情
	 */
	Detail(ctx context.Context, productId int64) (*dto.MemberProductCollection, error)

	/**
	 * 清空收藏
	 */
	Clear(ctx context.Context) (int64, error)
}
