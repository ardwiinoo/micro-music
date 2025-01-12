package main

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"

	"github.com/ardwiinoo/micro-music/mails/config"
	"github.com/ardwiinoo/micro-music/mails/internal/broker/rabbitmq"
	"github.com/ardwiinoo/micro-music/mails/internal/email"
	"github.com/ardwiinoo/micro-music/mails/internal/events"
	"github.com/ardwiinoo/micro-music/mails/internal/handlers"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	conn, err := amqp.Dial(config.Cfg.Rabbit.ConnString)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	mailer := email.NewMailer(config.Cfg.Mail.SMTPHost, config.Cfg.Mail.SMTPPort, config.Cfg.Mail.Username, config.Cfg.Mail.Password, config.Cfg.Mail.Sender)

	router := events.NewEventRouter()
	router.Register("user_registered", handlers.UserRegisteredHandler(mailer))

	sub, err := rabbitmq.NewSubscriber(conn, "email_events")
	if err != nil {
		log.Fatalf("Failed to initialize subscriber: %v", err)
	}
	defer sub.Close()

	err = sub.Subscribe(func(msg []byte) error {
		var event struct {
			Type    string `json:"type"`
			Payload []byte `json:"payload"`
		}
		if err := json.Unmarshal(msg, &event); err != nil {
			return err
		}
		return router.Route(event.Type, event.Payload)
	})

	if err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}

	log.Println("Mail service is running...")
	select {}
}