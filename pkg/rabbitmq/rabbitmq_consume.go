package rabbitmq

import (
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

// 定义消费者接口
type Consumer interface {
	Consume([]byte) //error
}

// 定义消费者单元
type consumerUnit struct {
	//启动数量
	limitNum int
	//消费者
	consume Consumer
}

// 启动RabbitMQ客户端，并初始化
func (r *RabbitMQQueueExchange) Consume() {

	//开启监听接受者接受任务
	if len(r.consumerList) > 0 {
		forever := make(chan bool)

		for i, oneConsumerUnit := range r.consumerList {
			go r.consumerHandle(oneConsumerUnit, i)
		}

		log.Println("RabbitMQ Consumer Queue Waiting for... ")

		//让程序不退出
		<-forever
	}

}

// 注册接收指定队列指定路由的数据接收者
func (r *RabbitMQQueueExchange) RegisterConsumer(addConsume Consumer, runNum int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, one := range r.consumerList {
		if addConsume == one.consume {
			log.Println("RabbitMQ RegisterConsumer already registered ")
			os.Exit(1)
		}
	}

	r.consumerList = append(r.consumerList, consumerUnit{
		limitNum: runNum,
		consume:  addConsume,
	})
}

// 监听接受者接收任务
func (r *RabbitMQQueueExchange) consumerHandle(consumerUnit consumerUnit, i int) {

	var err error

	// 声明交换机 如果存在也不会返回失败
	err = r.conn.channel.ExchangeDeclare(r.exchangeName, r.exchangeType, true, false, false, true, nil)
	if err != nil {
		log.Println("RabbitMQ Consumer channel.ExchangeDeclare error: ", err)
		return
	}

	// 声明队列
	// name： 队列名称； durable： 是否持久化，队列存盘，true服务重启后信息不会丢失，影响性能；autoDelete:是否自动删除；
	// noWait:是否非阻塞 true为是，不等待RMQ返回信息；args：参数，传nil即可；exclusive：是否设置排他,独占
	qu, err := r.conn.channel.QueueDeclare(r.queueName, true, false, false, true, nil)
	if err != nil {
		log.Println("RabbitMQ Consumer channel.QueueDeclare error: ", err)
		return
	}

	// 队列绑定
	err = r.conn.channel.QueueBind(qu.Name, r.routingKey, r.exchangeName, true, nil)
	if err != nil {
		log.Println("RabbitMQ Consumer channel.QueueBind error: ", err)
		return
	}

	//forever := make(chan bool)
	//获取消费通道，确保rabbitMQ一个一个发送消息
	err = r.conn.channel.Qos(1, 0, true)
	msgList, err := r.conn.channel.Consume(qu.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Println("RabbitMQ Consumer channel.Consume error: ", err)
		return
	}

	for msg := range msgList {

		for r.handleNum >= consumerUnit.limitNum {
			log.Printf("RabbitMQ Consumer Queue Id %v HandleNum: %v ", i, r.handleNum)
			time.Sleep(1 * time.Second)
		}

		r.handleMu.Lock()
		r.handleNum++
		r.handleMu.Unlock()
		//处理数据
		go func(r *RabbitMQQueueExchange, msg amqp.Delivery) {
			consumerUnit.consume.Consume(msg.Body)
			r.handleMu.Lock()
			r.handleNum--
			r.handleMu.Unlock()
			if r.handleNum == 0 {
				log.Println("RabbitMQ Consumer Queue Re-entry Waiting for... ")
			}
		}(r, msg)
	}

}
