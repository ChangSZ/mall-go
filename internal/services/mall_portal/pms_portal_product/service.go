package pms_portal_product

import (
	"context"
	"fmt"

	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_brand"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product_attribute"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product_attribute_value"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product_category"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product_full_reduction"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product_ladder"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_sku_stock"

	"github.com/ChangSZ/golib/copy"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Search(ctx context.Context, keyword string, brandId, productCategoryId int64,
	pageNum, pageSize, sort int) ([]dto.PmsProduct, int64, error) {
	qb := pms_product.NewQueryBuilder().
		WhereDeleteStatus(mysql.EqualPredicate, 0).
		WherePublishStatus(mysql.EqualPredicate, 1)
	if keyword != "" {
		qb = qb.WhereName(mysql.LikePredicate, "%"+keyword+"%")
	}
	if brandId != 0 {
		qb = qb.WhereBrandId(mysql.EqualPredicate, brandId)
	}
	if productCategoryId != 0 {
		qb = qb.WhereProductCategoryId(mysql.EqualPredicate, productCategoryId)
	}
	// 1->按新品；2->按销量；3->价格从低到高；4->价格从高到低
	switch sort {
	case 1:
		qb = qb.OrderById(false)
	case 2:
		qb = qb.OrderBySale(false)
	case 3:
		qb = qb.OrderByPrice(true)
	case 4:
		qb = qb.OrderByPrice(false)
	}
	count, err := qb.Count(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, 0, err
	}
	offset := (pageNum - 1) * pageSize
	list, err := qb.
		Limit(pageSize).
		Offset(offset).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, 0, err
	}

	listData := make([]dto.PmsProduct, 0, len(list))
	for _, v := range list {
		tmp := dto.PmsProduct{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, count, nil
}

func (s *service) CategoryTreeList(ctx context.Context) ([]dto.PmsProductCategoryNode, error) {
	allList, err := pms_product_category.NewQueryBuilder().
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}

	listData := make([]dto.PmsProductCategory, 0, len(allList))
	for _, v := range allList {
		tmp := dto.PmsProductCategory{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}

	var result = make([]dto.PmsProductCategoryNode, 0)
	for _, item := range listData {
		if item.ParentId == 0 {
			node := s.covert(item, listData)
			result = append(result, node)
		}
	}
	return result, nil
}

// 初始对象转化为节点对象
func (s *service) covert(item dto.PmsProductCategory, allList []dto.PmsProductCategory) dto.PmsProductCategoryNode {
	node := dto.PmsProductCategoryNode{
		PmsProductCategory: item,
	}
	for _, subItem := range allList {
		if subItem.ParentId == item.Id {
			childNode := s.covert(subItem, allList)
			node.Children = append(node.Children, childNode)
		}
	}
	return node
}

func (s *service) Detail(ctx context.Context, id int64) (*dto.PmsPortalProductDetail, error) {
	result := &dto.PmsPortalProductDetail{}
	// 获取商品信息
	product, err := pms_product.NewQueryBuilder().
		WhereId(mysql.EqualPredicate, id).
		First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, fmt.Errorf("未找到该商品")
	}
	copy.AssignStruct(product, &result.Product)

	// 获取品牌信息
	brand, err := pms_brand.NewQueryBuilder().
		WhereId(mysql.EqualPredicate, product.BrandId).
		First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if brand == nil {
		return nil, fmt.Errorf("未找到商品的品牌信息")
	}
	copy.AssignStruct(brand, &result.Brand)

	// 获取商品属性信息
	productAttributeList, err := pms_product_attribute.NewQueryBuilder().
		WhereProductAttributeCategoryId(mysql.EqualPredicate, product.ProductAttributeCategoryId).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if len(productAttributeList) > 0 {
		result.ProductAttributeList = make([]dto.PmsProductAttribute, 0, len(productAttributeList))
		attributeIds := make([]int64, 0, len(productAttributeList))
		for _, v := range productAttributeList {
			attributeIds = append(attributeIds, v.Id)
			tmp := dto.PmsProductAttribute{}
			copy.AssignStruct(v, &tmp)
			result.ProductAttributeList = append(result.ProductAttributeList, tmp)
		}

		// 获取商品属性值信息
		productAttributeValueList, err := pms_product_attribute_value.NewQueryBuilder().
			WhereProductId(mysql.EqualPredicate, id).
			WhereProductAttributeIdIn(attributeIds).
			QueryAll(mysql.DB().GetDbR().WithContext(ctx))
		if err != nil {
			return nil, err
		}
		result.ProductAttributeValueList = make([]dto.PmsProductAttributeValue, 0, len(productAttributeValueList))
		for _, v := range productAttributeValueList {
			tmp := dto.PmsProductAttributeValue{}
			copy.AssignStruct(v, &tmp)
			result.ProductAttributeValueList = append(result.ProductAttributeValueList, tmp)
		}
	}

	// 获取商品SKU库存信息
	skuStockList, err := pms_sku_stock.NewQueryBuilder().
		WhereProductId(mysql.EqualPredicate, id).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	result.SkuStockList = make([]dto.PmsSkuStock, 0, len(skuStockList))
	for _, v := range skuStockList {
		tmp := dto.PmsSkuStock{}
		copy.AssignStruct(v, &tmp)
		result.SkuStockList = append(result.SkuStockList, tmp)
	}

	switch product.PromotionType {
	case 3: // 商品阶梯价格设置
		productLadderList, err := pms_product_ladder.NewQueryBuilder().
			WhereProductId(mysql.EqualPredicate, id).
			QueryAll(mysql.DB().GetDbR().WithContext(ctx))
		if err != nil {
			return nil, err
		}
		result.ProductLadderList = make([]dto.PmsProductLadder, 0, len(productLadderList))
		for _, v := range productLadderList {
			tmp := dto.PmsProductLadder{}
			copy.AssignStruct(v, &tmp)
			result.ProductLadderList = append(result.ProductLadderList, tmp)
		}
	case 4: // 商品满减价格设置
		productFullReductionList, err := pms_product_full_reduction.NewQueryBuilder().
			WhereProductId(mysql.EqualPredicate, id).
			QueryAll(mysql.DB().GetDbR().WithContext(ctx))
		if err != nil {
			return nil, err
		}
		result.ProductFullReductionList = make([]dto.PmsProductFullReduction, 0, len(productFullReductionList))
		for _, v := range productFullReductionList {
			tmp := dto.PmsProductFullReduction{}
			copy.AssignStruct(v, &tmp)
			result.ProductFullReductionList = append(result.ProductFullReductionList, tmp)
		}
	}

	// 商品可用优惠券
	couponList, err := new(dao.CouponDao).GetAvailableCouponList(
		ctx, mysql.DB().GetDbR().WithContext(ctx), id, product.ProductCategoryId)
	if err != nil {
		return nil, fmt.Errorf("查询优惠券失败")
	}
	result.CouponList = make([]dto.SmsCoupon, 0, len(couponList))
	for _, v := range couponList {
		tmp := dto.SmsCoupon{}
		copy.AssignStruct(v, &tmp)
		result.CouponList = append(result.CouponList, tmp)
	}
	return result, nil
}
