package dao

import (
	"context"

	"gorm.io/gorm"

	"github.com/ChangSZ/mall-go/internal/dto"
)

type ProductDao struct{}

// GetCartProduct 获取购物车商品信息
func (t *ProductDao) GetCartProduct(ctx context.Context, tx *gorm.DB, id int64) (
	*dto.CartProduct, error) {
	res := &dto.CartProduct{}
	err := tx.
		Preload("ProductAttributeList", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name")
		}).
		Preload("SkuStockList", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, sku_code, price, stock, pic")
		}).
		Table("pms_product p").
		Select(`p.id, p.name, p.sub_title, p.price, p.pic, p.product_attribute_category_id, p.stock`).
		Joins("LEFT JOIN pms_product_attribute pa ON p.product_attribute_category_id = pa.product_attribute_category_id").
		Order("pa.sort DESC").
		Where("p.id = ? AND pa.type = 0", id).Find(res).Error
	return res, err
}

func (t *ProductDao) GetPromotionProductList(ctx context.Context, tx *gorm.DB, ids []int64) (
	[]dto.PromotionProduct, error) {
	res := []dto.PromotionProduct{}
	err := tx.
		Preload("SkuStockList", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, product_id, sku_code, price, promotion_price, stock, lock_stock")
		}).
		Preload("ProductLadderList", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, product_id, count, discount")
		}).
		Preload("ProductFullReductionList", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, product_id, full_price, reduce_price")
		}).
		Table("pms_product p").
		Select(`p.id, p.name, p.promotion_type, p.gift_growth, p.gift_point`).
		Where("p.id IN ?", ids).Find(&res).Error
	return res, err
}
