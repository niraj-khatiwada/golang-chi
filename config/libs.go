package config

import "gorm.io/gorm"

type Libs struct {
	DB *gorm.DB
}
