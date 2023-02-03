package config

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(config *Config) {
	dns := GetDatabaseDNS(config.DB)
    	m, err := migrate.New(
        	"file://migration/mysql/",
        	fmt.Sprintf("mysql://%s", dns),
	)
	if err != nil {
		log.Fatalf("Error migrating up: %s\n", err)
	}
    	err = m.Up()

	if err != nil {
		if err != migrate.ErrNoChange {
			log.Fatalf("Error migrating up: %s\n", err)
		}
	}
}
