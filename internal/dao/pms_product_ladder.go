package dao

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product_ladder"

	"github.com/ChangSZ/golib/copy"
	"gorm.io/gorm"
)

type PmsProductLadderDao struct{}

func (t *PmsProductLadderDao) InsertList(ctx context.Context,
	tx *gorm.DB, list []dto.PmsProductLadder, productId int64) error {
	if len(list) == 0 {
		return nil
	}

	dataList := make([]pms_product_ladder.PmsProductLadder, 0, len(list))
	for _, v := range list {
		data := pms_product_ladder.PmsProductLadder{}
		copy.AssignStruct(&v, &data)
		data.Id = 0
		data.ProductId = productId
		dataList = append(dataList, data)
	}
	return tx.CreateInBatches(dataList, len(dataList)).Error
}
