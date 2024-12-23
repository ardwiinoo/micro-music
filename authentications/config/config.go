package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App AppConfig `yaml:"app"`
	DB  DBConfig  `yaml:"db"`
}

type AppConfig struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}

type DBConfig struct {
	Host           string           `yaml:"host"`
	Port           int              `yaml:"port"`
	User           string           `yaml:"user"`
	Password       string           `yaml:"password"`
	DBName         string           `yaml:"dbname"`
	ConnectionPool DBConnectionPool `yaml:"connection_pool"`
}

type DBConnectionPool struct {
	MaxIdleConnection     uint8 `yaml:"max_idle_connection"`
	MaxOpenConnection     uint8 `yaml:"max_open_connection"`
	MaxLifetimeConnection uint8 `yaml:"max_lifetime_connection"`
	MaxIdletimeConnection uint8 `yaml:"max_idletime_connection"`
}

var Cfg Config

func LoadConfig(fileName string) (err error) {
	configByte, err := os.ReadFile(fileName)

	if err != nil {
		return err
	}

	return yaml.Unmarshal(configByte, &Cfg)
}