package dao

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/repository/mysql/cms_prefrence_area_product_relation"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/cms_subject_product_relation"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_member_price"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product_attribute_value"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product_full_reduction"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product_ladder"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_sku_stock"

	"gorm.io/gorm"
)

type PmsProductResult struct {
	PmsProductParam `json:",inline"`
	CateParentId    int64 `json:"cateParentId"` // 商品所选分类的父id
}

type PmsProductParam struct {
	pms_product.PmsProduct           `json:",inline"`
	ProductLadderList                []pms_product_ladder.PmsProductLadder
	ProductFullReductionList         []pms_product_full_reduction.PmsProductFullReduction
	MemberPriceList                  []pms_member_price.PmsMemberPrice
	SkuStockList                     []pms_sku_stock.PmsSkuStock
	ProductAttributeValueList        []pms_product_attribute_value.PmsProductAttributeValue
	SubjectProductRelationList       []cms_subject_product_relation.CmsSubjectProductRelation
	PrefrenceAreaProductRelationList []cms_prefrence_area_product_relation.CmsPrefrenceAreaProductRelation
}

type PmsProductDao struct{}

func (t *PmsProductDao) ListByKeyword(ctx context.Context,
	tx *gorm.DB, keyword string) ([]pms_product.PmsProduct, error) {
	res := make([]pms_product.PmsProduct, 0)
	query := tx.Where("delete_status = ?", 0)
	if keyword != "" {
		query = query.Where("name LIKE ?", "%"+keyword+"%").
			Or("product_sn LIKE ?", "%"+keyword+"%")
	}
	err := query.Find(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return res, nil
}
