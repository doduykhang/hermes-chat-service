package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormConnection(config *Config) (*gorm.DB, error) {
	dns := GetDatabaseDNS(config.DB)
  	return gorm.Open(mysql.Open(dns), &gorm.Config{})
}

func GetDatabaseDNS(config DB) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)
}

