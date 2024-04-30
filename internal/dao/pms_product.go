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
	PmsProductParam
	CateParentId int64 // 商品所选分类的父id
}

type PmsProductParam struct {
	pms_product.PmsProduct
	ProductLadderList                []pms_product_ladder.PmsProductLadder                                 `gorm:"foreignKey:ProductId"`
	ProductFullReductionList         []pms_product_full_reduction.PmsProductFullReduction                  `gorm:"foreignKey:ProductId"`
	MemberPriceList                  []pms_member_price.PmsMemberPrice                                     `gorm:"foreignKey:ProductId"`
	SkuStockList                     []pms_sku_stock.PmsSkuStock                                           `gorm:"foreignKey:ProductId"`
	ProductAttributeValueList        []pms_product_attribute_value.PmsProductAttributeValue                `gorm:"foreignKey:ProductId"`
	SubjectProductRelationList       []cms_subject_product_relation.CmsSubjectProductRelation              `gorm:"foreignKey:ProductId"`
	PrefrenceAreaProductRelationList []cms_prefrence_area_product_relation.CmsPrefrenceAreaProductRelation `gorm:"foreignKey:ProductId"`
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

func (t *PmsProductDao) GetUpdateInfo(ctx context.Context, tx *gorm.DB, id int64) (*PmsProductResult, error) {
	res := &PmsProductResult{}
	err := tx.Preload("ProductLadderList", func(db *gorm.DB) *gorm.DB {
		return db.Order("id DESC")
	}).
		Preload("ProductFullReductionList", func(db *gorm.DB) *gorm.DB {
			return db.Order("id DESC")
		}).
		Preload("MemberPriceList", func(db *gorm.DB) *gorm.DB {
			return db.Order("id DESC")
		}).
		Preload("ProductAttributeValueList", func(db *gorm.DB) *gorm.DB {
			return db.Order("id DESC")
		}).
		Preload("SubjectProductRelationList", func(db *gorm.DB) *gorm.DB {
			return db.Order("id DESC")
		}).
		Preload("PrefrenceAreaProductRelationList", func(db *gorm.DB) *gorm.DB {
			return db.Order("id DESC")
		}).
		Table("pms_product p").
		Select("p.*, pc.parent_id cate_parent_id").
		Joins("LEFT JOIN pms_product_category pc on pc.id = p.product_category_id").
		Where("p.id=?", id).Find(&res).Error
	return res, err
}
