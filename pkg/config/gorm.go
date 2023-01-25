package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormConnection() (*gorm.DB, error) {
	dns := GetDatabaseDNS()
  	return gorm.Open(mysql.Open(dns), &gorm.Config{})
}

func GetDatabaseDNS() string {
	config := LoadEnv(".")
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
}

