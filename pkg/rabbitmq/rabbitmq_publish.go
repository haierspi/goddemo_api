package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

// 生产者 生产消息 开始
func (r *RabbitMQQueueExchange) Publish(message string) error {
	//开启监听生产者发送任务
	r.producerHandle(message)
	return nil

}

// 发送任务
func (r *RabbitMQQueueExchange) producerHandle(message string) {

	var err error

	//声明交换机 如果存在也不会返回失败
	err = r.conn.channel.ExchangeDeclare(r.exchangeName, r.exchangeType, true, false, false, true, nil)
	if err != nil {
		log.Println("RabbitMQ Producer channel.ExchangeDeclare error: ", err)
		return
	}

	//声明队列
	//name： 队列名称； durable： 是否持久化，队列存盘，true服务重启后信息不会丢失，影响性能；autoDelete:是否自动删除；
	//noWait:是否非阻塞 true为是，不等待RMQ返回信息；args：参数，传nil即可；exclusive：是否设置排他
	qu, err := r.conn.channel.QueueDeclare(r.queueName, true, false, false, true, nil)
	if err != nil {
		log.Println("RabbitMQ Producer channel.QueueDeclare error: ", err)
		return
	}

	//队列绑定
	err = r.conn.channel.QueueBind(qu.Name, r.routingKey, r.exchangeName, true, nil)
	if err != nil {
		log.Println("RabbitMQ Producer channel.QueueBind error: ", err)
		return
	}

	//发送任务消息

	err = r.conn.channel.Publish(r.exchangeName, r.routingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})

	if err != nil {
		log.Println("RabbitMQ Producer channel.Publish error: ", err)
		return
	}

}
