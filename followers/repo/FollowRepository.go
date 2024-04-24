package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type FollowRepository struct {
	databaseConnection *gorm.DB
}

func (followRepository *FollowRepository) Init(databaseConnection *gorm.DB) {
	followRepository.databaseConnection = databaseConnection
}

func (followRepository *FollowRepository) Create(follow *model.Follow) (model.Follow, error) {
	dbResult := followRepository.databaseConnection.Create(follow)
	if dbResult != nil {
		return *follow, dbResult.Error
	}
	return *follow, nil
}

func (followRepository *FollowRepository) GetAll() ([]model.Follow, error) {
	var follows = []model.Follow{}
	dbResult := followRepository.databaseConnection.Find(&follows)
	if dbResult != nil {
		return follows, dbResult.Error
	}
	return follows, nil
}

func (followRepository *FollowRepository) Get(id int) (model.Follow, error) {
	var follow = model.Follow{}
	dbResult := followRepository.databaseConnection.Find(&follow, "\"Id\"=?", id)
	if dbResult != nil {
		return follow, dbResult.Error
	}
	return follow, nil
}

func (followRepository *FollowRepository) Delete(id int) error {
	var follow = model.Follow{}
	dbResult := followRepository.databaseConnection.Delete(&follow, id)
	if dbResult != nil {
		return dbResult.Error
	}
	return nil
}

func (followRepository *FollowRepository) Update(follow *model.Follow) error {
	dbResult := followRepository.databaseConnection.Save(follow)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}
