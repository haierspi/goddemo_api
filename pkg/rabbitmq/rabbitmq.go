package rabbitmq

import (
	"sync"
)

// 定义RabbitMQ对象
type RabbitMQQueueExchange struct {
	conn         *RabbitMQConnection
	queueName    string //队列名称
	routingKey   string //key名称
	exchangeName string //交换机名称
	exchangeType string //交换机类型
	consumerList []consumerUnit
	mu           sync.RWMutex
	handleNum    int
	handleMu     sync.RWMutex
}

// 定义队列交换机对象
type QueueExchange struct {
	QuName string //队列名称
	RtKey  string //key值
	ExName string //交换机名称
	ExType string //交换机类型
}

// 创建一个新的操作对象
func (r *RabbitMQConnection) NewQueueExchange(q *QueueExchange) *RabbitMQQueueExchange {
	return &RabbitMQQueueExchange{
		conn:         r,
		queueName:    q.QuName,
		routingKey:   q.RtKey,
		exchangeName: q.ExName,
		exchangeType: q.ExType,
	}
}
