package dao

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"

	"gorm.io/gorm"
)

type PmsProductCateDao struct{}

func (t *PmsProductCateDao) ListWithChildren(ctx context.Context, tx *gorm.DB) (
	[]dto.PmsProductCategoryWithChildrenItem, error) {
	res := make([]dto.PmsProductCategoryWithChildrenItem, 0)
	err := tx.Preload("Children", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, parent_id").Order("id DESC")
	}).Table("pms_product_category").Select("id, name").Find(&res).Error
	return res, err
}
