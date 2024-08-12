package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Config struct {
    DB       *gorm.DB
    JWTSecret string
}

func NewConfig() *Config {
    db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/medicalapp?charset=utf8mb4&parseTime=True&loc=Local")
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    return &Config{
        DB:       db,
        JWTSecret: os.Getenv("JWT_SECRET"),
    }
}
