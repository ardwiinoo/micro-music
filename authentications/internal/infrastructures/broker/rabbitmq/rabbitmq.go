package rabbitmq

import (
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	connection *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQ(connString string) (rabbitMQ *RabbitMQ, err error) {
	conn, err := amqp.Dial(connString)
	if err != nil {
		return
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return
	}

	return &RabbitMQ{
		connection: conn,
		channel: ch,
	}, nil
}

func (r *RabbitMQ) Close() {
	if r.channel != nil {
		r.channel.Close()
	}

	if r.connection != nil {
		r.connection.Close()
	}
}

func (r *RabbitMQ) PublishMessage(queueName string, body string) (err error) {
	_, err = r.channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return
	}

	err = r.channel.Publish("", queueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body: []byte(body),
	})	

	return
}