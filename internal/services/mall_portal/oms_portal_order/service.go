package oms_portal_order

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/oms_order"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/oms_order_item"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/oms_order_setting"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_sku_stock"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/sms_coupon_history"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_integration_consume_setting"
	"github.com/ChangSZ/mall-go/internal/repository/redis"
	"github.com/ChangSZ/mall-go/internal/services/mall_portal/oms_cart_item"
	"github.com/ChangSZ/mall-go/internal/services/mall_portal/ums_member"
	"github.com/ChangSZ/mall-go/internal/services/mall_portal/ums_member_coupon"
	"github.com/ChangSZ/mall-go/internal/services/mall_portal/ums_member_receive_address"
	"github.com/ChangSZ/mall-go/pkg/copy"
	"github.com/ChangSZ/mall-go/pkg/math"

	"github.com/ChangSZ/golib/log"
)

var (
	REDIS_KEY_ORDER_ID string = configs.Get().Redis.Key.OrderId
	REDIS_DATABASE     string = configs.Get().Redis.Database
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) GenerateConfirmOrder(ctx context.Context, cartIds []int64) (*dto.ConfirmOrderResult, error) {
	result := &dto.ConfirmOrderResult{}
	// 获取购物车信息
	currentMember, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return nil, err
	}
	cartPromotionItemList, err := oms_cart_item.New().ListPromotion(ctx, cartIds)
	if err != nil {
		return nil, err
	}
	result.CartPromotionItemList = cartPromotionItemList

	// 获取用户收货地址列表
	memberReceiveAddressList, err := ums_member_receive_address.New().List(ctx)
	if err != nil {
		return nil, err
	}
	result.MemberReceiveAddressList = memberReceiveAddressList

	// 获取用户可用优惠券列表
	couponHistoryDetailList, err := ums_member_coupon.New().ListCart(ctx, cartPromotionItemList, 1)
	if err != nil {
		return nil, err
	}
	result.CouponHistoryDetailList = couponHistoryDetailList

	// 获取用户积分
	result.MemberIntegration = currentMember.Integration

	// 获取积分使用规则
	integrationConsumeSetting, err := ums_integration_consume_setting.NewQueryBuilder().
		WhereId(mysql.EqualPredicate, 1).
		First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if integrationConsumeSetting != nil {
		tmp := dto.UmsIntegrationConsumeSetting{}
		copy.AssignStruct(integrationConsumeSetting, &tmp)
		result.IntegrationConsumeSetting = tmp
	}

	// 计算总金额、活动优惠、应付金额
	calcAmount := s.CalcCartAmount(cartPromotionItemList)
	result.CalcAmount = *calcAmount
	return result, nil
}

