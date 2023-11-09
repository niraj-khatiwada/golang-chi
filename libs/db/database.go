package db

import (
	"fmt"
	"go-web/config"
	"go-web/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

func InitializeDB(conf config.Database) (*gorm.DB, error) {
	var dbConf config.Database
	if conf != (config.Database{}) {
		dbConf = conf
	} else {
		dbConf = getDefaultConfig()
	}
	dsn := createDSN(&dbConf)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		utils.CatchRuntimeErrors(err)
		return nil, err
	}
	if sqlDB, err := db.DB(); err != nil {
		utils.CatchRuntimeErrors(err)
		return nil, err
	} else {
		sqlDB.SetMaxOpenConns(25)
		sqlDB.SetMaxIdleConns(25)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}
	return db, nil
}

func getDefaultConfig() config.Database {
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USER")
	name := os.Getenv("DATABASE_NAME")
	password := os.Getenv("DATABASE_PASSWORD")
	return config.Database{Host: host, Port: port, User: user, Name: name, Password: password}
}

func createDSN(config *config.Database) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.Name)
}
