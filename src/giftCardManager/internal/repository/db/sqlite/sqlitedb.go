package sqlitedb

import (
	"GiftCardManager/internal/models"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() {
	var err error
	Db, err = gorm.Open(sqlite.Open("giftcards.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	Db.AutoMigrate(models.GIftCard{})

}
