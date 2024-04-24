package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type ICRUDRepository interface {
	Create(user *model.User) (model.User, error)
	Get(id int) (model.User, error)
	GetAll() ([]model.User, error)
	Delete(id int) error
	Update(user *model.User) error
}

type IUserRepository interface {
	ICRUDRepository
	Init(databaseConnection *gorm.DB)
}

type IProfileRepository interface {
	Create(profile *model.Profile) (model.Profile, error)
	Get(id int) (model.Profile, error)
	GetAll() ([]model.Profile, error)
	Init(databaseConnection *gorm.DB)
	Delete(id int) error
	Update(blogComment *model.Profile) error
}

type IFollowRepository interface {
	Create(profile *model.Follow) (model.Follow, error)
	Get(id int) (model.Follow, error)
	GetAll() ([]model.Follow, error)
	Init(databaseConnection *gorm.DB)
	Delete(id int) error
	Update(blogComment *model.Follow) error
}

type ITourPreferenceRepository interface {
	Create(profile *model.TourPreference) (model.TourPreference, error)
	Get(id int) (model.TourPreference, error)
	GetAll() ([]model.TourPreference, error)
	Init(databaseConnection *gorm.DB)
	Delete(id int) error
	Update(blogComment *model.TourPreference) error
}
