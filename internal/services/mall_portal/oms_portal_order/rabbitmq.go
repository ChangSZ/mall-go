package oms_portal_order

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/ChangSZ/golib/log"
	"github.com/ChangSZ/golib/queue"
	"github.com/robfig/cron/v3"

	"github.com/ChangSZ/mall-go/internal/repository/rabbitmq"
)

func CancelOrderReceive(ctx context.Context) {
	const maxRetryTimes = 3
	retryTimes := 0
	log.Info("[CancelOrderReceiver] 已启动")
	defer log.Info("[CancelOrderReceiver] 已结束")

RETRY:
	queue, err := rabbitmq.GetCancelQueue()
	if err == nil {
		log.Info("[CancelOrderReceiver] 连接rabbitmq成功")
		retryTimes = 0
	} else {
		log.Infof("[CancelOrderReceiver] 连接rabbitmq失败: %v", err)
		retryTimes++
		if retryTimes < maxRetryTimes {
			log.Infof("[CancelOrderReceiver] 准备第%v次重试", retryTimes)
			time.Sleep(3 * time.Second) // 3秒后重试
			goto RETRY
		} else {
			log.Infof("[CancelOrderReceiver] 重试%v次后仍失败, 退出", maxRetryTimes)
		}
		return
	}

	err = queue.ConsumeFuncWithCtx(ctx, func(s string) bool {
		log.Infof("[CancelOrderReceiver] process orderId: %s", s)
		orderId, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Errorf("[CancelOrderReceiver] 非预期的消息: %s", s)
			return true
		}
		if err = New().CancelOrder(ctx, orderId); err != nil {
			log.Errorf("[CancelOrderReceiver] 取消订单(%s)时失败: %v", s, err)
			return false
		}
		return true
	})
	if err != nil {
		log.Error(err)
		retryTimes++
		if retryTimes < maxRetryTimes {
			log.Infof("[CancelOrderReceiver] 准备第%v次重试", retryTimes)
			time.Sleep(3 * time.Second) // 3秒后重试
			goto RETRY
		} else {
			log.Infof("[CancelOrderReceiver] 重试%v次后仍失败, 退出", maxRetryTimes)
		}
		return
	}
}

func CancelOrderSend(ctx context.Context, orderId, delayTimes int64) error {
	q, err := rabbitmq.GetTtlCancelQueue()
	if err != nil {
		return err
	}
	opts := queue.PublishOptions{Expiration: fmt.Sprintf("%d", delayTimes)}
	return q.ProduceWithCtx(ctx, fmt.Sprintf("%d", orderId), opts)
}

func CancelTimeOutOrderCron(ctx context.Context) {
	/**
	 * cron表达式：https://crontab.guru/
	 * 每10分钟扫描一次，扫描超时未支付订单，进行取消操作
	 */
	c := cron.New()

	// 添加定时任务
	_, err := c.AddFunc("*/10 * * * *", func() {
		count, err := New().CancelTimeOutOrder(ctx)
		if err != nil {
			log.Errorf("[CancelTimeOutOrder] 取消超时订单时失败: %v", err)
			return
		}
		log.Infof("[CancelTimeOutOrder] 取消订单, 并根据sku编号释放锁定库存, 取消订单数量: %d", count)
	})

	if err != nil {
		log.Errorf("[CancelTimeOutOrder] 添加cron任务时失败: %v", err)
		return
	}

	// 启动定时任务
	c.Start()
	log.Info("[CancelTimeOutOrderCron] 已启动")
}
