package dao

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"

	"gorm.io/gorm"
)

type PmsProductAttributeDao struct{}

func (t *PmsProductAttributeDao) GetProductAttrInfo(ctx context.Context, tx *gorm.DB, productCategoryId int64) (
	[]dto.PmsProductAttrInfo, error) {
	res := make([]dto.PmsProductAttrInfo, 0)
	sql := `
SELECT
	pa.id  attribute_id,
	pac.id attribute_category_id
FROM
	pms_product_category_attribute_relation pcar
	LEFT JOIN pms_product_attribute pa ON pa.id = pcar.product_attribute_id
	LEFT JOIN pms_product_attribute_category pac ON pa.product_attribute_category_id = pac.id
WHERE
	pcar.product_category_id = ?
	`
	err := tx.Raw(sql, productCategoryId).Scan(&res).Error
	return res, err
}
