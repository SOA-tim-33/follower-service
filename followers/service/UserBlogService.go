package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type UserBlogService struct {
	UserBlogRepo repo.IUserBlogRepository
}

func (blogService *UserBlogService) Init(blogRepository repo.IUserBlogRepository) {
	blogService.UserBlogRepo = blogRepository
}

func (service *UserBlogService) Create(userBlog *model.Blog) (*model.Blog, error) {
	createdBlog, err := service.UserBlogRepo.Create(userBlog)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return &createdBlog, nil
}

func (service *UserBlogService) GetAll() ([]model.Blog, error) {
	blogs, err := service.UserBlogRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return blogs, nil
}

func (service *UserBlogService) Get(id int) (*model.Blog, error) {
	blog, err := service.UserBlogRepo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return &blog, nil
}

func (service *UserBlogService) GetByUser(userId int) ([]model.Blog, error) {
	blog, err := service.UserBlogRepo.GetByUser(userId)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return blog, nil
}

func (service *UserBlogService) Delete(id int) error {
	err := service.UserBlogRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("error")
	}
	return nil
}

func (service *UserBlogService) Update(userBlog *model.Blog) error {
	err := service.UserBlogRepo.Update(userBlog)
	if err != nil {
		return fmt.Errorf("error")
	}
	return nil
}

// func (service *UserBlogService) FindUserBlog(id string) (*model.UserBlog, error) {
// 	userBlog, err := service.UserBlogRepo.FindUserBlogById(id)
// 	if err != nil {
// 		return nil, fmt.Errorf(fmt.Sprintf("user blog with id %s not found", id))
// 	}
// 	return &userBlog, nil
// }

// func (service *UserBlogService) FindAllUserBlogs() ([]model.UserBlog, error) {
// 	userBlogs, err := service.UserBlogRepo.FindAllUserBlogs()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return userBlogs, nil
// }

// func (service *UserBlogService) UpdateUserBlog(userBlog *model.UserBlog) error {
// 	err := service.UserBlogRepo.UpdateUserBlog(userBlog)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (service *UserBlogService) DeleteUserBlog(id string) error {
// 	userBlog, err := service.UserBlogRepo.FindUserBlogById(id)
// 	if err != nil {
// 		return fmt.Errorf(fmt.Sprintf("user blog with id %s not found", id))
// 	}
// 	err = service.UserBlogRepo.DeleteUserBlog(&userBlog)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
