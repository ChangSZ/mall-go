package dao

import (
	"context"

	"gorm.io/gorm"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product"
)

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

func (t *PmsProductDao) GetUpdateInfo(ctx context.Context, tx *gorm.DB, id int64) (*dto.PmsProductResult, error) {
	res := &dto.PmsProductResult{}
	err := tx.Preload("ProductLadderList", func(db *gorm.DB) *gorm.DB {
		return db.Order("id DESC")
	}).
		Preload("ProductFullReductionList", func(db *gorm.DB) *gorm.DB {
			return db.Order("id DESC")
		}).
		Preload("MemberPriceList", func(db *gorm.DB) *gorm.DB {
			return db.Order("id DESC")
		}).
		Preload("SkuStockList", func(db *gorm.DB) *gorm.DB {
			return db.Order("id DESC")
		}).
		Preload("ProductAttributeValueList", func(db *gorm.DB) *gorm.DB {
			return db.Order("id DESC")
		}).
		Preload("SubjectProductRelationList", func(db *gorm.DB) *gorm.DB {
			return db.Order("id ASC")
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