func (s *service) GenerateOrder(ctx context.Context, orderParam dto.OrderParam) (*dto.Order, error) {
	// 校验收货地址
	if orderParam.MemberReceiveAddressId == 0 {
		return nil, fmt.Errorf("请选择收货地址！")
	}
	// 获取购物车及优惠信息
	currentMember, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return nil, err
	}
	cartPromotionItemList, err := oms_cart_item.New().ListPromotion(ctx, orderParam.CartIds)
	if err != nil {
		return nil, err
	}
	orderItemList := make([]dto.OmsOrderItem, 0)
	for _, cartPromotionItem := range cartPromotionItemList {
		// 生成下单商品信息
		orderItemList = append(orderItemList, dto.OmsOrderItem{
			ProductId:         cartPromotionItem.ProductId,
			ProductName:       cartPromotionItem.ProductName,
			ProductPic:        cartPromotionItem.ProductPic,
			ProductAttr:       cartPromotionItem.ProductAttr,
			ProductBrand:      cartPromotionItem.ProductBrand,
			ProductSn:         cartPromotionItem.ProductSn,
			ProductPrice:      cartPromotionItem.Price,
			ProductQuantity:   cartPromotionItem.Quantity,
			ProductSkuId:      cartPromotionItem.ProductSkuId,
			ProductSkuCode:    cartPromotionItem.ProductSkuCode,
			ProductCategoryId: cartPromotionItem.ProductCategoryId,
			PromotionAmount:   cartPromotionItem.ReduceAmount,
			PromotionName:     cartPromotionItem.PromotionMessage,
			GiftIntegration:   cartPromotionItem.Integration,
			GiftGrowth:        cartPromotionItem.Growth,
		})
	}

	// 判断购物车中商品是否都有库存
	if !s.HasStock(cartPromotionItemList) {
		return nil, fmt.Errorf("库存不足，无法下单")
	}

	// 使用优惠券
	if orderParam.CouponId != 0 {
		couponHistoryDetail, err := s.GetUseCoupon(ctx, cartPromotionItemList, orderParam.CouponId)
		if err != nil {
			return nil, err
		}
		if couponHistoryDetail == nil {
			return nil, fmt.Errorf("该优惠券不可用")
		}
		// 对下单商品的优惠券进行处理
		s.HandleCouponAmount(orderItemList, *couponHistoryDetail)
	}

	// 使用积分
	if orderParam.UseIntegration != 0 {
		totalAmount := s.CalcTotalAmount(orderItemList)
		integrationAmount := s.GetUseIntegrationAmount(ctx,
			orderParam.UseIntegration, totalAmount, currentMember, orderParam.CouponId != 0)
		if integrationAmount == 0 {
			return nil, fmt.Errorf("积分不可用")
		} else {
			// 可用情况下分摊到可用商品中
			for i, orderItem := range orderItemList {
				perAmount := math.RoundHalfEven(orderItem.ProductPrice/totalAmount, 3) * integrationAmount
				orderItemList[i].IntegrationAmount = perAmount
			}
		}
	}

	// 计算order_item的实付金额
	s.HandleRealAmount(orderItemList)

	// 进行库存锁定
	if err := s.LockStock(ctx, cartPromotionItemList); err != nil {
		return nil, err
	}

	// 根据商品合计、运费、活动优惠、优惠券、积分计算应付金额
	order := dto.OmsOrder{
		TotalAmount:     s.CalcTotalAmount(orderItemList),
		PromotionAmount: s.CalcPromotionAmount(orderItemList),
		PromotionInfo:   s.GetOrderPromotionInfo(orderItemList),
	}
	if orderParam.CouponId != 0 {
		order.CouponId = orderParam.CouponId
		order.CouponAmount = s.CalcCouponAmount(orderItemList)
	}
	if orderParam.UseIntegration != 0 {
		order.Integration = orderParam.UseIntegration
		order.IntegrationAmount = s.CalcIntegrationAmount(orderItemList)
	}

	order.PayAmount = s.CalcPayAmount(order)
	// 转化为订单信息并插入数据库
	order.MemberId = currentMember.Id
	order.CreateTime = time.Now()
	order.MemberUsername = currentMember.Username
	// 支付方式：0->未支付；1->支付宝；2->微信
	order.PayType = orderParam.PayType
	// 订单来源：0->PC订单；1->app订单
	order.SourceType = 1
	// 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	order.Status = 0
	// 订单类型：0->正常订单；1->秒杀订单
	order.OrderType = 0
	// 收货人信息：姓名、电话、邮编、地址
	address, err := ums_member_receive_address.New().GetItem(ctx, orderParam.MemberReceiveAddressId)
	if err != nil {
		return nil, err
	}
	order.ReceiverName = address.Name
	order.ReceiverPhone = address.PhoneNumber
	order.ReceiverPostCode = address.PostCode
	order.ReceiverProvince = address.Province
	order.ReceiverCity = address.City
	order.ReceiverRegion = address.Region
	order.ReceiverDetailAddress = address.DetailAddress
	// 0->未确认；1->已确认
	order.ConfirmStatus = 0
	order.DeleteStatus = 0
	// 计算赠送积分
	order.Integration = s.CalcGifIntegration(orderItemList)
	// 计算赠送成长值
	order.Growth = s.CalcGiftGrowth(orderItemList)
	//生成订单号
	order.OrderSn = s.GenerateOrderSn(order)

	// 设置自动收货天数
	orderSetting, err := oms_order_setting.NewQueryBuilder().
		WhereId(mysql.EqualPredicate, 1).
		First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if orderSetting != nil {
		order.AutoConfirmDay = orderSetting.ConfirmOvertime
	}

	// TODO: bill_*,delivery_*
	// 插入order表和order_item表
	orderData := oms_order.NewModel()
	copy.AssignStruct(&order, orderData)
	orderId, err := orderData.Create(mysql.DB().GetDbW().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	order.Id = orderId

	for i := range orderItemList {
		orderItemList[i].OrderId = order.Id
		orderItemList[i].OrderSn = order.OrderSn
	}
	if err := new(dao.OrderDao).InsertItemList(ctx,
		mysql.DB().GetDbW().WithContext(ctx), orderItemList); err != nil {
		return nil, err
	}

	// 如使用优惠券更新优惠券使用状态
	if orderParam.CouponId != 0 {
		if err := s.UpdateCouponStatus(ctx, orderParam.CouponId, currentMember.Id, 1); err != nil {
			return nil, err
		}
	}

	// 如使用积分需要扣除积分
	if orderParam.UseIntegration != 0 {
		order.UseIntegration = orderParam.UseIntegration
		if err := ums_member.New().UpdateIntegration(ctx,
			currentMember.Id, currentMember.Integration-orderParam.UseIntegration); err != nil {
			return nil, err
		}
	}

	// 删除购物车中的下单商品
	s.DeleteCartItemList(ctx, cartPromotionItemList, currentMember)

	// 发送延迟消息取消订单
	if err := s.SendDelayMessageCancelOrder(ctx, order.Id); err != nil {
		log.WithTrace(ctx).Errorf("发送延迟消息取消订单失败: %v", err)
	}

	result := &dto.Order{
		Order:         order,
		OrderItemList: orderItemList,
	}
	return result, nil
}

