package pms_product_cate

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 商品分类管理Service
type Service interface {
	i()

	/**
	 * 创建商品分类
	 */
	Create(ctx context.Context, param dto.PmsProductCategoryParam) (int64, error)

	/**
	 * 修改商品分类
	 */
	Update(ctx context.Context, id int64, param dto.PmsProductCategoryParam) (int64, error)

	/**
	 * 分页获取商品分类
	 */
	List(ctx context.Context, parentId int64, pageSize, pageNum int) ([]dto.PmsProductCategory, int64, error)

	/**
	 * 删除商品分类
	 */
	Delete(ctx context.Context, id int64) (int64, error)

	/**
	 * 根据ID获取商品分类
	 */
	GetItem(ctx context.Context, id int64) (*dto.PmsProductCategory, error)

	/**
	 * 批量修改导航状态
	 */
	UpdateNavStatus(ctx context.Context, ids []int64, navStatus int32) (int64, error)

	/**
	 * 批量修改显示状态
	 */
	UpdateShowStatus(ctx context.Context, ids []int64, showStatus int32) (int64, error)

	/**
	 * 以层级形式获取商品分类
	 */
	ListWithChildren(ctx context.Context) ([]dto.PmsProductCategoryWithChildrenItem, error)
}
