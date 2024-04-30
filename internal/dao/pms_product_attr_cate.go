package dao

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product_attribute"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product_attribute_category"

	"gorm.io/gorm"
)

type PmsProductAttrCateItem struct {
	pms_product_attribute_category.PmsProductAttributeCategory
	ProductAttributeList []pms_product_attribute.PmsProductAttribute `gorm:"foreignKey:ProductAttributeCategoryId"`
}

type PmsProductAttrCateDao struct{}

func (t *PmsProductAttrCateDao) ListWithAttr(ctx context.Context, tx *gorm.DB) ([]PmsProductAttrCateItem, error) {
	res := make([]PmsProductAttrCateItem, 0)
	err := tx.Preload("ProductAttributeList", "type=?", 1, func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, product_attribute_category_id").Order("id DESC")
	}).Table("pms_product_attribute_category").Select("id, name").Find(&res).Error
	return res, err
}
