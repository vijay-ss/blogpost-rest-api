package initializers

import (
	"blogpostApi/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}