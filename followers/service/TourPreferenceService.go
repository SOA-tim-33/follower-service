package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type TourPreferenceService struct {
	TourPreferenceRepo repo.ITourPreferenceRepository
}

func (tourPreferenceService *TourPreferenceService) Init(tourPreferenceRepository repo.ITourPreferenceRepository) {
	tourPreferenceService.TourPreferenceRepo = tourPreferenceRepository
}

func (service *TourPreferenceService) Create(tourPreference *model.TourPreference) (*model.TourPreference, error) {
	createdTourPreference, err := service.TourPreferenceRepo.Create(tourPreference)
	if err != nil {
		return nil, fmt.Errorf("error creating tour preference")
	}
	return &createdTourPreference, nil
}

func (service *TourPreferenceService) GetAll() ([]model.TourPreference, error) {
	tourPreferences, err := service.TourPreferenceRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error getting all tour preferences")
	}
	return tourPreferences, nil
}

func (service *TourPreferenceService) Get(id int) (model.TourPreference, error) {
	tourPreference, err := service.TourPreferenceRepo.Get(id)
	if err != nil {
		return model.TourPreference{}, fmt.Errorf("error getting tour preference")
	}
	return tourPreference, nil
}

func (service *TourPreferenceService) Update(tourPreference *model.TourPreference) error {
	err := service.TourPreferenceRepo.Update(tourPreference)
	if err != nil {
		return fmt.Errorf("error updating tour preference")
	}
	return nil
}

func (service *TourPreferenceService) Delete(id int) error {
	err := service.TourPreferenceRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting tour preference")
	}
	return nil
}
