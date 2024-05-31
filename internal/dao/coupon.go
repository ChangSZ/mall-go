package dao

import (
	"context"
	"time"

	"github.com/ChangSZ/mall-go/internal/dto"

	"gorm.io/gorm"
)

type CouponDao struct{}

func (t *CouponDao) GetHistoryDetailList(ctx context.Context, tx *gorm.DB, memberId int64) (
	[]dto.SmsCouponHistoryDetail, error) {
	res := []dto.SmsCouponHistoryDetail{}
	err := tx.
		Preload("ProductRelationList", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, product_id").Order("id DESC")
		}).
		Preload("CategoryRelationList", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, product_category_id").Order("id DESC")
		}).
		Table("sms_coupon_history ch").
		Select(`ch.*, c.id c_id, c.name c_name, c.amount c_amount, c.min_point c_min_point,	c.platform c_platform,
		c.start_time c_start_time, c.end_time c_end_time, c.note c_note, c.use_type c_use_type, c.type c_type`).
		Joins("LEFT JOIN sms_coupon c ON ch.coupon_id = c.id").
		Where("ch.member_id = ? AND ch.use_status = 0", memberId).Find(&res).Error
	return res, err
}

func (t *CouponDao) GetCouponList(ctx context.Context, tx *gorm.DB, memberId int64, useStatus int32) (
	[]dto.SmsCoupon, error) {
	res := make([]dto.SmsCoupon, 0)
	db := tx.Table("sms_coupon c").
		Select("c.*").
		Joins("LEFT JOIN sms_coupon_history ch ON ch.coupon_id = c.id").
		Where("ch.member_id = ?", memberId)
	if useStatus != 2 {
		db = db.Where("ch.use_status = ? AND NOW() > c.start_time AND c.end_time > NOW()", useStatus)
	} else {
		db = db.Where("NOW() > c.end_time")
	}
	err := db.Find(&res).Error
	return res, err
}

func (t *CouponDao) ListByCouponIds(ctx context.Context, tx *gorm.DB, allCouponIds []int64) ([]dto.SmsCoupon, error) {
	res := make([]dto.SmsCoupon, 0)
	now := time.Now()
	db := tx.Table("sms_coupon")
	if len(allCouponIds) > 0 {
		db = db.Where(`(end_time > ? AND start_time < ? AND use_type = ?) OR 
(end_time > ? AND start_time < ? AND use_type != ? AND id IN ?)`,
			now, now, 0, now, now, 0, allCouponIds)
	} else {
		db = db.Where("end_time > ? AND start_time < ? AND use_type = ?", now, now, 0)
	}
	err := db.Find(&res).Error
	return res, err
}
