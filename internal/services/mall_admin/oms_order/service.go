package oms_order

import (
	"context"
	"fmt"
	"time"

	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/oms_order"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/oms_order_operate_history"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) List(ctx context.Context, queryParam dto.OmsOrderQueryParam, pageSize, pageNum int) (
	[]dto.OmsOrder, int64, error) {
	return new(dao.OmsOrderDao).List(ctx, mysql.DB().GetDbR().WithContext(ctx), queryParam, pageSize, pageNum)
}

func (s *service) Delivery(ctx context.Context, deliveryParamList []dto.OmsOrderDeliveryParam) (int64, error) {
	if len(deliveryParamList) == 0 {
		return 0, nil
	}
	// 批量发货
	count, err := new(dao.OmsOrderDao).Delivery(ctx, mysql.DB().GetDbW().WithContext(ctx), deliveryParamList)
	if err != nil {
		return 0, err
	}
	// 添加操作记录
	operateHistoryList := make([]oms_order_operate_history.OmsOrderOperateHistory, 0, len(deliveryParamList))
	for _, v := range deliveryParamList {
		operateHistoryList = append(operateHistoryList, oms_order_operate_history.OmsOrderOperateHistory{
			OrderId:     v.OrderId,
			OperateMan:  "后台管理员",
			OrderStatus: 2,
			Note:        "完成发货",
		})
	}
	return count, mysql.DB().GetDbW().WithContext(ctx).CreateInBatches(
		operateHistoryList, len(operateHistoryList)).Error
}

func (s *service) Close(ctx context.Context, ids []int64, note string) (int64, error) {
	data := map[string]interface{}{
		"status": 4,
	}
	qb := oms_order.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	count, err := qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
	if err != nil {
		return 0, err
	}

	operateHistoryList := make([]oms_order_operate_history.OmsOrderOperateHistory, 0, len(ids))
	for _, id := range ids {
		operateHistoryList = append(operateHistoryList, oms_order_operate_history.OmsOrderOperateHistory{
			OrderId:     id,
			OperateMan:  "后台管理员",
			OrderStatus: 4,
			Note:        "订单关闭:" + note,
		})
	}
	return count, mysql.DB().GetDbW().WithContext(ctx).CreateInBatches(
		operateHistoryList, len(operateHistoryList)).Error
}

func (s *service) Delete(ctx context.Context, ids []int64) (int64, error) {
	qb := oms_order.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	return qb.Delete(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) GetItem(ctx context.Context, id int64) (*dto.OmsOrderDetail, error) {
	return new(dao.OmsOrderDao).GetDetail(ctx, mysql.DB().GetDbR().WithContext(ctx), id)
}

func (s *service) UpdateReceiverInfo(ctx context.Context, param dto.OmsReceiverInfoParam) (int64, error) {
	data := map[string]interface{}{
		"id":                      param.OrderID,
		"receiver_name":           param.ReceiverName,
		"receiver_phone":          param.ReceiverPhone,
		"receiver_post_code":      param.ReceiverPostCode,
		"receiver_detail_address": param.ReceiverDetailAddress,
		"receiver_province":       param.ReceiverProvince,
		"receiver_city":           param.ReceiverCity,
		"receiver_region":         param.ReceiverRegion,
		"modify_time":             time.Now(),
	}
	qb := oms_order.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, param.OrderID)
	count, err := qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
	if err != nil {
		return 0, err
	}
	// 插入操作记录
	history := &oms_order_operate_history.OmsOrderOperateHistory{
		OrderId:     param.OrderID,
		OperateMan:  "后台管理员",
		OrderStatus: param.Status,
		Note:        "修改收货人信息",
	}
	if _, err := history.Create(mysql.DB().GetDbW().WithContext(ctx)); err != nil {
		return 0, fmt.Errorf("插入操作记录失败: %v", err)
	}
	return count, nil
}

func (s *service) UpdateMoneyInfo(ctx context.Context, param dto.OmsMoneyInfoParam) (int64, error) {
	data := map[string]interface{}{
		"id":              param.OrderID,
		"freight_amount":  param.FreightAmount,
		"discount_amount": param.DiscountAmount,
		"modify_time":     time.Now(),
	}
	qb := oms_order.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, param.OrderID)
	count, err := qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
	if err != nil {
		return 0, err
	}
	// 插入操作记录
	history := &oms_order_operate_history.OmsOrderOperateHistory{
		OrderId:     param.OrderID,
		OperateMan:  "后台管理员",
		OrderStatus: param.Status,
		Note:        "修改费用信息",
	}
	if _, err := history.Create(mysql.DB().GetDbW().WithContext(ctx)); err != nil {
		return 0, fmt.Errorf("插入操作记录失败: %v", err)
	}
	return count, nil
}

func (s *service) UpdateNote(ctx context.Context, id int64, note string, status int32) (int64, error) {
	data := map[string]interface{}{
		"id":          id,
		"note":        note,
		"modify_time": time.Now(),
	}
	qb := oms_order.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	count, err := qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
	if err != nil {
		return 0, err
	}
	// 插入操作记录
	history := &oms_order_operate_history.OmsOrderOperateHistory{
		OrderId:     id,
		OperateMan:  "后台管理员",
		OrderStatus: status,
		Note:        "修改备注信息:" + note,
	}
	if _, err := history.Create(mysql.DB().GetDbW().WithContext(ctx)); err != nil {
		return 0, fmt.Errorf("插入操作记录失败: %v", err)
	}
	return count, nil
}
