package dao

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/cms_prefrence_area_product_relation"

	"github.com/ChangSZ/golib/copy"
	"gorm.io/gorm"
)

type CmsPrefrenceAreaProductRelationDao struct{}

func (t *CmsPrefrenceAreaProductRelationDao) InsertList(ctx context.Context,
	tx *gorm.DB, list []dto.CmsPrefrenceAreaProductRelation, productId int64) error {
	if len(list) == 0 {
		return nil
	}

	dataList := make([]cms_prefrence_area_product_relation.CmsPrefrenceAreaProductRelation, 0, len(list))
	for _, v := range list {
		data := cms_prefrence_area_product_relation.CmsPrefrenceAreaProductRelation{}
		copy.AssignStruct(&v, &data)
		data.Id = 0
		data.ProductId = productId
		dataList = append(dataList, data)
	}
	return tx.CreateInBatches(dataList, len(dataList)).Error
}
