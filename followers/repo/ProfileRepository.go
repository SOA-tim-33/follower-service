package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type ProfileRepository struct {
	databaseConnection *gorm.DB
}

func (profileRepository *ProfileRepository) Init(databaseConnection *gorm.DB) {
	profileRepository.databaseConnection = databaseConnection
}

func (profileRepository *ProfileRepository) Create(profile *model.Profile) (model.Profile, error) {
	dbResult := profileRepository.databaseConnection.Create(profile)
	if dbResult != nil {
		return *profile, dbResult.Error
	}
	return *profile, nil
}

func (profileRepository *ProfileRepository) GetAll() ([]model.Profile, error) {
	var profiles = []model.Profile{}
	dbResult := profileRepository.databaseConnection.Find(&profiles)
	if dbResult != nil {
		return profiles, dbResult.Error
	}
	return profiles, nil
}

func (profileRepository *ProfileRepository) Get(id int) (model.Profile, error) {
	var profile = model.Profile{}
	dbResult := profileRepository.databaseConnection.Find(&profile, "\"Id\"=?", id)
	if dbResult != nil {
		return profile, dbResult.Error
	}
	return profile, nil
}

func (profileRepository *ProfileRepository) Delete(id int) error {
	var profile = model.Profile{}
	dbResult := profileRepository.databaseConnection.Delete(&profile, id)
	if dbResult != nil {
		return dbResult.Error
	}
	return nil
}

func (profileRepository *ProfileRepository) Update(profile *model.Profile) error {
	dbResult := profileRepository.databaseConnection.Save(profile)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}
