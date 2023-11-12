package db

import (
	"go-web/config"
	"go-web/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DB struct {
	Client *gorm.DB
}

func (database *DB) Init(conf config.Database) error {
	var dbConf config.Database
	if conf != (config.Database{}) {
		dbConf = conf
	} else {
		dbConf = config.GetDefaultDatabaseConfig()
	}
	dsn := config.CreateDatabaseDSN(&dbConf)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		utils.CatchRuntimeErrors(err)
		return err
	}
	if sqlDB, err := db.DB(); err != nil {
		utils.CatchRuntimeErrors(err)
		return err
	} else {
		sqlDB.SetMaxOpenConns(config.DbMaxOpenConnections)
		sqlDB.SetMaxIdleConns(config.DbMaxIdleConnections)
		sqlDB.SetConnMaxLifetime(config.DbConnectionMaxLifetime)
	}
	database.Client = db
	return nil
}
