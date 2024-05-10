package pms_brand

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 商品品牌管理Service
type Service interface {
	i()

	/**
	 * 创建品牌
	 */
	Create(ctx context.Context, param dto.PmsBrandParam) (int64, error)

	/**
	 * 修改品牌
	 */
	Update(ctx context.Context, id int64, param dto.PmsBrandParam) (int64, error)

	/**
	 * 删除品牌
	 */
	Delete(ctx context.Context, id int64) (int64, error)

	/**
	 * 批量删除品牌
	 */
	DeleteBatch(ctx context.Context, ids []int64) (int64, error)

	/**
	 * 获取所有品牌
	 */
	ListAll(ctx context.Context) ([]dto.PmsBrand, error)

	/**
	 * 分页查询品牌
	 */
	List(ctx context.Context, keyword string, showStatus int32, pageSize, pageNum int) (
		[]dto.PmsBrand, int64, error)

	/**
	 * 获取品牌详情
	 */
	GetItem(ctx context.Context, id int64) (*dto.PmsBrand, error)

	/**
	 * 修改显示状态
	 */
	UpdateShowStatus(ctx context.Context, ids []int64, showStatus int32) (int64, error)

	/**
	 * 修改厂家制造商状态
	 */
	UpdateFactoryStatus(ctx context.Context, ids []int64, factoryStatus int32) (int64, error)
}
