package config

import (
	"fmt"

	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type dbInfo struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func InitDatabase() {
	db := dbInfo{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "mysecretpassword",
		Name:     "postgres",
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", db.Host, db.Port, db.User, db.Password, db.Name)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = database

	migrateEntities()
}

func migrateEntities() {
	DB.AutoMigrate(&entities.Customer{})
}
