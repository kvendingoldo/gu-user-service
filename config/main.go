package config

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Config AppConfig

func init() {
	bindEnvVars([]string{})

	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.SetConfigType("yml")

	var cfg RawConfig

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading checks files, %s", err)
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	Config = AppConfig{
		RestPort: cfg.Server.RestPort,
		GRPCPort: cfg.Server.GRPCPort,
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       cfg.DB.DSN, // data source name
		DefaultStringSize:         256,        // default size for string fields
		DisableDatetimePrecision:  true,       // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,       // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,       // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,      // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		// TODO: Add later
	}

	Config.DB = db
}

func bindEnvVars(vars []string) {
	for _, v := range vars {
		err := viper.BindEnv(v)
		if err != nil {
			log.Fatalf("unable to bind '%v' env var", v)
		}
	}
}
