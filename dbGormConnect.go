package PostgresTypeEnv

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func ConnectPostgres(config ConfigEnv) *gorm.DB {
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		config.PostgresUrl, config.PostgresPort, config.PostgresUsername, config.PostgresPassword, config.PostgresDb, config.SslMode, config.TimeZone)

	dbClient, err := gorm.Open(postgres.Open(cnn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic("error 1")
	}
	sqlDb, _ := dbClient.DB()
	err = sqlDb.Ping()
	if err != nil {
		panic("error 2")
	}
	sqlDb.SetMaxIdleConns(config.SetMaxIdleConns)
	sqlDb.SetMaxOpenConns(config.SetMaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(config.SetConnMaxLifetime) * time.Hour)

	return dbClient
}
