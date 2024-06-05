package pms_product_attr

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 商品属性管理Service
type Service interface {
	i()

	/**
	 * 根据分类ID和类型分页获取商品属性
	 * @param cid 分类id
	 * @param type 0->规格；1->参数
	 */

	List(ctx context.Context, cid int64, attrType int32, pageSize, pageNum int) ([]dto.PmsProductAttribute, int64, error)

	/**
	 * 添加商品属性
	 */
	Create(ctx context.Context, param dto.PmsProductAttrParam) (int64, error)

	/**
	 * 修改商品属性
	 */
	Update(ctx context.Context, id int64, param dto.PmsProductAttrParam) (int64, error)

	/**
	 * 获取单个商品属性信息
	 */
	GetItem(ctx context.Context, id int64) (*dto.PmsProductAttribute, error)

	/**
	 * 批量删除商品属性
	 */
	Delete(ctx context.Context, ids []int64) (int64, error)

	/**
	 * 获取商品分类对应属性列表
	 */
	GetProductAttrInfo(ctx context.Context, productCategoryId int64) ([]dto.PmsProductAttrInfo, error)
}
