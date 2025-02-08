package config

import (
	"encoding/base64"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	App           AppConfig
	DB            DBConfig
	Cache         CacheConfig
	StorageConfig StorageConfig
}

type AppConfig struct {
	Name      string
	Port      int
	AppSecret Secret
}

type Secret struct {
	AppPublicKey []byte
}

type CacheConfig struct {
	Host     string
	Password string
	DB       int
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

type StorageConfig struct {
	CredentialsFile string
	BucketName      string
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

	publicKeyBase64 := os.Getenv("APP_PUBLIC_KEY")
	publicKey, _ := base64.StdEncoding.DecodeString(publicKeyBase64)

	firebaseCredBase64 := os.Getenv("FIREBASE_CREDENTIALS_BASE64")
	bucketName := os.Getenv("FIREBASE_BUCKET_NAME")

	decodedCred, _ := base64.StdEncoding.DecodeString(firebaseCredBase64)

	// write to temp file
	firebaseCred := "/tmp/serviceAccount.json"
	if err := os.WriteFile(firebaseCred, decodedCred, 0644); err != nil {
		return err
	}

	Cfg = Config{
		App: AppConfig{
			Name: os.Getenv("APP_NAME"),
			Port: port,
			AppSecret: Secret{
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
		Cache: CacheConfig{
			Host:     os.Getenv("REDIS_SERVER"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		},
		StorageConfig: StorageConfig{
			CredentialsFile: firebaseCred,
			BucketName:      bucketName,
		},
	}

	return nil
}
