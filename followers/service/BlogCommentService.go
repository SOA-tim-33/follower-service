package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type BlogCommentService struct {
	BlogCommentRepo repo.IBlogCommentRepository
}

func (blogCommentService *BlogCommentService) Init(blogRepository repo.IBlogCommentRepository) {
	blogCommentService.BlogCommentRepo = blogRepository
}

func (service *BlogCommentService) GetAll() ([]model.BlogComment, error) {
	blogs, err := service.BlogCommentRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return blogs, nil
}

func (service *BlogCommentService) Create(comment *model.BlogComment) (*model.BlogComment, error) {
	createdComment, err := service.BlogCommentRepo.Create(comment)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return &createdComment, nil
}

func (service *BlogCommentService) GetByBlog(blogId int) ([]model.BlogComment, error) {
	comment, err := service.BlogCommentRepo.GetByBlog(blogId)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return comment, nil
}

func (service *BlogCommentService) Delete(id int) error {
	err := service.BlogCommentRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("error")
	}
	return nil
}

func (service *BlogCommentService) Update(comment *model.BlogComment) error {
	err := service.BlogCommentRepo.Update(comment)
	if err != nil {
		return fmt.Errorf("error")
	}
	return nil
}

// func (service *BlogCommentService) CreateBlogComment(blogComment *model.BlogComment) error {
// 	err := service.BlogCommentRepo.CreateBlogComment(blogComment)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (service *BlogCommentService) FindBlogComment(id string) (*model.BlogComment, error) {
// 	blogComment, err := service.BlogCommentRepo.FindBlogCommentById(id)
// 	if err != nil {
// 		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
// 	}
// 	return &blogComment, nil
// }

// func (service *BlogCommentService) FindAllBlogComments() ([]model.BlogComment, error) {
// 	blogComments, err := service.BlogCommentRepo.FindAllBlogComments()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return blogComments, nil
// }

// func (service *BlogCommentService) UpdateBlogComment(blogComment *model.BlogComment) error {
// 	err := service.BlogCommentRepo.UpdateBlogComment(blogComment)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (service *BlogCommentService) DeleteBlogComment(id string) error {
// 	blogComment, err := service.BlogCommentRepo.FindBlogCommentById(id)
// 	if err != nil {
// 		return fmt.Errorf(fmt.Sprintf("blog comment with id %s not found", id))
// 	}
// 	err = service.BlogCommentRepo.DeleteBlogComment(&blogComment)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
