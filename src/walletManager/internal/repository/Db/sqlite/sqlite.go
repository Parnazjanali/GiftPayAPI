package sqliteDb

import (
	"log"
	"walletManager/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() {
	var err error
	DB, err = gorm.Open(sqlite.Open("wallet.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	DB.AutoMigrate(&models.Wallet{})

}
