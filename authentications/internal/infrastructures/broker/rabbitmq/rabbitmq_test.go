package rabbitmq_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ardwiinoo/micro-music/authentications/config"
	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures/broker/rabbitmq"
)

func init() {
	if err := config.LoadConfig("../../../../.env"); err != nil {
		panic(err)
	}
}

func TestRabbitMQ(t *testing.T) {
	t.Run("RabbitMQ Connect Successfully", func(t *testing.T) {
		rabbit, err := rabbitmq.NewRabbitMQ(config.Cfg.Rabbit.ConnString)

		require.Nil(t, err)
		defer rabbit.Close()
	})

	t.Run("RabbitMQ Connection Fail", func(t *testing.T) {
		invalidConnString := "amqp://invalid:invalid@localhost:5672/"

		rabbit, err := rabbitmq.NewRabbitMQ(invalidConnString)

		require.NotNil(t, err)
		require.Nil(t, rabbit)
	})

	t.Run("RabbitMQ Publish Event Successfully", func(t *testing.T) {
		rabbit, err := rabbitmq.NewRabbitMQ(config.Cfg.Rabbit.ConnString)

		require.Nil(t, err)
		defer rabbit.Close()

		queueName := "test_events"
		eventType := "test_event"
		payload := map[string]string{
			"message": "Hello, RabbitMQ!",
		}

		err = rabbit.PublishEvent(queueName, eventType, payload)
		require.Nil(t, err)
	})

	t.Run("RabbitMQ Publish Event Fail", func(t *testing.T) {
		rabbit, err := rabbitmq.NewRabbitMQ(config.Cfg.Rabbit.ConnString)
		require.Nil(t, err)
		defer rabbit.Close()

		rabbit.Close()

		queueName := "test_events"
		eventType := "test_event"
		payload := map[string]string{
			"message": "This will fail",
		}

		err = rabbit.PublishEvent(queueName, eventType, payload)
		require.NotNil(t, err)
	})
}
