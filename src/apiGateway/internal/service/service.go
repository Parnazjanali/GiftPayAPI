package Service

import (
	"apiGateway/internal/model"
	sqliteDB "apiGateway/repository/db/sqlite"
)

func AuthenticateUser(username, password string) (*model.User, error) {

	user, err := sqliteDB.AuthenticateUser(sqliteDB.DB, username, password)
	if err != nil {
		return nil, err
	}
	return user, nil

}
