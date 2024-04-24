package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type ProfileService struct {
	ProfileRepo repo.IProfileRepository
}

func (profileService *ProfileService) Init(profileRepository repo.IProfileRepository) {
	profileService.ProfileRepo = profileRepository
}

func (service *ProfileService) Create(profile *model.Profile) (*model.Profile, error) {
	createdProfile, err := service.ProfileRepo.Create(profile)
	if err != nil {
		return nil, fmt.Errorf("error creating profile")
	}
	return &createdProfile, nil
}

func (service *ProfileService) GetAll() ([]model.Profile, error) {
	profiles, err := service.ProfileRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error getting all profiles")
	}
	return profiles, nil
}

func (service *ProfileService) Get(id int) (model.Profile, error) {
	profile, err := service.ProfileRepo.Get(id)
	if err != nil {
		return model.Profile{}, fmt.Errorf("error getting profile: %v", err)
	}
	return profile, nil
}

func (service *ProfileService) Update(profile *model.Profile) error {
	err := service.ProfileRepo.Update(profile)
	if err != nil {
		return fmt.Errorf("error updating profile")
	}
	return nil
}

func (service *ProfileService) Delete(id int) error {
	err := service.ProfileRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting profile")
	}
	return nil
}
