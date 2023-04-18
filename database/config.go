package database

import (
    "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UrlStore struct {
    gorm.Model
    OriginalUrl string
}

func InitializeGorm(dsn string) (*gorm.DB, error) {
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    return db, err
}
