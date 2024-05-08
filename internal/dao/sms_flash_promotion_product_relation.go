package dao

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"

	"gorm.io/gorm"
)

type SmsFlashPromotionProductRelationDao struct{}

func (t *SmsFlashPromotionProductRelationDao) GetList(ctx context.Context, tx *gorm.DB,
	flashPromotionId, flashPromotionSessionId int64, pageSize, pageNum int) (
	[]dto.SmsFlashPromotionProductRelation, int64, error) {
	res := make([]dto.SmsFlashPromotionProductRelation, 0)
	sql := `
SELECT
	r.id,
	r.flash_promotion_price,
	r.flash_promotion_count,
	r.flash_promotion_limit,
	r.flash_promotion_id,
	r.flash_promotion_session_id,
	r.product_id,
	r.sort,
	p.id p_id,
	p.name p_name,
	p.product_sn p_product_sn,
	p.price p_price,
	p.stock p_stock
FROM
	sms_flash_promotion_product_relation r
	LEFT JOIN pms_product p ON r.product_id = p.id
WHERE
	r.flash_promotion_id = ?
	AND r.flash_promotion_session_id = ?
ORDER BY r.sort DESC
`

	var count int64
	if err := tx.Raw(sql, flashPromotionId, flashPromotionSessionId).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	offset := (pageNum - 1) * pageSize
	err := tx.Raw(sql, flashPromotionId, flashPromotionSessionId).
		Limit(pageSize).
		Offset(offset).
		Scan(&res).Error
	return res, count, err
}
