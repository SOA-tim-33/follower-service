package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type RatingService struct {
	RatingRepo repo.IRatingRepository
}

func (ratingService *RatingService) Init(ratingRepo repo.IRatingRepository) {
	ratingService.RatingRepo = ratingRepo
}

func (service *RatingService) Create(rating *model.Rating) (*model.Rating, error) {
	createdRating, err := service.RatingRepo.Create(rating)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return &createdRating, nil
}

func (service *RatingService) GetPositiveByBlog(blogId int) ([]model.Rating, error) {
	rating, err := service.RatingRepo.GetPositiveByBlog(blogId)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return rating, nil
}

func (service *RatingService) GetNegativeByBlog(blogId int) ([]model.Rating, error) {
	rating, err := service.RatingRepo.GetNegativeByBlog(blogId)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return rating, nil
}

// func (service *RatingService) CreateRating(rating *model.Rating) error {
// 	err := service.RatingRepo.CreateRating(rating)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (service *RatingService) FindRating(id string) (*model.Rating, error) {
// 	rating, err := service.RatingRepo.FindRatingById(id)
// 	if err != nil {
// 		return nil, fmt.Errorf(fmt.Sprintf("rating with id %s not found", id))
// 	}
// 	return &rating, nil
// }

// func (service *RatingService) FindAllRatings() ([]model.Rating, error) {
// 	ratings, err := service.RatingRepo.FindAllRatings()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return ratings, nil
// }

// func (service *RatingService) UpdateRating(rating *model.Rating) error {
// 	err := service.RatingRepo.UpdateRating(rating)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (service *RatingService) DeleteRating(id string) error {
// 	rating, err := service.RatingRepo.FindRatingById(id)
// 	if err != nil {
// 		return fmt.Errorf(fmt.Sprintf("rating with id %s not found", id))
// 	}
// 	err = service.RatingRepo.DeleteRating(&rating)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
