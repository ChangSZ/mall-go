package pms_product_attr

import (
	"context"

	"github.com/ChangSZ/golib/copy"
	"github.com/ChangSZ/golib/log"

	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product_attribute"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product_attribute_category"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) List(ctx context.Context, cid int64, attrType int32, pageSize, pageNum int) (
	[]dto.PmsProductAttribute, int64, error) {
	qb := pms_product_attribute.NewQueryBuilder()
	qb = qb.WhereProductAttributeCategoryId(mysql.EqualPredicate, cid)
	qb = qb.WhereType(mysql.EqualPredicate, attrType)
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

	listData := make([]dto.PmsProductAttribute, 0, len(list))
	for _, v := range list {
		tmp := dto.PmsProductAttribute{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, count, err
}

func (s *service) Create(ctx context.Context, param dto.PmsProductAttrParam) (int64, error) {
	data := pms_product_attribute.NewModel()
	copy.AssignStruct(&param, data)
	id, err := data.Create(mysql.DB().GetDbW().WithContext(ctx))
	if err != nil {
		return 0, err
	}

	// 新增商品属性以后需要更新商品属性分类数量
	qb := pms_product_attribute_category.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, data.ProductAttributeCategoryId)
	pmsProductAttributeCategory, err := qb.First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return id, err
	}
	updateData := map[string]interface{}{}
	if data.Type == 0 {
		updateData["attribute_count"] = pmsProductAttributeCategory.AttributeCount + 1
	} else if data.Type == 1 {
		updateData["param_count"] = pmsProductAttributeCategory.ParamCount + 1
	}
	if len(updateData) > 0 {
		if _, err := qb.Updates(mysql.DB().GetDbW().WithContext(ctx), updateData); err != nil {
			log.WithTrace(ctx).Warnf("更新商品属性分类数量时失败: %v", err)
			return id, err
		}
	}
	return id, nil
}

func (s *service) Update(ctx context.Context, id int64, param dto.PmsProductAttrParam) (int64, error) {
	data := map[string]interface{}{
		"id":                            id,
		"product_attribute_category_id": param.ProductAttributeCategoryId,
		"name":                          param.Name,
		"select_type":                   param.SelectType,
		"input_type":                    param.InputType,
		"input_list":                    param.InputList,
		"sort":                          param.Sort,
		"filter_type":                   param.FilterType,
		"search_type":                   param.SearchType,
		"related_status":                param.RelatedStatus,
		"hand_add_status":               param.HandAddStatus,
		"type":                          param.Type,
	}

	qb := pms_product_attribute.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) GetItem(ctx context.Context, id int64) (*dto.PmsProductAttribute, error) {
	qb := pms_product_attribute.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	data, err := qb.First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	res := &dto.PmsProductAttribute{}
	copy.AssignStruct(data, res)
	return res, nil
}

func (s *service) Delete(ctx context.Context, ids []int64) (int64, error) {
	if len(ids) == 0 {
		return 0, nil
	}
	// 获取分类
	var data *pms_product_attribute.PmsProductAttribute
	var pmsProductAttributeCategory *pms_product_attribute_category.PmsProductAttributeCategory
	var err error
	{
		qb := pms_product_attribute.NewQueryBuilder()
		qb = qb.WhereId(mysql.EqualPredicate, ids[0])
		data, err = qb.First(mysql.DB().GetDbR().WithContext(ctx))
		if err != nil {
			return 0, err
		}

		cateQb := pms_product_attribute_category.NewQueryBuilder()
		cateQb = cateQb.WhereId(mysql.EqualPredicate, data.ProductAttributeCategoryId)
		pmsProductAttributeCategory, err = cateQb.First(mysql.DB().GetDbR().WithContext(ctx))
		if err != nil {
			return 0, err
		}
	}

	qb := pms_product_attribute.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	cnt, err := qb.Delete(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return cnt, err
	}

	updateData := map[string]interface{}{}
	if data.Type == 0 {
		if int64(pmsProductAttributeCategory.AttributeCount) > cnt {
			updateData["attribute_count"] = int64(pmsProductAttributeCategory.AttributeCount) - cnt
		} else {
			updateData["attribute_count"] = 0
		}
	} else if data.Type == 1 {
		if int64(pmsProductAttributeCategory.ParamCount) > cnt {
			updateData["param_count"] = int64(pmsProductAttributeCategory.ParamCount) - cnt
		} else {
			updateData["param_count"] = 0
		}
	}
	if len(updateData) > 0 {
		if _, err := qb.Updates(mysql.DB().GetDbW().WithContext(ctx), updateData); err != nil {
			log.WithTrace(ctx).Warnf("更新商品属性分类数量时失败: %v", err)
			return cnt, err
		}
	}
	return cnt, nil
}

func (s *service) GetProductAttrInfo(ctx context.Context, productCategoryId int64) ([]dto.PmsProductAttrInfo, error) {
	return new(dao.PmsProductAttributeDao).GetProductAttrInfo(
		ctx, mysql.DB().GetDbR().WithContext(ctx), productCategoryId)
}