func (s *service) PaySuccess(ctx context.Context, orderId int64, payType int32) (int64, error) {
	data := map[string]interface{}{
		"status":       1,
		"payment_time": time.Now(),
		"pay_type":     payType,
	}
	// 只修改未付款状态的订单
	updateCount, err := oms_order.NewQueryBuilder().
		WhereId(mysql.EqualPredicate, orderId).
		WhereDeleteStatus(mysql.EqualPredicate, 0).
		WhereStatus(mysql.EqualPredicate, 0).
		Updates(mysql.DB().GetDbW().WithContext(ctx), data)
	if err != nil {
		return 0, err
	}
	if updateCount == 0 {
		return 0, fmt.Errorf("订单不存在或订单状态不是未支付！")
	}
	// 恢复所有下单商品的锁定库存，扣减真实库存
	orderDetail, err := new(dao.OmsOrderDao).GetDetail(ctx, mysql.DB().GetDbR().WithContext(ctx), orderId)
	if err != nil {
		return 0, err
	}
	var totalCount int64
	for _, orderItem := range orderDetail.OrderItemList {
		count, err := new(dao.OrderDao).ReduceSkuStock(ctx,
			mysql.DB().GetDbW().WithContext(ctx), orderItem.ProductSkuId, orderItem.ProductQuantity)
		if err != nil {
			return 0, err
		}
		totalCount += count
	}
	return totalCount, nil
}

func (s *service) CancelTimeOutOrder(ctx context.Context) (int64, error) {
	orderSetting, err := oms_order_setting.NewQueryBuilder().
		WhereId(mysql.EqualPredicate, 1).
		First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return 0, err
	}
	if orderSetting == nil {
		return 0, fmt.Errorf("未查询到订单配置信息")
	}
	// 查询超时、未支付的订单及订单详情
	timeOutOrders, err := new(dao.OrderDao).GetTimeOutOrders(ctx,
		mysql.DB().GetDbR().WithContext(ctx), orderSetting.NormalOrderOvertime)
	if err != nil {
		return 0, err
	}
	if len(timeOutOrders) == 0 {
		return 0, nil
	}

	// 修改订单状态为交易取消
	ids := make([]int64, 0)
	for _, timeOutOrder := range timeOutOrders {
		ids = append(ids, timeOutOrder.Id)
	}
	data := map[string]interface{}{"status": 4}
	_, err = oms_order.NewQueryBuilder().WhereIdIn(ids).Updates(mysql.DB().GetDbW().WithContext(ctx), data)
	if err != nil {
		return 0, err
	}
	for _, timeOutOrder := range timeOutOrders {
		// 解除订单商品库存锁定
		_, err := new(dao.OrderDao).ReleaseSkuStockLock(ctx,
			mysql.DB().GetDbR().WithContext(ctx), timeOutOrder.OrderItemList)
		if err != nil {
			return 0, err
		}

		// 修改优惠券使用状态
		if err = s.UpdateCouponStatus(ctx, timeOutOrder.CouponId, timeOutOrder.MemberId, 0); err != nil {
			return 0, err
		}

		// 返还使用积分
		if timeOutOrder.UseIntegration != 0 {
			member, err := ums_member.New().GetById(ctx, timeOutOrder.MemberId)
			if err != nil {
				return 0, err
			}
			if err = ums_member.New().UpdateIntegration(ctx,
				timeOutOrder.MemberId, member.Integration+timeOutOrder.UseIntegration); err != nil {
				return 0, err
			}
		}
	}
	return int64(len(timeOutOrders)), nil
}

