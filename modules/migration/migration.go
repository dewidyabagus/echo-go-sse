package migration

import (
	"gorm.io/gorm"

	"go-sse/modules/message"
)

func MigrationTables(db *gorm.DB) {
	db.AutoMigrate(&message.Message{})
}
