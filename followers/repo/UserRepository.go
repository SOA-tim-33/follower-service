package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	databaseConnection *gorm.DB
}

func (userRepository *UserRepository) Init(databaseConnection *gorm.DB) {
	userRepository.databaseConnection = databaseConnection
}

func (userRepository *UserRepository) Create(user *model.User) (model.User, error) {
	dbResult := userRepository.databaseConnection.Create(user)
	if dbResult != nil {
		return *user, dbResult.Error
	}
	return *user, nil
}

func (userRepository *UserRepository) GetAll() ([]model.User, error) {
	var users = []model.User{}
	dbResult := userRepository.databaseConnection.Find(&users)
	if dbResult != nil {
		return users, dbResult.Error
	}
	return users, nil
}

func (userRepository *UserRepository) Get(id int) (model.User, error) {
	var user = model.User{}
	dbResult := userRepository.databaseConnection.Find(&user, "\"Id\"=?", id)
	if dbResult != nil {
		return user, dbResult.Error
	}
	return user, nil
}

func (userRepository *UserRepository) Delete(id int) error {
	var user = model.User{}
	dbResult := userRepository.databaseConnection.Delete(&user, id)
	if dbResult != nil {
		return dbResult.Error
	}
	return nil
}

func (userRepository *UserRepository) Update(user *model.User) error {
	dbResult := userRepository.databaseConnection.Save(user)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}
