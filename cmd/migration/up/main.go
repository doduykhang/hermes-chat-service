package main

import (
	"doduykhang/hermes-chat/pkg/config"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	dns := config.GetDatabaseDNS()
    	m, err := migrate.New(
        	"file://migration/mysql/",
        	fmt.Sprintf("mysql://%s", dns),
	)
	if err != nil {
		log.Fatalf("Error migrating up: %s\n", err)
	}
    	err = m.Up()

	if err != nil {
		log.Fatalf("Error migrating up: %s\n", err)
	}
}
