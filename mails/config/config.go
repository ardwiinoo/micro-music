package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Rabbit RabbitMQConfig
	Mail MailConfig
}

type RabbitMQConfig struct {
	ConnString string
}

type MailConfig struct {
	SMTPHost string
	SMTPPort string
	Username string
	Password string
	Sender   string
}

var Cfg Config

func LoadConfig(envPath ...string) error {
	path := ".env"
	if len(envPath) > 0 {
		path = envPath[0]
	}

	if err := godotenv.Load(path); err != nil {
		return err
	}

	Cfg = Config{
		Rabbit: RabbitMQConfig{
			ConnString: os.Getenv("RABBITMQ_SERVER"),
		},
		Mail: MailConfig{
			SMTPHost: os.Getenv("SMTP_HOST"),
			SMTPPort: os.Getenv("SMTP_PORT"),
			Username: os.Getenv("SMTP_USERNAME"),
			Password: os.Getenv("SMTP_PASSWORD"),
			Sender:   os.Getenv("SMTP_SENDER"),
		},
	}

	return nil
}