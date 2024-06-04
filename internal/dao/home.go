package dao

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"

	"gorm.io/gorm"
)

type HomeDao struct{}

func (t *HomeDao) GetRecommendBrandList(ctx context.Context, tx *gorm.DB, pageNum, pageSize int) ([]dto.PmsBrand, error) {
	res := make([]dto.PmsBrand, 0)
	offset := (pageNum - 1) * pageSize
	err := tx.Table("pms_brand b").
		Select("b.*").
		Joins("LEFT JOIN sms_home_brand hb ON hb.brand_id = b.id").
		Where("hb.recommend_status = 1 AND b.show_status = 1").
		Order("hb.sort DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&res).
		Error
	return res, err
}
