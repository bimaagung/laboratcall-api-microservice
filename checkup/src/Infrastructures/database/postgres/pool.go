package database

import (
	"log"
	"os"

	postgres "go.elastic.co/apm/module/apmgormv2/v2/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB() *gorm.DB {
	dsn := os.Getenv("DB_URL_TEST")
	DB, err :=  gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatalln(err)
	}

	return DB
}
