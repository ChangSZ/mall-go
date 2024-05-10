package dao

import (
	"context"
	"fmt"

	"github.com/ChangSZ/mall-go/internal/dto"

	"gorm.io/gorm"
)

type OmsOrderDao struct{}

func (t *OmsOrderDao) List(ctx context.Context, tx *gorm.DB,
	queryParam dto.OmsOrderQueryParam, pageSize, pageNum int) (
	[]dto.OmsOrder, int64, error) {
	res := make([]dto.OmsOrder, 0)
	sql := "SELECT * FROM oms_order WHERE delete_status = 0"
	if queryParam.OrderSn != "" {
		sql += fmt.Sprintf(" AND order_sn = %s", queryParam.OrderSn)
	}
	if queryParam.Status != nil {
		sql += fmt.Sprintf(" AND status = %d", *queryParam.Status)
	}
	if queryParam.SourceType != nil {
		sql += fmt.Sprintf(" AND source_type = %d", *queryParam.SourceType)
	}
	if queryParam.OrderType != nil {
		sql += fmt.Sprintf(" AND order_type = %d", *queryParam.OrderType)
	}
	if queryParam.CreateTime != "" {
		sql += fmt.Sprintf(" AND create_time LIKE \"%s\"", queryParam.CreateTime+"%")
	}
	if queryParam.ReceiverKeyword != "" {
		sql += fmt.Sprintf(" AND (receiver_name LIKE \"%s\") OR receiver_phone LIKE \"%s\"",
			"%"+queryParam.ReceiverKeyword+"%", "%"+queryParam.ReceiverKeyword+"%")
	}

	var count int64
	if err := tx.Table("oms_order").Exec(sql).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	offset := (pageNum - 1) * pageSize
	err := tx.Raw(sql).
		Limit(pageSize).
		Offset(offset).
		Scan(&res).Error
	return res, count, err
}

func (t *OmsOrderDao) Delivery(ctx context.Context, tx *gorm.DB, deliveryParamList []dto.OmsOrderDeliveryParam) (
	int64, error) {
	var (
		orderIds          []int64
		deliverySns       []string
		deliveryCompanies []string
	)

	// 构建订单ID、物流单号和物流公司的切片
	for _, delivery := range deliveryParamList {
		orderIds = append(orderIds, delivery.OrderId)
		deliveryCompanies = append(deliveryCompanies, delivery.DeliveryCompany)
		deliverySns = append(deliverySns, delivery.DeliverySn)
	}

	// 执行原始 SQL
	sql := `UPDATE oms_order
				SET
					delivery_sn = CASE id
						` + t.buildCaseClause(orderIds, deliverySns) + `,
					delivery_company = CASE id
						` + t.buildCaseClause(orderIds, deliveryCompanies) + `,
					delivery_time = CASE id
						` + t.buildNowCaseClause(orderIds) + `,
					status = CASE id
						` + t.buildStatusCaseClause(orderIds) + `
				WHERE
					id IN (?) AND status = 1`

	ret := tx.Exec(sql, orderIds)
	return ret.RowsAffected, ret.Error
}

// 构建 CASE 子句
func (t *OmsOrderDao) buildCaseClause(ids []int64, values []string) string {
	var clause string
	for i, id := range ids {
		clause += fmt.Sprintf("WHEN %d THEN %s ", id, values[i])
	}
	return clause + "END"
}

// 构建 CASE 子句（处理 delivery_time）
func (t *OmsOrderDao) buildNowCaseClause(ids []int64) string {
	var clause string
	for _, id := range ids {
		clause += fmt.Sprintf("WHEN %d THEN now() ", id)
	}
	return clause + "END"
}

// 构建 CASE 子句（处理 status）
func (t *OmsOrderDao) buildStatusCaseClause(ids []int64) string {
	var clause string
	for _, id := range ids {
		clause += fmt.Sprintf("WHEN %d THEN 2 ", id)
	}
	return clause + "END"
}

func (t *OmsOrderDao) GetDetail(ctx context.Context, tx *gorm.DB, id int64) (*dto.OmsOrderDetail, error) {
	res := &dto.OmsOrderDetail{}
	err := tx.
		Preload("OrderItemList", func(db *gorm.DB) *gorm.DB {
			return db.Order("id DESC")
		}).
		Preload("HistoryList", func(db *gorm.DB) *gorm.DB {
			return db.Order("id DESC")
		}).
		Table("oms_order").
		Where("id=?", id).Find(&res).Error
	return res, err
}
