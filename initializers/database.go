package initializers

import (
	"gofinance/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=gofinance password=gofinance dbname=gofinance port=5432 sslmode=disable TimeZone=Europe/Berlin"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(models.Category{})

	DB = database
}
