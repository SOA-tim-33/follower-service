package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type FollowService struct {
	FollowRepo repo.IFollowRepository
}

func (followService *FollowService) Init(followRepository repo.IFollowRepository) {
	followService.FollowRepo = followRepository
}

func (service *FollowService) Create(follow *model.Follow) (*model.Follow, error) {
	createdFollow, err := service.FollowRepo.Create(follow)
	if err != nil {
		return nil, fmt.Errorf("error creating follow")
	}
	return &createdFollow, nil
}

func (service *FollowService) GetAll() ([]model.Follow, error) {
	follows, err := service.FollowRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error getting all follows")
	}
	return follows, nil
}

func (service *FollowService) Get(id int) (*model.Follow, error) {
	follow, err := service.FollowRepo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("error getting follow")
	}
	return &follow, nil
}

func (service *FollowService) Update(follow *model.Follow) error {
	err := service.FollowRepo.Update(follow)
	if err != nil {
		return fmt.Errorf("error updating follow")
	}
	return nil
}

func (service *FollowService) Delete(id int) error {
	err := service.FollowRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting follow")
	}
	return nil
}
