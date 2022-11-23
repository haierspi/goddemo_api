package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

var (
	Dsn string //rabbitmq 连接
)

// rabbitmq连接对象
type RabbitMQConnection struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewRdabbitMQCon(c string) (*RabbitMQConnection, error) {
	Dsn = c
	mq := &RabbitMQConnection{}
	return mq.MqConnect()
}

// 链接rabbitMQ
func (r *RabbitMQConnection) MqConnect() (*RabbitMQConnection, error) {
	var err error
	RabbitUrl := Dsn
	mqConn, err := amqp.Dial(RabbitUrl)
	//defer mqConn.Close()
	r.connection = mqConn

	if err != nil {
		log.Println("rabbitMQ connection Error: ", err)
		return nil, err
	}
	mqChan, err := mqConn.Channel()
	r.channel = mqChan //赋值给RabbitMQ对象
	if err != nil {
		log.Println("rabbitMQ Channel Error: ", err)
		return nil, err
	}
	return r, nil
}

// 关闭rabbitMQ链接
func (r *RabbitMQConnection) MqClose() error {

	//先关闭管道，在关闭链接
	err := r.channel.Close()
	if err != nil {
		log.Println("rabbitMQ channel.Close Error: ", err)
		return err
	}
	r.channel = nil
	err = r.connection.Close()
	if err != nil {
		log.Println("rabbitMQ connection.Close Error: ", err)
		return err
	}
	r.connection = nil
	return nil
}
