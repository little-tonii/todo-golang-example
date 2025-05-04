package config

import (
	"errors"
	"os"
)

type environment struct {
	DATABASE_HOST      string
	DATABASE_PORT      string
	DATABASE_USER      string
	DATABASE_PASSWORD  string
	DATABASE_NAME      string
	DATABASE_SSL_MODE  string
	DATABASE_TIME_ZONE string
}

var Environment *environment

func LoadEnvironment() []error {
	errorList := make([]error, 0)
	databaseHost, exists := os.LookupEnv("DATABASE_HOST")
	if !exists {
		errorList = append(errorList, errors.New("Biến môi trường DATABASE_HOST chưa được thiết lập"))
	}
	databasePort, exists := os.LookupEnv("DATABASE_PORT")
	if !exists {
		errorList = append(errorList, errors.New("Biến môi trường DATABASE_PORT chưa được thiết lập"))
	}
	databaseUser, exists := os.LookupEnv("DATABASE_USER")
	if !exists {
		errorList = append(errorList, errors.New("Biến môi trường DATABASE_USER chưa được thiết lập"))
	}
	databasePassword, exists := os.LookupEnv("DATABASE_PASSWORD")
	if !exists {
		errorList = append(errorList, errors.New("Biến môi trường DATABASE_PASSWORD chưa được thiết lập"))
	}
	databaseName, exists := os.LookupEnv("DATABASE_NAME")
	if !exists {
		errorList = append(errorList, errors.New("Biến môi trường DATABASE_NAME chưa được thiết lập"))
	}
	databaseSSLMode, exists := os.LookupEnv("DATABASE_SSL_MODE")
	if !exists {
		errorList = append(errorList, errors.New("Biến môi trường DATABASE_SSL_MODE chưa được thiết lập"))
	}
	databaseTimeZone, exists := os.LookupEnv("DATABASE_TIME_ZONE")
	if !exists {
		errorList = append(errorList, errors.New("Biến môi trường DATABASE_TIME_ZONE chưa được thiết lập"))
	}
	Environment = &environment{
		DATABASE_HOST:      databaseHost,
		DATABASE_PORT:      databasePort,
		DATABASE_USER:      databaseUser,
		DATABASE_PASSWORD:  databasePassword,
		DATABASE_NAME:      databaseName,
		DATABASE_SSL_MODE:  databaseSSLMode,
		DATABASE_TIME_ZONE: databaseTimeZone,
	}
	return errorList
}
