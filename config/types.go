package config

import "gorm.io/gorm"

type AppConfig struct {
	DB       *gorm.DB
	RestPort int
	GRPCPort int
}

type RawConfig struct {
	Server struct {
		RestPort int
		GRPCPort int
	}
	DB struct {
		DSN string
	}
}