// CancelOrder 取消订单
func (s *service) CancelOrder(ctx context.Context, orderId int64) error {
	// 查询未付款的取消订单
	order, err := oms_order.NewQueryBuilder().
		WhereId(mysql.EqualPredicate, orderId).
		WhereStatus(mysql.EqualPredicate, 0).
		WhereDeleteStatus(mysql.EqualPredicate, 0).
		First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return err
	}
	if order == nil {
		return nil
	}
	// 修改订单状态为取消
	data := map[string]interface{}{"status": 4}
	_, err = oms_order.NewQueryBuilder().
		WhereId(mysql.EqualPredicate, orderId).
		Updates(mysql.DB().GetDbW().WithContext(ctx), data)
	if err != nil {
		return err
	}

	orderItemList, err := oms_order_item.NewQueryBuilder().
		WhereOrderId(mysql.EqualPredicate, orderId).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return err
	}
	// 解除订单商品库存锁定
	if len(orderItemList) != 0 {
		for _, orderItem := range orderItemList {
			count, err := new(dao.OrderDao).ReleaseStockBySkuId(ctx,
				mysql.DB().GetDbW().WithContext(ctx), orderItem.ProductSkuId, orderItem.ProductQuantity)
			if err != nil {
				return err
			}
			if count == 0 {
				return fmt.Errorf("库存不足，无法释放！")
			}
		}
	}

	// 修改优惠券使用状态
	if err := s.UpdateCouponStatus(ctx, order.CouponId, order.MemberId, 0); err != nil {
		return err
	}

	// 返还使用积分
	if order.UseIntegration != 0 {
		member, err := ums_member.New().GetById(ctx, order.MemberId)
		if err != nil {
			return err
		}
		return ums_member.New().UpdateIntegration(ctx, order.MemberId, member.Integration+order.UseIntegration)
	}
	return nil
}

func (s *service) SendDelayMessageCancelOrder(ctx context.Context, orderId int64) error {
	// 获取订单超时时间
	orderSetting, err := oms_order_setting.NewQueryBuilder().
		WhereId(mysql.EqualPredicate, 1).
		First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return err
	}
	if orderSetting == nil {
		return fmt.Errorf("未查询到订单配置信息")
	}
	delayTimes := int64(orderSetting.NormalOrderOvertime) * 60 * 1000

	// 发送延迟消息
	return CancelOrderSend(ctx, orderId, delayTimes)
}

func (s *service) ConfirmReceiveOrder(ctx context.Context, orderId int64) error {
	member, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return err
	}
	order, err := oms_order.NewQueryBuilder().
		WhereId(mysql.EqualPredicate, orderId).
		First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return err
	}
	if member.Id != order.MemberId {
		return fmt.Errorf("不能确认他人订单！")
	}

	if order.Status != 2 {
		return fmt.Errorf("该订单还未发货！")
	}

	data := map[string]interface{}{
		"status":         3,
		"confirm_status": 1,
		"receive_time":   time.Now(),
	}
	_, err = oms_order.NewQueryBuilder().
		WhereId(mysql.EqualPredicate, orderId).
		Updates(mysql.DB().GetDbW().WithContext(ctx), data)
	return err
}

