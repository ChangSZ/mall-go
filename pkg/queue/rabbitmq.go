package queue

import (
	"context"
	"fmt"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConfig struct {
	Addr        string
	Port        int64
	VirtualHost string
	User        string
	Password    string

	QueueName string

	ExchangeName string
	ExchangeType string
	RoutingKey   string
}

type RabbitMQ struct {
	sync.Once
	delivery <-chan amqp.Delivery
	conn     *amqp.Connection
	channel  *amqp.Channel
	queue    *amqp.Queue
}

type PublishOptions struct {
	Expiration  string
	ContentType string
}

func NewRabbitMQ(ctx context.Context, config *RabbitMQConfig) (*RabbitMQ, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/%s",
		config.User, config.Password, config.Addr, config.Port, config.VirtualHost)
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	if err := ch.Qos(1, 0, false); err != nil {
		return nil, err
	}
	// 使用Publish方法的Confirm字段设置为true时，可等待服务器确认消息是否成功发送，否则不需要等待
	if err = ch.Confirm(false); err != nil {
		return nil, err
	}

	if config.ExchangeType == "" {
		config.ExchangeType = "direct"
	}

	args := make(amqp.Table)
	if config.RoutingKey != "" {
		args["x-dead-letter-routing-key"] = config.RoutingKey // 到期后转发的路由键
	}
	if config.ExchangeName != "" {
		// Declare the order exchange
		err = ch.ExchangeDeclare(
			config.ExchangeName, // name
			config.ExchangeType, // type
			true,                // durable
			false,               // auto-deleted
			false,               // internal
			false,               // no-wait
			nil,                 // arguments
		)
		if err != nil {
			return nil, err
		}
		args["x-dead-letter-exchange"] = config.ExchangeName // 到期后转发的交换机
	}

	q, err := ch.QueueDeclare(
		config.QueueName, // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		args,             // arguments
	)
	queue := &q
	if err != nil {
		return nil, err
	}

	if config.ExchangeName != "" {
		// Bind the order queue to the exchange
		err = ch.QueueBind(
			config.QueueName,    // queue name
			config.RoutingKey,   // routing key
			config.ExchangeName, // exchange
			false,
			nil,
		)
		if err != nil {
			return nil, err
		}
	}

	obj := &RabbitMQ{
		conn:    conn,
		channel: ch,
		queue:   queue,
	}
	return obj, nil
}

// Close 关闭connection和channel
func (q *RabbitMQ) Close() error {
	if err := q.channel.Close(); err != nil {
		return fmt.Errorf("failed to close channel: %w", err)
	}
	if err := q.conn.Close(); err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}
	return nil
}

// IsClosed returns true if the channel or connection is marked as closed, otherwise false is returned.
func (q *RabbitMQ) IsClosed() bool {
	return q.channel.IsClosed() || q.conn.IsClosed()
}

// Delete 删除队列
func (q *RabbitMQ) Delete() error {
	_, err := q.channel.QueueDelete(q.queue.Name, false, false, false)
	if err != nil {
		fmt.Println("failed to delete queue: ", q.queue.Name, ", err: ", err)
		return fmt.Errorf("无法删除队列: %v", err)
	}
	return nil
}

// ProduceWithCtx 生产消息
func (q *RabbitMQ) ProduceWithCtx(ctx context.Context, value string, options PublishOptions) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if options.ContentType == "" {
			options.ContentType = "text/plain"
		}

		// Publish the message with the provided options
		err := q.channel.PublishWithContext(
			ctx,
			"",
			q.queue.Name,
			false,
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  options.ContentType,
				Body:         []byte(value),
				Expiration:   options.Expiration,
			},
		)
		if err != nil {
			return fmt.Errorf("failed to publish message: %w", err)
		}
		return nil
	}
}

func (q *RabbitMQ) setDelivery() error {
	var err error
	q.Once.Do(func() {
		q.delivery, err = q.channel.Consume(
			q.queue.Name,
			"",
			false,
			false,
			false,
			false,
			nil,
		)
	})
	return err
}

// ConsumeWithCtx 每次消费1个消息
func (q *RabbitMQ) ConsumeWithCtx(ctx context.Context) (string, error) {
	if err := q.setDelivery(); err != nil {
		return "", err
	}

	timer := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case err := <-q.conn.NotifyClose(make(chan *amqp.Error)):
			return "", fmt.Errorf("connection.NotifyClose: %w", err)
		case err := <-q.channel.NotifyClose(make(chan *amqp.Error)):
			return "", fmt.Errorf("channel.NotifyClose: %w", err)
		case msg := <-q.delivery:
			// 确认消息
			if err := msg.Ack(false); err != nil {
				return "", err
			}
			return string(msg.Body), nil
		case <-timer.C:
			if q.IsClosed() {
				return "", fmt.Errorf("可能由于网络原因断连了, conn: %v,  chan: %v", q.conn.IsClosed(), q.channel.IsClosed())
			}
		}
	}

}

func (q *RabbitMQ) ConsumeFuncWithCtx(ctx context.Context, f func(string) bool) error {
	if err := q.setDelivery(); err != nil {
		return err
	}

	timer := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-q.conn.NotifyClose(make(chan *amqp.Error)):
			return fmt.Errorf("connection.NotifyClose: %w", err)
		case err := <-q.channel.NotifyClose(make(chan *amqp.Error)):
			return fmt.Errorf("channel.NotifyClose: %w", err)
		case msg := <-q.delivery:
			// 调用回调函数处理消息，如果返回true，则确认消息，否则拒绝消息
			if f(string(msg.Body)) {
				if err := msg.Ack(false); err != nil {
					return err
				}
			} else {
				if err := msg.Nack(false, true); err != nil {
					return err
				}
			}
		case <-timer.C:
			if q.IsClosed() {
				return fmt.Errorf("可能由于网络原因断连了, conn: %v,  chan: %v", q.conn.IsClosed(), q.channel.IsClosed())
			}
		}
	}
}
