package service

import (
	"database-example/model"
	"database-example/repo"
)

type ICRUDService interface {
	Create(userBlog *model.Blog) (*model.Blog, error)
	GetAll() ([]model.Blog, error)
	Get(id int) (*model.Blog, error)
	Delete(id int) error
	GetByUser(userId int) ([]model.Blog, error)
	Update(userBlog *model.Blog) error
}

type IUserBlogService interface {
	ICRUDService
	Init(crudRepository repo.IUserBlogRepository)
}
type IBlogCommentService interface {
	Init(crudRepository repo.IBlogCommentRepository)
	GetAll() ([]model.BlogComment, error)
	Create(userBlog *model.BlogComment) (*model.BlogComment, error)
	GetByBlog(blogId int) ([]model.BlogComment, error)
	Delete(id int) error
	Update(blogComment *model.BlogComment) error
}
type IRatingService interface {
	Init(crudRepository repo.IRatingRepository)
	Create(rating *model.Rating) (*model.Rating, error)
	GetPositiveByBlog(blogId int) ([]model.Rating, error)
	GetNegativeByBlog(blogId int) ([]model.Rating, error)
}
