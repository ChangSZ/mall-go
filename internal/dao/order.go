package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/ChangSZ/golib/copy"
	"gorm.io/gorm"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/oms_order_item"
)

type OrderDao struct{}

// LockStockBySkuId 根据商品的skuId来锁定库存
func (t *OrderDao) LockStockBySkuId(ctx context.Context,
	tx *gorm.DB, productSkuId int64, quantity int32) (int64, error) {
	ret := tx.Exec(`UPDATE pms_sku_stock SET lock_stock = lock_stock + ? WHERE id = ? AND lock_stock + ? <= stock`,
		quantity, productSkuId, quantity)
	return ret.RowsAffected, ret.Error
}

// ReduceSkuStock 根据商品的skuId扣减真实库存
func (t *OrderDao) ReduceSkuStock(ctx context.Context,
	tx *gorm.DB, productSkuId int64, quantity int32) (int64, error) {
	ret := tx.Exec(`UPDATE pms_sku_stock SET lock_stock = lock_stock - ?, stock = stock - ? 
WHERE id = ? AND stock - ? >= 0 AND lock_stock - ? >= 0`,
		quantity, quantity, productSkuId, quantity, quantity)
	return ret.RowsAffected, ret.Error
}

// ReduceSkuStock 根据商品的skuId释放库存
func (t *OrderDao) ReleaseStockBySkuId(ctx context.Context,
	tx *gorm.DB, productSkuId int64, quantity int32) (int64, error) {
	ret := tx.Exec(`UPDATE pms_sku_stock SET lock_stock = lock_stock - ? WHERE id = ? AND lock_stock - ? >= 0`,
		quantity, productSkuId, quantity)
	return ret.RowsAffected, ret.Error
}

/**
 * 获取超时订单
 * @param minute 超时时间（分）
 */
func (t *OrderDao) GetTimeOutOrders(ctx context.Context,
	tx *gorm.DB, minute int32) ([]dto.OrderDetail, error) {
	res := make([]dto.OrderDetail, 0)
	timeout := time.Now().Add(time.Duration(-minute) * time.Minute)
	err := tx.
		Preload("OrderItemList", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, order_id, product_name, product_sku_id, product_sku_code, product_quantity")
		}).
		Table("oms_order o").
		Where("o.status = 0 AND o.create_time < ?", timeout).Find(&res).Error
	return res, err
}

// ReleaseSkuStockLock 解除取消订单的库存锁定
func (t *OrderDao) ReleaseSkuStockLock(ctx context.Context,
	tx *gorm.DB, orderItemList []dto.OmsOrderItem) (int64, error) {
	cases := ""
	ids := make([]int64, 0, len(orderItemList))
	for _, orderItem := range orderItemList {
		ids = append(ids, orderItem.ProductSkuId)
		cases += fmt.Sprintf("WHEN %d THEN lock_stock - %d ", orderItem.ProductSkuId, orderItem.ProductQuantity)
	}
	sql := fmt.Sprintf(`UPDATE pms_sku_stock SET lock_stock = CASE id %s END WHERE id IN ?`, cases)
	ret := tx.Exec(sql, ids)
	return ret.RowsAffected, ret.Error
}

func (t *OrderDao) InsertItemList(ctx context.Context, tx *gorm.DB, list []dto.OmsOrderItem) error {
	if len(list) == 0 {
		return nil
	}

	dataList := make([]oms_order_item.OmsOrderItem, 0, len(list))
	for _, v := range list {
		data := oms_order_item.OmsOrderItem{}
		copy.AssignStruct(&v, &data)
		data.Id = 0
		dataList = append(dataList, data)
	}
	return tx.CreateInBatches(dataList, len(dataList)).Error
}
