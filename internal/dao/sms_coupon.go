package dao

import (
	"context"

	"gorm.io/gorm"

	"github.com/ChangSZ/mall-go/internal/dto"
)

type SmsCouponDao struct{}

func (t *SmsCouponDao) GetItem(ctx context.Context, tx *gorm.DB, id int64) (*dto.SmsCouponParam, error) {
	res := &dto.SmsCouponParam{}
	err := tx.Preload("ProductRelationList", func(db *gorm.DB) *gorm.DB {
		return db.Order("id DESC")
	}).
		Preload("ProductCategoryRelationList", func(db *gorm.DB) *gorm.DB {
			return db.Order("id DESC")
		}).
		Table("sms_coupon c").
		Where("c.id=?", id).Find(&res).Error
	return res, err
}
