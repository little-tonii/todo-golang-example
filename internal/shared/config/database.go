package config

import (
	"fmt"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	database *gorm.DB
	once     sync.Once
)

func InitializeDatabase() error {
	var error error
	once.Do(func() {
		destination := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
			Environment.POSTGRES_HOST,
			Environment.POSTGRES_PORT,
			Environment.POSTGRES_USER,
			Environment.POSTGRES_PASSWORD,
			Environment.POSTGRES_DB,
			Environment.POSTGRES_SSL_MODE,
			Environment.POSTGRES_TIME_ZONE,
		)
		connection, err := gorm.Open(postgres.Open(destination), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			error = err
			return
		}
		sqlDatabase, err := connection.DB()
		if err != nil {
			error = err
			return
		}
		sqlDatabase.SetMaxIdleConns(2)
		sqlDatabase.SetMaxOpenConns(10)
		sqlDatabase.SetConnMaxLifetime(5 * time.Minute)
		database = connection
	})
	return error
}

func CloseDatabase() error {
	sqlDatabase, error := database.DB()
	if error != nil {
		return error
	}
	return sqlDatabase.Close()
}

func GetDatabase() *gorm.DB {
	return database
}
