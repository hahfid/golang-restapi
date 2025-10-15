package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config interface {
	DSN() string
}

func NewMySQLConnection(cfg Config) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(cfg.DSN()), &gorm.Config{})
}
