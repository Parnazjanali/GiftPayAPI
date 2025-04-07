package sqliteDB

import (
	"apiGateway/internal/model"
	"fmt"

	"gorm.io/gorm"
)

func AuthenticateUser(db *gorm.DB, username, password string) (*model.User, error) {

	var user model.User
	result := db.Where("username = ? AND password = ?", username, password).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("invalid credentials")
		}
		return nil, result.Error
	}
	return &user, nil
}
