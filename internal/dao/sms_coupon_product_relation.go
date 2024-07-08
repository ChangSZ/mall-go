package dao

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/sms_coupon_product_relation"

	"github.com/ChangSZ/golib/copy"
	"gorm.io/gorm"
)

type SmsCouponProductRelationDao struct{}

func (t *SmsCouponProductRelationDao) InsertList(ctx context.Context,
	tx *gorm.DB, list []dto.SmsCouponProductRelation) error {
	if len(list) == 0 {
		return nil
	}

	dataList := make([]sms_coupon_product_relation.SmsCouponProductRelation, 0, len(list))
	for _, v := range list {
		data := sms_coupon_product_relation.SmsCouponProductRelation{}
		copy.AssignStruct(&v, &data)
		data.Id = 0
		dataList = append(dataList, data)
	}
	return tx.CreateInBatches(dataList, len(dataList)).Error
}
