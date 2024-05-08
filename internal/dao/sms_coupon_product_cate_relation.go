package dao

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/sms_coupon_product_category_relation"
	"github.com/ChangSZ/mall-go/pkg/copy"
	"gorm.io/gorm"
)

type SmsCouponProductCategoryRelationDao struct{}

func (t *SmsCouponProductCategoryRelationDao) InsertList(ctx context.Context,
	tx *gorm.DB, list []dto.SmsCouponProductCategoryRelation) error {
	if len(list) == 0 {
		return nil
	}

	dataList := make([]sms_coupon_product_category_relation.SmsCouponProductCategoryRelation, 0, len(list))
	for _, v := range list {
		data := sms_coupon_product_category_relation.SmsCouponProductCategoryRelation{}
		copy.AssignStruct(&v, &data)
		data.Id = 0
		dataList = append(dataList, data)
	}
	return tx.CreateInBatches(dataList, len(dataList)).Error
}
