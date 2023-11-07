package models
import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dbURL := "postgres://user:pass@localhost:5432/db"

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        panic(err)
    }

    db.AutoMigrate(&Book{})

    DB = db
}