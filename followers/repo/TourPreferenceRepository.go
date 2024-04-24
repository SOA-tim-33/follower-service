package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type TourPreferenceRepository struct {
	databaseConnection *gorm.DB
}

func (tourPreferenceRepository *TourPreferenceRepository) Init(databaseConnection *gorm.DB) {
	tourPreferenceRepository.databaseConnection = databaseConnection
}

func (tourPreferenceRepository *TourPreferenceRepository) Create(tourPreference *model.TourPreference) (model.TourPreference, error) {
	dbResult := tourPreferenceRepository.databaseConnection.Create(tourPreference)
	if dbResult != nil {
		return *tourPreference, dbResult.Error
	}
	return *tourPreference, nil
}

func (tourPreferenceRepository *TourPreferenceRepository) GetAll() ([]model.TourPreference, error) {
	var tourPreferences = []model.TourPreference{}
	dbResult := tourPreferenceRepository.databaseConnection.Find(&tourPreferences)
	if dbResult != nil {
		return tourPreferences, dbResult.Error
	}
	return tourPreferences, nil
}

func (tourPreferenceRepository *TourPreferenceRepository) Get(id int) (model.TourPreference, error) {
	var tourPreference = model.TourPreference{}
	dbResult := tourPreferenceRepository.databaseConnection.Find(&tourPreference, "\"Id\"=?", id)
	if dbResult != nil {
		return tourPreference, dbResult.Error
	}
	return tourPreference, nil
}

func (tourPreferenceRepository *TourPreferenceRepository) Delete(id int) error {
	var tourPreference = model.TourPreference{}
	dbResult := tourPreferenceRepository.databaseConnection.Delete(&tourPreference, id)
	if dbResult != nil {
		return dbResult.Error
	}
	return nil
}

func (tourPreferenceRepository *TourPreferenceRepository) Update(tourPreference *model.TourPreference) error {
	dbResult := tourPreferenceRepository.databaseConnection.Save(tourPreference)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}
