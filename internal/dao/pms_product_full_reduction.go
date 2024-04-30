package dao

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product_full_reduction"
	"github.com/ChangSZ/mall-go/pkg/copy"

	"gorm.io/gorm"
)

type PmsProductFullReductionDao struct{}

func (t *PmsProductFullReductionDao) InsertList(ctx context.Context,
	tx *gorm.DB, list []dto.PmsProductFullReduction, productId int64) error {
	if len(list) == 0 {
		return nil
	}

	dataList := make([]pms_product_full_reduction.PmsProductFullReduction, 0, len(list))
	for _, v := range list {
		data := pms_product_full_reduction.PmsProductFullReduction{}
		copy.AssignStruct(&v, &data)
		data.Id = 0
		data.ProductId = productId
		dataList = append(dataList, data)
	}
	return tx.CreateInBatches(dataList, len(dataList)).Error
}
