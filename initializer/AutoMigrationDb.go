package initializer

import (
	"fmt"
	"go-jwt/models"
)

func AutoMigrationDb() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Note{},
		&models.Tag{},
	)

	if err != nil {
		fmt.Println("can't running migration")
	}
}
