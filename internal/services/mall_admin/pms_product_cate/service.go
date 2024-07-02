package pms_product_cate

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product_category"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product_category_attribute_relation"
	"github.com/ChangSZ/mall-go/pkg/copy"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Create(ctx context.Context, param dto.PmsProductCategoryParam) (int64, error) {
	data := pms_product_category.NewModel()
	data.ProductCount = 0
	data.Level = s.getCategoryLevel(ctx, param.ParentId)
	id, err := data.Create(mysql.DB().GetDbW().WithContext(ctx))
	if err != nil {
		return 0, err
	}
	if len(param.ProductAttributeIdList) > 0 {
		if err := s.insertRelationList(ctx, id, param.ProductAttributeIdList); err != nil {
			return id, err
		}
	}
	return id, nil
}

func (s *service) Update(ctx context.Context, id int64, param dto.PmsProductCategoryParam) (int64, error) {
	data := map[string]interface{}{
		"id":           id,
		"parent_id":    param.ParentId,
		"name":         param.Name,
		"level":        s.getCategoryLevel(ctx, param.ParentId),
		"product_unit": param.ProductUnit,
		"nav_status":   param.NavStatus,
		"show_status":  param.ShowStatus,
		"sort":         param.Sort,
		"icon":         param.Icon,
		"keywords":     param.Keywords,
		"description":  param.Description,
	}

	// 更新商品分类时要更新商品中的名称
	{
		qb := pms_product.NewQueryBuilder()
		qb = qb.WhereProductCategoryId(mysql.EqualPredicate, id)
		data := map[string]interface{}{
			"name": param.Name,
		}
		qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
	}

	// 删除属性信息
	{
		qb := pms_product_category_attribute_relation.NewQueryBuilder()
		qb = qb.WhereProductCategoryId(mysql.EqualPredicate, id)
		if _, err := qb.Delete(mysql.DB().GetDbW().WithContext(ctx)); err != nil {
			return 0, err
		}
	}

	// 新增属性信息
	if len(param.ProductAttributeIdList) > 0 {
		s.insertRelationList(ctx, id, param.ProductAttributeIdList)
	}

	qb := pms_product_category.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) List(ctx context.Context, parentId int64, pageSize, pageNum int) ([]dto.PmsProductCategory, int64, error) {
	qb := pms_product_category.NewQueryBuilder()
	qb = qb.WhereParentId(mysql.EqualPredicate, parentId)
	count, err := qb.Count(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, 0, err
	}

	offset := (pageNum - 1) * pageSize
	list, err := qb.
		Limit(pageSize).
		Offset(offset).
		OrderBySort(false).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, 0, err
	}

	listData := make([]dto.PmsProductCategory, 0, len(list))
	for _, v := range list {
		tmp := dto.PmsProductCategory{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, count, err
}

func (s *service) Delete(ctx context.Context, id int64) (int64, error) {
	qb := pms_product_category.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Delete(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) GetItem(ctx context.Context, id int64) (*dto.PmsProductCategory, error) {
	qb := pms_product_category.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	data, err := qb.First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	res := &dto.PmsProductCategory{}
	copy.AssignStruct(data, res)
	return res, nil
}

func (s *service) UpdateNavStatus(ctx context.Context, ids []int64, navStatus int32) (int64, error) {
	data := map[string]interface{}{
		"nav_status": navStatus,
	}
	qb := pms_product_category.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) UpdateShowStatus(ctx context.Context, ids []int64, showStatus int32) (int64, error) {
	data := map[string]interface{}{
		"show_status": showStatus,
	}
	qb := pms_product_category.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}
func (s *service) ListWithChildren(ctx context.Context) ([]dto.PmsProductCategoryWithChildrenItem, error) {
	return new(dao.PmsProductCateDao).ListWithChildren(ctx, mysql.DB().GetDbR().WithContext(ctx))
}

/**
 * 根据分类的parentId获取分类的level
 */
func (s *service) getCategoryLevel(ctx context.Context, parentId int64) int32 {
	if parentId == 0 {
		return 0
	} else {
		qb := pms_product_category.NewQueryBuilder()
		qb = qb.WhereParentId(mysql.EqualPredicate, parentId)
		data, _ := qb.First(mysql.DB().GetDbR().WithContext(ctx))
		if data != nil {
			return data.Level + 1
		}
		return 0
	}
}

/**
 * 批量插入商品分类与筛选属性关系表
 * @param productCategoryId 商品分类id
 * @param productAttributeIdList 相关商品筛选属性id集合
 */
func (s *service) insertRelationList(ctx context.Context, productCategoryId int64, productAttributeIdList []int64) error {
	relationList := make([]pms_product_category_attribute_relation.PmsProductCategoryAttributeRelation, 0)
	for _, productAttrId := range productAttributeIdList {
		relationList = append(relationList, pms_product_category_attribute_relation.PmsProductCategoryAttributeRelation{
			ProductCategoryId:  productCategoryId,
			ProductAttributeId: productAttrId,
		})
	}
	return mysql.DB().GetDbW().WithContext(ctx).CreateInBatches(relationList, len(relationList)).Error
}
