package config

import (
	"errors"
	"github.com/gin-gonic/gin"
	cgl "github.com/mathandcrypto/cryptomath-gorm-logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
	"time"
)

var Config AppCfg

type AppCfg struct {
	DB       *gorm.DB
	RestPort int
	GRPCPort int

	ServerLogLevel string
	Logger         *logrus.Logger
}

// Setup ... setups AppCfg; The logic of this function is not implemented as init()
// for the opportunity to reload configuration in runtime by triggering it somehow.
func Setup() error {
	// List of mandatory env variables that should be specified
	bindEnvVars([]string{})

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading checks files, %s", err)
	}

	// default values
	//
	viper.SetDefault("ports.rest", 8080)
	viper.SetDefault("ports.grpc", 9090)

	viper.SetDefault("server.logLevel", "INFO")

	viper.SetDefault("db.logLevel", "SILENT")
	//
	//

	// configure logging
	//
	appLogger := logrus.New()

	gormLoggerConfig := cgl.Config{
		SlowThreshold:         time.Second,
		SkipErrRecordNotFound: true,
	}

	switch viper.GetString("server.logLevel") {
	case "DEBUG":
		gin.SetMode(gin.DebugMode)
		logrus.SetLevel(logrus.DebugLevel)
		gormLoggerConfig.LogLevel = logger.Info
	case "WARN":
		gin.SetMode(gin.ReleaseMode)
		logrus.SetLevel(logrus.WarnLevel)
		gormLoggerConfig.LogLevel = logger.Info
	case "INFO":
		gin.SetMode(gin.ReleaseMode)
		logrus.SetLevel(logrus.InfoLevel)
		gormLoggerConfig.LogLevel = logger.Silent
	default:
		gin.SetMode(gin.ReleaseMode)
		logrus.SetLevel(logrus.ErrorLevel)
		gormLoggerConfig.LogLevel = logger.Silent
	}

	appLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "02/Jan/2006:15:04:05 -0700",
		FullTimestamp:   true,
	})
	appLogger.SetOutput(os.Stdout)
	//
	//

	dsn := viper.GetString("db.dsn")
	if dsn == "" {
		return errors.New("DB DSN has not been specified via YAML config or ENV variables")
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // autoconfigure based on currently MySQL version
	}), &gorm.Config{
		Logger: cgl.New(appLogger, gormLoggerConfig),
	})

	if err != nil {
		return err
	}

	Config = AppCfg{
		RestPort: viper.GetInt("ports.rest"),
		GRPCPort: viper.GetInt("ports.grpc"),
		DB:       db,
		Logger:   appLogger,
	}

	return nil
}

func bindEnvVars(vars []string) {
	for _, v := range vars {
		err := viper.BindEnv(v)
		if err != nil {
			log.Fatalf("unable to bind '%v' env var", v)
		}
	}
}
