package config

import (
	"fmt"
	"go-web/utils"
	"os"
	"time"
)

type Database struct {
	Host     string `env:"DATABASE_HOST"`
	Port     string `env:"DATABASE_PORT"`
	User     string `env:"DATABASE_USER"`
	Name     string `env:"DATABASE_NAME"`
	Password string `env:"DATABASE_PASSWORD"`
}

const (
	DbMaxOpenConnections    = 25
	DbMaxIdleConnections    = 100
	DbConnectionMaxLifetime = time.Hour
)

func GetDefaultDatabaseConfig() Database {
	utils.LoadEnv()
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USER")
	name := os.Getenv("DATABASE_NAME")
	password := os.Getenv("DATABASE_PASSWORD")
	return Database{Host: host, Port: port, User: user, Name: name, Password: password}
}

func CreateDatabaseDSN(config *Database) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.Name)
}
