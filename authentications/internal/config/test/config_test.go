package test

import (
	"log"
	"testing"

	"github.com/ardwiinoo/micro-music/authentications/internal/config"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	
	t.Run(("LoadConfig should load the config file"), func(t *testing.T) {
		fileName := "../../../config.yaml"
		err := config.LoadConfig(fileName)

		require.Nil(t, err)
		log.Printf("Config: %v\n", config.Cfg)
	})

	t.Run(("LoadConfig should return an error if the file does not exist"), func(t *testing.T) {
		fileName := "config.yaml"
		err := config.LoadConfig(fileName)

		require.NotNil(t, err)
	})
}