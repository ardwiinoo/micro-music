package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type Subscriber struct {
	channel *amqp.Channel
	queue string
}

func NewSubscriber(conn *amqp.Connection, queue string) (*Subscriber, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return &Subscriber{channel: ch, queue: queue}, nil
}

func (s *Subscriber) Subscribe(processFunc func([]byte) error) error {
	msgs, err := s.channel.Consume(s.queue, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			if err := processFunc(d.Body); err != nil {
				log.Printf("Failed to process message: %v", err)
			}
		}
	}()

	return nil
}

func (s *Subscriber) Close() {
	s.channel.Close()
}