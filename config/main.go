package config

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
)

var Config AppConfig

type AppConfig struct {
	DB *gorm.DB
}

func GetConfig() AppConfig {
	bindEnvVars([]string{})

	var config AppConfig

	db, err := GormOpen()
	if err != nil {
		// todo
	}

	config.DB = db

	return config
}

func bindEnvVars(vars []string) {
	for _, v := range vars {
		err := viper.BindEnv(v)
		if err != nil {
			log.Fatalf("unable to bind '%v' env var", v)
		}
	}
}
