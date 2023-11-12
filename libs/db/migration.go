package db

import (
	"fmt"
	"go-web/config"
	"go-web/models"
	"log"
)

func Migrate() {
	log.Println("[info] Migration: Started")

	database := &DB{}
	if err := database.Init(config.Database{}); err != nil {
		log.Fatal("[error] Migration: Error during database initialization ", err)
	}

	var tables []interface{}
	for name, table := range models.GetAllModels() {
		fmt.Println(name, table)
		tables = append(tables, table)
	}

	fmt.Println(tables)

	if err := database.Client.AutoMigrate(tables...); err != nil {
		log.Fatal("[error] Migration: Error during migration")
	}
	log.Println("[info] Migration: Completed")
}