func (s *service) List(ctx context.Context, status int32, pageNum, pageSize int) ([]dto.OrderDetail, int64, error) {
	member, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return nil, 0, err
	}

	qb := oms_order.NewQueryBuilder().
		WhereDeleteStatus(mysql.EqualPredicate, 0).
		WhereMemberId(mysql.EqualPredicate, member.Id)
	if status != -1 {
		qb = qb.WhereStatus(mysql.EqualPredicate, status)
	}

	count, err := qb.Count(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, 0, err
	}

	offset := (pageNum - 1) * pageSize
	orderList, err := qb.
		Limit(pageSize).
		Offset(offset).
		OrderByCreateTime(false).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, count, err
	}
	if len(orderList) == 0 {
		return nil, count, nil
	}

	orderIds := make([]int64, 0)
	for _, omsOrder := range orderList {
		orderIds = append(orderIds, omsOrder.Id)
	}
	orderItemList, err := oms_order_item.NewQueryBuilder().
		WhereOrderIdIn(orderIds).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, count, err
	}
	relatedItemsMap := make(map[int64][]dto.OmsOrderItem)
	for _, orderItem := range orderItemList {
		tmp := dto.OmsOrderItem{}
		copy.AssignStruct(orderItem, &tmp)
		if _, ok := relatedItemsMap[orderItem.OrderId]; !ok {
			itemList := []dto.OmsOrderItem{tmp}
			relatedItemsMap[orderItem.OrderId] = itemList
		} else {
			relatedItemsMap[orderItem.OrderId] = append(relatedItemsMap[orderItem.OrderId], tmp)
		}
	}

	orderDetailList := make([]dto.OrderDetail, 0)
	for _, omsOrder := range orderList {
		orderDetail := dto.OrderDetail{}
		copy.AssignStruct(omsOrder, &orderDetail)
		if relatedItemList, ok := relatedItemsMap[omsOrder.Id]; ok {
			orderDetail.OrderItemList = relatedItemList
		}
		orderDetailList = append(orderDetailList, orderDetail)
	}
	return orderDetailList, count, err
}

func (s *service) Detail(ctx context.Context, orderId int64) (*dto.OrderDetail, error) {
	omsOrder, err := oms_order.NewQueryBuilder().
		WhereId(mysql.EqualPredicate, orderId).
		First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if omsOrder == nil {
		return nil, fmt.Errorf("未找到订单信息")
	}
	orderItemList, err := oms_order_item.NewQueryBuilder().
		WhereOrderId(mysql.EqualPredicate, orderId).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	detail := &dto.OrderDetail{}
	copy.AssignStruct(omsOrder, detail)
	detail.OrderItemList = make([]dto.OmsOrderItem, 0, len(orderItemList))
	for _, item := range orderItemList {
		tmp := dto.OmsOrderItem{}
		copy.AssignStruct(item, &tmp)
		detail.OrderItemList = append(detail.OrderItemList, tmp)
	}
	return detail, nil
}

func (s *service) DeleteOrder(ctx context.Context, orderId int64) error {
	member, err := ums_member.New().GetCurrentMember(ctx)
	if err != nil {
		return err
	}
	order, err := oms_order.NewQueryBuilder().
		WhereId(mysql.EqualPredicate, orderId).
		First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return err
	}
	if member.Id != order.MemberId {
		return fmt.Errorf("不能删除他人订单！")
	}
	if order.Status == 3 || order.Status == 4 {
		data := map[string]interface{}{"delete_status": 1}
		_, err := oms_order.NewQueryBuilder().
			WhereId(mysql.EqualPredicate, orderId).
			Updates(mysql.DB().GetDbW().WithContext(ctx), data)
		return err
	}
	return fmt.Errorf("只能删除已完成或已关闭的订单！")
}

func (s *service) PaySuccessByOrderSn(ctx context.Context, orderSn string, payType int32) error {
	order, err := oms_order.NewQueryBuilder().
		WhereOrderSn(mysql.EqualPredicate, orderSn).
		WhereStatus(mysql.EqualPredicate, 0).
		WhereDeleteStatus(mysql.EqualPredicate, 0).
		First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return err
	}
	if order == nil {
		return fmt.Errorf("未找到订单信息")
	}
	cnt, err := s.PaySuccess(ctx, order.Id, payType)
	if err != nil {
		log.WithTrace(ctx).Warnf("订单支付失败: %v", err)
		return err
	}
	log.WithTrace(ctx).Infof("订单支付成功, 总计商品数: %v", cnt)
	return nil
}

// GenerateOrderSn 生成18位订单编号:8位日期+2位平台号码+2位支付方式+6位以上自增id
func (s *service) GenerateOrderSn(order dto.OmsOrder) string {
	ctx := context.Background()
	var sb strings.Builder
	date := time.Now().Format("20060102")
	key := REDIS_DATABASE + ":" + REDIS_KEY_ORDER_ID + ":" + date
	increment := redis.Cache().Incr(ctx, key)
	sb.WriteString(date)
	sb.WriteString(fmt.Sprintf("%02d", order.SourceType))
	sb.WriteString(fmt.Sprintf("%02d", order.PayType))
	incrementStr := fmt.Sprintf("%06d", increment)
	if len(incrementStr) > 6 {
		sb.WriteString(incrementStr)
	} else {
		sb.WriteString(incrementStr[:6])
	}
	return sb.String()
}

