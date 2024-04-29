package dao

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_sku_stock"
	"github.com/ChangSZ/mall-go/pkg/copy"

	"gorm.io/gorm"
)

type PmsSkuStockDao struct{}

func (t *PmsSkuStockDao) InsertList(ctx context.Context,
	tx *gorm.DB, list []dto.PmsSkuStock, productId int64) error {
	if len(list) == 0 {
		return nil
	}

	dataList := make([]pms_sku_stock.PmsSkuStock, 0, len(list))
	for _, v := range list {
		data := pms_sku_stock.PmsSkuStock{}
		copy.AssignStruct(&v, &data)
		data.Id = 0
		data.ProductId = productId
		dataList = append(dataList, data)
	}
	return tx.CreateInBatches(dataList, len(dataList)).Error
}
