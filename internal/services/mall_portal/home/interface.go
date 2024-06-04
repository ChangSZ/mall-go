package home

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 首页内容管理Service
type Service interface {
	i()

	/**
	 * 获取首页内容
	 */
	Content(ctx context.Context) (*dto.HomeContentResult, error)

	/**
	 * 首页商品推荐
	 */
	RecommendProductList(ctx context.Context, pageNum, pageSize int) ([]dto.PmsProduct, error)

	/**
	 * 获取商品分类
	 * @param parentId 0:获取一级分类；其他：获取指定二级分类
	 */
	GetProductCateList(ctx context.Context, parentId int64) ([]dto.PmsProductCategory, error)

	/**
	 * 根据专题分类分页获取专题
	 * @param cateId 专题分类id
	 */
	GetSubjectList(ctx context.Context, cateId int64, pageNum, pageSize int) ([]dto.CmsSubject, error)

	/**
	 * 分页获取人气推荐商品
	 */
	HotProductList(ctx context.Context, pageNum, pageSize int) ([]dto.PmsProduct, error)

	/**
	 * 分页获取新品推荐商品
	 */
	NewProductList(ctx context.Context, pageNum, pageSize int) ([]dto.PmsProduct, error)
}
