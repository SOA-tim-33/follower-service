package service

import (
	"database-example/model"
	"database-example/repo"
)

type ICRUDService interface {
	Create(user *model.User) (*model.User, error)
	GetAll() ([]model.User, error)
	Get(id int) (*model.User, error)
	Delete(id int) error
	Update(user *model.User) error
}

type IUserService interface {
	ICRUDService
	Init(crudRepository repo.IUserRepository)
}

type IProfileService interface {
	Init(crudRepository repo.IProfileRepository)
	GetAll() ([]model.Profile, error)
	Get(id int) ([]model.Profile, error)
	Create(profile *model.Profile) (*model.Profile, error)
	Delete(id int) error
	Update(profile *model.Profile) error
}

type IFollowService interface {
	Init(crudRepository repo.IFollowRepository)
	GetAll() ([]model.Follow, error)
	Get(id int) ([]model.Follow, error)
	Create(follow *model.Follow) (*model.Follow, error)
	Delete(id int) error
	Update(follow *model.Profile) error
}

type ITourPreferenceService interface {
	Init(crudRepository repo.ITourPreferenceRepository)
	GetAll() ([]model.TourPreference, error)
	Get(id int) ([]model.TourPreference, error)
	Create(tourPreference *model.TourPreference) (*model.TourPreference, error)
	Delete(id int) error
	Update(tourPreference *model.TourPreference) error
}
