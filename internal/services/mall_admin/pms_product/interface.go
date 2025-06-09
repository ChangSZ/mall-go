package pms_product

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 商品管理Service
type Service interface {
	i()

	/**
	 * 创建商品
	 */
	Create(ctx context.Context, param dto.PmsProductParam) (int64, error)

	/**
	 * 根据商品ID获取商品信息（用于更新商品）
	 */
	GetUpdateInfo(ctx context.Context, id int64) (*dto.PmsProductResult, error)

	/**
	 * 更新商品
	 */
	Update(ctx context.Context, id int64, param dto.PmsProductParam) (int64, error)

	/**
	 * 分页查询商品
	 */
	List(ctx context.Context, queryParam dto.PmsProductQueryParam, pageSize, pageNum int) (
		[]dto.PmsProduct, int64, error)

	/**
	 * 批量修改审核状态
	 * @param ids 商品ID列表
	 * @param verifyStatus 审核状态
	 * @param detail 审核详情
	 */
	UpdateVerifyStatus(ctx context.Context, ids []int64, verifyStatus int32, detail string) (int64, error)

	/**
	 * 批量修改商品上架状态
	 */
	UpdatePublishStatus(ctx context.Context, ids []int64, publishStatus int32) (int64, error)

	/**
	 * 批量修改商品推荐状态
	 */
	UpdateRecommendStatus(ctx context.Context, ids []int64, recommendStatus int32) (int64, error)

	/**
	 * 批量修改新品状态
	 */
	UpdateNewStatus(ctx context.Context, ids []int64, newStatus int32) (int64, error)

	/**
	 * 批量删除商品
	 */
	UpdateDeleteStatus(ctx context.Context, ids []int64, deleteStatus int32) (int64, error)

	/**
	 * 根据商品名称或者货号模糊查询
	 */
	SimpleList(ctx context.Context, keyword string) ([]dto.PmsProduct, error)
}
