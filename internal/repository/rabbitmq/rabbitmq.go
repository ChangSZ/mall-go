package rabbitmq

import (
	"context"

	"github.com/ChangSZ/golib/queue"

	"github.com/ChangSZ/mall-go/configs"
)

type QueueEnum struct {
	Exchange     string // 交换名称
	ExchangeType string // 交换类型
	Name         string // 队列名称
	RouteKey     string // 路由键
}

var (
	cancelQueue    *queue.RabbitMQ
	ttlCancelQueue *queue.RabbitMQ

	// 消息通知队列
	QUEUE_ORDER_CANCEL = QueueEnum{
		"mall.order.direct", "direct", "mall.order.cancel", "mall.order.cancel"}

	// 消息通知ttl队列
	QUEUE_TTL_ORDER_CANCEL = QueueEnum{
		"mall.order.direct.ttl", "direct", "mall.order.cancel.ttl", "mall.order.cancel.ttl"}
)

func GetCancelQueue() (*queue.RabbitMQ, error) {
	if cancelQueue != nil && !cancelQueue.IsClosed() {
		return cancelQueue, nil
	}
	conf := configs.Get().Rabbitmq
	rConf := &queue.RabbitMQConfig{
		Addr:        conf.Host,
		Port:        conf.Port,
		VirtualHost: conf.VirtualHost,
		User:        conf.Username,
		Password:    conf.Password,

		QueueName: QUEUE_ORDER_CANCEL.Name,
	}

	var err error
	cancelQueue, err = queue.NewRabbitMQ(context.Background(), rConf)
	return cancelQueue, err
}

func GetTtlCancelQueue() (*queue.RabbitMQ, error) {
	if ttlCancelQueue != nil && !ttlCancelQueue.IsClosed() {
		return ttlCancelQueue, nil
	}

	conf := configs.Get().Rabbitmq
	rConf := &queue.RabbitMQConfig{
		Addr:        conf.Host,
		Port:        conf.Port,
		VirtualHost: conf.VirtualHost,
		User:        conf.Username,
		Password:    conf.Password,

		QueueName: QUEUE_TTL_ORDER_CANCEL.Name,

		ExchangeName: QUEUE_ORDER_CANCEL.Exchange,
		ExchangeType: QUEUE_ORDER_CANCEL.ExchangeType,
		RoutingKey:   QUEUE_ORDER_CANCEL.RouteKey,
	}

	var err error
	ttlCancelQueue, err = queue.NewRabbitMQ(context.Background(), rConf)
	return ttlCancelQueue, err
}
