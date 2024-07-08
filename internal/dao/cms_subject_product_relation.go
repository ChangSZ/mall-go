package dao

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/cms_subject_product_relation"

	"github.com/ChangSZ/golib/copy"
	"gorm.io/gorm"
)

type CmsSubjectProductRelationDao struct{}

func (t *CmsSubjectProductRelationDao) InsertList(ctx context.Context,
	tx *gorm.DB, list []dto.CmsSubjectProductRelation, productId int64) error {
	if len(list) == 0 {
		return nil
	}

	dataList := make([]cms_subject_product_relation.CmsSubjectProductRelation, 0, len(list))
	for _, v := range list {
		data := cms_subject_product_relation.CmsSubjectProductRelation{}
		copy.AssignStruct(&v, &data)
		data.Id = 0
		data.ProductId = productId
		dataList = append(dataList, data)
	}
	return tx.CreateInBatches(dataList, len(dataList)).Error
}
