package config

import (
	"errors"
	"os"
	"strconv"
)

type DatabaseConfig struct {
	Host     string
	Port     int64
	User     string
	Password string
	Name     string
	SSLMode  string
}

func LoadDatabaseConfig() (*DatabaseConfig, error) {
	host, exists := os.LookupEnv("DATABASE_HOST")
	if !exists {
		return nil, errors.New("Biến môi trường DATABASE_HOST chưa được thiết lập")
	}
	portString, exists := os.LookupEnv("DATABASE_PORT")
	if !exists {
		return nil, errors.New("Biến môi trường DATABASE_PORT chưa được thiết lập")
	}
	port, error := strconv.ParseInt(portString, 10, 64)
	if error != nil {
		return nil, errors.New("Biến môi trường DATABASE_PORT không hợp lệ")
	}
	user, exists := os.LookupEnv("DATABASE_USER")
	if !exists {
		return nil, errors.New("Biến môi trường DATABASE_USER chưa được thiết lập")
	}
	password, exists := os.LookupEnv("DATABASE_PASSWORD")
	if !exists {
		return nil, errors.New("Biến môi trường DATABASE_PASSWORD chưa được thiết lập")
	}
	name, exists := os.LookupEnv("DATABASE_NAME")
	if !exists {
		return nil, errors.New("Biến môi trường DATABASE_NAME chưa được thiết lập")
	}
	sslMode, exists := os.LookupEnv("DATABASE_SSL_MODE")
	if !exists {
		return nil, errors.New("Biến môi trường DATABASE_SSL_MODE chưa được thiết lập")
	}
	return &DatabaseConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Name:     name,
		SSLMode:  sslMode,
	}, nil
}
