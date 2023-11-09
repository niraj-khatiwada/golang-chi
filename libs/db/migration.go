package db

import (
	"fmt"
	"go-web/config"
	"go-web/models"
	"log"
)

func Migrate() {
	log.Println("[info] Migration: Started")
	db, err := InitDB(config.Database{})
	if err != nil {
		log.Fatal("[error] Migration: Error during database initialization ", err)
	}

	var tables []interface{}
	for name, table := range models.GetAllModels() {
		fmt.Println(name, table)
		tables = append(tables, table)
	}

	fmt.Println(tables)

	if err := db.AutoMigrate(tables...); err != nil {
		log.Fatal("[error] Migration: Error during migration")
	}
	log.Println("[info] Migration: Completed")
}
