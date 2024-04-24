package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type UserService struct {
	UserRepo repo.IUserRepository
}

func (userService *UserService) Init(userRepository repo.IUserRepository) {
	userService.UserRepo = userRepository
}

func (service *UserService) Create(user *model.User) (*model.User, error) {
	createdUser, err := service.UserRepo.Create(user)
	if err != nil {
		return nil, fmt.Errorf("error creating user")
	}
	return &createdUser, nil
}

func (service *UserService) GetAll() ([]model.User, error) {
	users, err := service.UserRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error getting all users")
	}
	return users, nil
}

func (service *UserService) Get(id int) (*model.User, error) {
	user, err := service.UserRepo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("error getting user")
	}
	return &user, nil
}

func (service *UserService) Update(user *model.User) error {
	err := service.UserRepo.Update(user)
	if err != nil {
		return fmt.Errorf("error updating user")
	}
	return nil
}

func (service *UserService) Delete(id int) error {
	err := service.UserRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting user")
	}
	return nil
}
