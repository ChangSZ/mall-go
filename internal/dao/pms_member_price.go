package dao

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_member_price"
	"github.com/ChangSZ/mall-go/pkg/copy"

	"gorm.io/gorm"
)

type PmsMemberPriceDao struct{}

func (t *PmsMemberPriceDao) InsertList(ctx context.Context,
	tx *gorm.DB, list []dto.PmsMemberPrice, productId int64) error {
	if len(list) == 0 {
		return nil
	}

	dataList := make([]pms_member_price.PmsMemberPrice, 0, len(list))
	for _, v := range list {
		data := pms_member_price.PmsMemberPrice{}
		copy.AssignStruct(&v, &data)
		data.Id = 0
		data.ProductId = productId
		dataList = append(dataList, data)
	}
	return tx.CreateInBatches(dataList, len(dataList)).Error
}