// DeleteCartItemList 删除下单商品的购物车信息
func (s *service) DeleteCartItemList(ctx context.Context,
	cartPromotionItemList []dto.CartPromotionItem, currentMember *dto.UmsMember) {
	ids := make([]int64, 0)
	for _, cartPromotionItem := range cartPromotionItemList {
		ids = append(ids, cartPromotionItem.Id)
	}
	oms_cart_item.New().Delete(ctx, ids)
}

// CalcGiftGrowth 计算该订单赠送的成长值
func (s *service) CalcGiftGrowth(orderItemList []dto.OmsOrderItem) int32 {
	var sum int32
	for _, orderItem := range orderItemList {
		sum += orderItem.GiftGrowth * orderItem.ProductQuantity
	}
	return sum
}

// CalcGifIntegration 计算该订单赠送的积分
func (s *service) CalcGifIntegration(orderItemList []dto.OmsOrderItem) int32 {
	var sum int32
	for _, orderItem := range orderItemList {
		sum += orderItem.GiftIntegration * orderItem.ProductQuantity
	}
	return sum
}

/**
 * 将优惠券信息更改为指定状态
 *
 * @param couponId  优惠券id
 * @param memberId  会员id
 * @param useStatus 0->未使用；1->已使用
 */
func (s *service) UpdateCouponStatus(ctx context.Context, couponId, memberId int64, useStatus int32) error {
	if couponId == 0 {
		return nil
	}
	var useStatusQ int32 = 0
	if useStatus == 0 {
		useStatusQ = 1
	}
	couponHistory, err := sms_coupon_history.NewQueryBuilder().
		WhereMemberId(mysql.EqualPredicate, memberId).
		WhereCouponId(mysql.EqualPredicate, couponId).
		WhereUseStatus(mysql.EqualPredicate, useStatusQ).
		First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return err
	}
	if couponHistory == nil {
		return nil
	}
	data := map[string]interface{}{
		"use_time":   time.Now(),
		"use_status": useStatus,
	}
	_, err = sms_coupon_history.NewQueryBuilder().WhereId(mysql.EqualPredicate, couponHistory.Id).
		Updates(mysql.DB().GetDbW().WithContext(ctx), data)
	return err
}

func (s *service) HandleRealAmount(orderItemList []dto.OmsOrderItem) {
	for i, orderItem := range orderItemList {
		// 原价-促销优惠-优惠券抵扣-积分抵扣
		realAmount := orderItem.ProductPrice - orderItem.PromotionAmount -
			orderItem.CouponAmount - orderItem.IntegrationAmount
		orderItemList[i].RealAmount = realAmount
	}
}

// GetOrderPromotionInfo 获取订单促销信息
func (s *service) GetOrderPromotionInfo(orderItemList []dto.OmsOrderItem) string {
	promotionNameList := make([]string, 0, len(orderItemList))
	for _, orderItem := range orderItemList {
		promotionNameList = append(promotionNameList, orderItem.PromotionName)
	}
	return strings.Join(promotionNameList, ";")
}

// CalcPayAmount 计算订单应付金额
func (s *service) CalcPayAmount(order dto.OmsOrder) float64 {
	// 总金额+运费-促销优惠-优惠券优惠-积分抵扣
	return order.TotalAmount + order.FreightAmount -
		order.PromotionAmount - order.CouponAmount - order.IntegrationAmount
}

// CalcIntegrationAmount 计算订单积分金额
func (s *service) CalcIntegrationAmount(orderItemList []dto.OmsOrderItem) float64 {
	var integrationAmount float64
	for _, orderItem := range orderItemList {
		if orderItem.IntegrationAmount != 0 {
			integrationAmount += orderItem.IntegrationAmount * float64(orderItem.ProductQuantity)
		}
	}
	return integrationAmount
}

