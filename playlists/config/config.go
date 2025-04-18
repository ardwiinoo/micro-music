package config

import (
	"encoding/base64"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	App AppConfig
	DB  DBConfig
	Rabbit RabbitMQConfig
}

type AppConfig struct {
	Name string
	Port int
	AppSecret Secret
}

type Secret struct {
	AppPrivateKey []byte
	AppPublicKey []byte
}

type DBConfig struct {
	Host           string
	Port           int
	User           string
	Password       string
	DBName         string
	ConnectionPool DBConnectionPool
}

type DBConnectionPool struct {
	MaxIdleConnection     int
	MaxOpenConnection     int
	MaxLifetimeConnection int
	MaxIdletimeConnection int
}

type RabbitMQConfig struct {
	ConnString string
}

var Cfg Config

func LoadConfig(envPath ...string) error {

	// default
	path := ".env"
	if len(envPath) > 0 {
		path = envPath[0]
	}

	if err := godotenv.Load(path); err != nil {
		return err
	}

	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	maxIdle, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTION"))
	maxOpen, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTION"))
	maxLifetime, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTION"))
	maxIdletime, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLETIME_CONNECTION"))

	privateKeyBase64 := os.Getenv("APP_PRIVATE_KEY")
	publicKeyBase64 := os.Getenv("APP_PUBLIC_KEY")

	privateKey, _ := base64.StdEncoding.DecodeString(privateKeyBase64)
	publicKey, _ := base64.StdEncoding.DecodeString(publicKeyBase64)

	Cfg = Config{
		App: AppConfig{
			Name: os.Getenv("APP_NAME"),
			Port: port,
			AppSecret: Secret{
				AppPrivateKey: privateKey,
				AppPublicKey: publicKey,
			},
		},
		DB: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     dbPort,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
			ConnectionPool: DBConnectionPool{
				MaxIdleConnection:     maxIdle,
				MaxOpenConnection:     maxOpen,
				MaxLifetimeConnection: maxLifetime,
				MaxIdletimeConnection: maxIdletime,
			},
		},
		Rabbit: RabbitMQConfig{
			ConnString: os.Getenv("RABBITMQ_SERVER"),
		},
	}

	return nil
}