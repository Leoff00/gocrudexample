package migrations

import (
	"github.com/Leoff00/go-crud-api/models"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(models.User{})
}
