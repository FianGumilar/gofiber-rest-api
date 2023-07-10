package migrations

import (
	"fmt"
	"log"

	"github.com/FianGumilar/gofiber-rest-api/database"
	"github.com/FianGumilar/gofiber-rest-api/models"
)

func Migration() {
	err := database.DB.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatal("Failed to migrate...")
	}

	fmt.Println("Migrated successfully")
}
