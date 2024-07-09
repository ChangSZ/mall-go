package dao

import (
	"context"

	"github.com/ChangSZ/golib/copy"
	"gorm.io/gorm"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product_attribute_value"
)

type PmsProductAttributeValueDao struct{}

func (t *PmsProductAttributeValueDao) InsertList(ctx context.Context,
	tx *gorm.DB, list []dto.PmsProductAttributeValue, productId int64) error {
	if len(list) == 0 {
		return nil
	}

	dataList := make([]pms_product_attribute_value.PmsProductAttributeValue, 0, len(list))
	for _, v := range list {
		data := pms_product_attribute_value.PmsProductAttributeValue{}
		copy.AssignStruct(&v, &data)
		data.Id = 0
		data.ProductId = productId
		dataList = append(dataList, data)
	}
	return tx.CreateInBatches(dataList, len(dataList)).Error
}
