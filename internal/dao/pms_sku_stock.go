package dao

import (
	"context"

	"github.com/ChangSZ/golib/copy"
	"gorm.io/gorm"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_sku_stock"
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

func (t *PmsSkuStockDao) ReplaceList(ctx context.Context,
	tx *gorm.DB, list []dto.PmsSkuStock) (int64, error) {
	var rowsAffected int64
	db := tx.Begin()
	for _, v := range list {
		data := map[string]interface{}{
			"id":              v.Id,
			"product_id":      v.ProductId,
			"sku_code":        v.SkuCode,
			"price":           v.Price,
			"stock":           v.Stock,
			"low_stock":       v.LowStock,
			"pic":             v.Pic,
			"sale":            v.Sale,
			"promotion_price": v.PromotionPrice,
			"lock_stock":      v.LockStock,
			"sp_data":         v.SpData,
		}
		ret := db.Model(pms_sku_stock.NewModel()).
			Table("pms_sku_stock").
			Where(map[string]interface{}{"id": v.Id}).
			Save(data) // TODO: 此处是REPLACE INTO操作
		if ret.Error != nil {
			db.Rollback()
			return 0, ret.Error
		}
		rowsAffected += ret.RowsAffected
	}
	if err := db.Commit().Error; err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
