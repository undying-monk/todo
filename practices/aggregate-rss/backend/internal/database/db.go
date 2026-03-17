package database

import (
	"fmt"
	"log"
	"github.com/dungtc/aggregate-rss/backend/internal/config"
	"github.com/dungtc/aggregate-rss/backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		config.AppConfig.Database.Host,
		config.AppConfig.Database.User,
		config.AppConfig.Database.Password,
		config.AppConfig.Database.Name,
		config.AppConfig.Database.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Enable pg_trgm for GIN index on text search
	db.Exec("CREATE EXTENSION IF NOT EXISTS pg_trgm;")

	err = db.AutoMigrate(&models.Article{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = db
	log.Println("Database connection established and migrations run.")
}
