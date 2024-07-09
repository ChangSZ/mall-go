package dao

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/ChangSZ/mall-go/internal/dto"
)

type OmsOrderReturnApplyDao struct{}

func (t *OmsOrderReturnApplyDao) List(ctx context.Context, tx *gorm.DB,
	queryParam dto.OmsReturnApplyQueryParam, pageSize, pageNum int) (
	[]dto.OmsOrderReturnApply, int64, error) {
	res := make([]dto.OmsOrderReturnApply, 0)
	sql := `
SELECT
	id,
	create_time,
	member_username,
	product_real_price,
	product_count,
	return_name,
	status,
	handle_time
FROM
	oms_order_return_apply
WHERE
	1 = 1
`
	if queryParam.Id != 0 {
		sql += fmt.Sprintf(" AND id = %d", queryParam.Id)
	}
	if queryParam.Status != nil {
		sql += fmt.Sprintf(" AND status = %d", *queryParam.Status)
	}
	if queryParam.HandleMan != "" {
		sql += fmt.Sprintf(" AND handle_man = %d", &queryParam.HandleMan)
	}
	if queryParam.CreateTime != "" {
		sql += fmt.Sprintf(" AND create_time LIKE \"%s\"", queryParam.CreateTime+"%")
	}
	if queryParam.HandleTime != "" {
		sql += fmt.Sprintf(" AND handle_time LIKE \"%s\"", queryParam.HandleTime+"%")
	}
	if queryParam.ReceiverKeyword != "" {
		sql += fmt.Sprintf(" AND (return_name LIKE \"%s\") OR return_phone LIKE \"%s\"",
			"%"+queryParam.ReceiverKeyword+"%", "%"+queryParam.ReceiverKeyword+"%")
	}

	var count int64
	if err := tx.Table("oms_order_return_apply").Exec(sql).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	offset := (pageNum - 1) * pageSize
	err := tx.Raw(sql).
		Limit(pageSize).
		Offset(offset).
		Scan(&res).Error
	return res, count, err
}

func (t *OmsOrderReturnApplyDao) GetDetail(ctx context.Context, tx *gorm.DB, id int64) (
	*dto.OmsOrderReturnApplyResult, error) {
	res := &dto.OmsOrderReturnApplyResult{}
	sql := `
SELECT
	ra.*, ca.id ca_id,
		  ca.address_name ca_address_name,
		  ca.name ca_name,
		  ca.phone ca_phone,
		  ca.province ca_province,
		  ca.city ca_city,
		  ca.region ca_region,
		  ca.detail_address ca_detail_address
FROM
	oms_order_return_apply ra
	LEFT JOIN oms_company_address ca ON ra.company_address_id = ca.id
WHERE ra.id=?;
`
	err := tx.Raw(sql, id).Scan(&res).Error
	return res, err
}
