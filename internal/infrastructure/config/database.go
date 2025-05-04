package config

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database *sql.DB

func InitializeDatabase() error {
	destination := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		Environment.DATABASE_HOST,
		Environment.DATABASE_PORT,
		Environment.DATABASE_USER,
		Environment.DATABASE_PASSWORD,
		Environment.DATABASE_NAME,
		Environment.DATABASE_SSL_MODE,
		Environment.DATABASE_TIME_ZONE,
	)
	connection, error := gorm.Open(postgres.Open(destination), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if error != nil {
		return error
	}
	database, error := connection.DB()
	if error != nil {
		return error
	}
	database.SetMaxIdleConns(2)
	database.SetMaxOpenConns(10)
	database.SetConnMaxLifetime(5 * time.Minute)
	Database = database
	return nil
}
