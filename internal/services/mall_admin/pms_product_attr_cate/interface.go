package pms_product_attr_cate

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/pagehelper"
)

var _ Service = (*service)(nil)

// 商品属性分类管理Service
type Service interface {
	i()

	/**
	 * 创建属性分类
	 */
	Create(ctx context.Context, name string) (int64, error)

	/**
	 * 修改属性分类
	 */
	Update(ctx context.Context, id int64, name string) (int64, error)

	/**
	 * 删除属性分类
	 */
	Delete(ctx context.Context, id int64) (int64, error)

	/**
	 * 获取属性分类详情
	 */
	GetItem(ctx context.Context, id int64) (*dto.PmsProductAttributeCategory, error)

	/**
	 * 分页查询属性分类
	 */
	List(ctx context.Context, pageSize, pageNum int) (*pagehelper.ListData[dto.PmsProductAttributeCategory], error)

	/**
	 * 获取包含属性的属性分类
	 */
	ListWithAttr(ctx context.Context) ([]dto.PmsProductAttrCateItem, error)
}