// CalcCouponAmount 计算订单优惠券金额
func (s *service) CalcCouponAmount(orderItemList []dto.OmsOrderItem) float64 {
	var couponAmount float64
	for _, orderItem := range orderItemList {
		if orderItem.CouponAmount != 0 {
			couponAmount += orderItem.CouponAmount * float64(orderItem.ProductQuantity)
		}
	}
	return couponAmount
}

// CalcPromotionAmount 计算订单活动优惠
func (s *service) CalcPromotionAmount(orderItemList []dto.OmsOrderItem) float64 {
	var promotionAmount float64
	for _, orderItem := range orderItemList {
		if orderItem.PromotionAmount != 0 {
			promotionAmount += orderItem.PromotionAmount * float64(orderItem.ProductQuantity)
		}
	}
	return promotionAmount
}

/**
 * 获取可用积分抵扣金额
 *
 * @param useIntegration 使用的积分数量
 * @param totalAmount    订单总金额
 * @param currentMember  使用的用户
 * @param hasCoupon      是否已经使用优惠券
 */
func (s *service) GetUseIntegrationAmount(ctx context.Context, useIntegration int32, totalAmount float64,
	currentMember *dto.UmsMember, hasCoupon bool) float64 {
	// 判断用户是否有这么多积分
	if useIntegration > currentMember.Integration {
		return 0
	}

	// 根据积分使用规则判断是否可用
	integrationConsumeSetting, err := ums_integration_consume_setting.NewQueryBuilder().
		WhereId(mysql.EqualPredicate, 1).
		First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		log.WithTrace(ctx).Errorf("查询积分使用规则时失败: %v", err)
		return 0
	}
	if integrationConsumeSetting == nil {
		log.WithTrace(ctx).Error("未查询到积分使用规则")
		return 0
	}

	// 是否可与优惠券共用
	if hasCoupon && integrationConsumeSetting.CouponStatus == 0 {
		return 0
	}

	// 是否达到最低使用积分门槛
	if useIntegration < integrationConsumeSetting.UseUnit {
		return 0
	}

	// 是否超过订单抵用最高百分比
	var integrationAmount float64 = math.RoundHalfEven(
		float64(useIntegration)/float64(integrationConsumeSetting.UseUnit), 2)
	var maxPercent float64 = math.RoundHalfEven(
		float64(integrationConsumeSetting.MaxPercentPerOrder)/100.0, 2)
	if integrationAmount > totalAmount*maxPercent {
		return 0
	}

	return integrationAmount
}

/**
 * 对优惠券优惠进行处理
 *
 * @param orderItemList       order_item列表
 * @param couponHistoryDetail 可用优惠券详情
 */
func (s *service) HandleCouponAmount(orderItemList []dto.OmsOrderItem, couponHistoryDetail dto.SmsCouponHistoryDetail) {
	coupon := couponHistoryDetail.Coupon
	switch coupon.UseType {
	case 0: // 全场通用
		s.CalcPerCouponAmount(orderItemList, coupon)
	case 1: // 指定分类
		couponOrderItemList := s.GetCouponOrderItemByRelation(couponHistoryDetail, orderItemList, 0)
		s.CalcPerCouponAmount(couponOrderItemList, coupon)
	case 2: // 指定商品
		couponOrderItemList := s.GetCouponOrderItemByRelation(couponHistoryDetail, orderItemList, 1)
		s.CalcPerCouponAmount(couponOrderItemList, coupon)
	}
}

/**
 * 对每个下单商品进行优惠券金额分摊的计算
 *
 * @param orderItemList 可用优惠券的下单商品商品
 */
func (s *service) CalcPerCouponAmount(orderItemList []dto.OmsOrderItem, coupon dto.SmsCoupon) {
	totalAmount := s.CalcTotalAmount(orderItemList)
	for i, orderItem := range orderItemList {
		// (商品价格/可用商品总价)*优惠券面额
		couponAmount := math.RoundHalfEven((orderItem.ProductPrice/totalAmount)*coupon.Amount, 3)
		orderItemList[i].CouponAmount = couponAmount
	}
}

/**
 * 获取与优惠券有关系的下单商品
 *
 * @param couponHistoryDetail 优惠券详情
 * @param orderItemList       下单商品
 * @param refType             使用关系类型：0->相关分类；1->指定商品
 */
