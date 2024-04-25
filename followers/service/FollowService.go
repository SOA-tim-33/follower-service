package service

import "database-example/repo"

type FollowService struct {
	FollowRepository *repo.FollowRepository
}

func (service *FollowService) Follow(id1, id2 int) error {
	err := service.FollowRepository.Following(id1, id2)
	return err
}

func (service *FollowService) CheckFollowing(id1, id2 int) (bool, error) {
	return service.FollowRepository.CheckFollowing(id1, id2)
}

func (service *FollowService) GetRecommendations(id int) ([]int, error) {
	recommendations, err := service.FollowRepository.GetRecommendation(id)
	if len(recommendations) == 0 {
		return []int{}, err
	}
	return recommendations, err
}