func (s *service) GetCouponOrderItemByRelation(couponHistoryDetail dto.SmsCouponHistoryDetail,
	orderItemList []dto.OmsOrderItem, refType int32) []dto.OmsOrderItem {
	result := make([]dto.OmsOrderItem, 0)
	switch refType {
	case 0:
		categoryIdsMap := make(map[int64]bool, 0)
		for _, productCategoryRelation := range couponHistoryDetail.CategoryRelationList {
			categoryIdsMap[productCategoryRelation.ProductCategoryId] = true
		}
		for i, orderItem := range orderItemList {
			if _, ok := categoryIdsMap[orderItem.ProductCategoryId]; ok {
				result = append(result, orderItem)
			} else {
				orderItemList[i].CouponAmount = 0
			}
		}
	case 1:
		productIdsMap := make(map[int64]bool, 0)
		for _, productRelation := range couponHistoryDetail.ProductRelationList {
			productIdsMap[productRelation.ProductId] = true
		}
		for i, orderItem := range orderItemList {
			if _, ok := productIdsMap[orderItem.ProductId]; ok {
				result = append(result, orderItem)
			} else {
				orderItemList[i].CouponAmount = 0
			}
		}
	}
	return result
}

/**
 * 获取该用户可以使用的优惠券
 *
 * @param cartPromotionItemList 购物车优惠列表
 * @param couponId              使用优惠券id
 */
func (s *service) GetUseCoupon(ctx context.Context,
	cartPromotionItemList []dto.CartPromotionItem, couponId int64) (*dto.SmsCouponHistoryDetail, error) {
	couponHistoryDetailList, err := ums_member_coupon.New().ListCart(ctx, cartPromotionItemList, 1)
	if err != nil {
		return nil, err
	}
	for _, couponHistoryDetail := range couponHistoryDetailList {
		if couponHistoryDetail.Coupon.Id == couponId {
			return &couponHistoryDetail, nil
		}
	}
	return nil, nil
}

// CalcTotalAmount 计算总金额
func (s *service) CalcTotalAmount(orderItemList []dto.OmsOrderItem) float64 {
	var totalAmount float64
	for _, item := range orderItemList {
		totalAmount += item.ProductPrice * float64(item.ProductQuantity)
	}
	return totalAmount
}

// LockStock 锁定下单商品的所有库存
func (s *service) LockStock(ctx context.Context, cartPromotionItemList []dto.CartPromotionItem) error {
	for _, cartPromotionItem := range cartPromotionItemList {
		skuStock, err := pms_sku_stock.NewQueryBuilder().
			WhereId(mysql.EqualPredicate, cartPromotionItem.ProductSkuId).
			First(mysql.DB().GetDbR().WithContext(ctx))
		if err != nil {
			return err
		}
		if skuStock == nil {
			return fmt.Errorf("未找到商品库存信息, ProductID: %v", cartPromotionItem.ProductId)
		}
		count, err := new(dao.OrderDao).LockStockBySkuId(ctx,
			mysql.DB().GetDbW().WithContext(ctx), cartPromotionItem.ProductSkuId, cartPromotionItem.Quantity)
		if err != nil {
			return fmt.Errorf("修改库存时失败: %v", err)
		}
		if count == 0 {
			return fmt.Errorf("库存不足, 无法下单")
		}
	}
	return nil
}

// HasStock 判断下单商品是否都有库存
func (s *service) HasStock(cartPromotionItemList []dto.CartPromotionItem) bool {
	for _, cartPromotionItem := range cartPromotionItemList {
		if cartPromotionItem.RealStock <= 0 || // 判断真实库存是否≤0
			cartPromotionItem.RealStock < cartPromotionItem.Quantity { // 判断真实库存是否＜下单的数量
			return false
		}
	}
	return true
}

// CalcCartAmount 计算购物车中商品的价格
func (s *service) CalcCartAmount(cartPromotionItemList []dto.CartPromotionItem) *dto.CalcAmount {
	calcAmount := &dto.CalcAmount{}
	var totalAmount float64
	var promotionAmount float64
	for _, cartPromotionItem := range cartPromotionItemList {
		totalAmount += cartPromotionItem.Price * float64(cartPromotionItem.Quantity)
		promotionAmount += cartPromotionItem.ReduceAmount * float64(cartPromotionItem.Quantity)
	}
	calcAmount.TotalAmount = totalAmount
	calcAmount.PromotionAmount = promotionAmount
	calcAmount.PayAmount = totalAmount - promotionAmount
	return calcAmount
}
