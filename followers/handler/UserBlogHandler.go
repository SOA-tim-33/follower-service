package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type UserBlogHandler struct {
	UserBlogService    service.IUserBlogService
	BlogCommentService service.IBlogCommentService
	RatingService      service.IRatingService
}

func (handler *UserBlogHandler) InitRouter(blogService service.IUserBlogService, commentService service.IBlogCommentService,
	ratingService service.IRatingService) *chi.Mux {
	handler.UserBlogService = blogService
	handler.BlogCommentService = commentService
	handler.RatingService = ratingService

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/all", handler.GetAll)
	router.Post("/", handler.Create)
	router.Get("/{id}", handler.Get)
	router.Delete("/{id}", handler.Delete)
	router.Get("/byUser/{userId}", handler.GetByUser)
	router.Put("/{id}", handler.Update)

	//comment
	router.Get("/allComments", handler.GetAllComments)
	router.Post("/createComment", handler.CreateComment)
	router.Get("/byBlog/{blogId}", handler.GetByBlog)
	router.Delete("/deleteComment/{id}", handler.DeleteComment)
	router.Put("/updateComment/{id}", handler.UpdateComment)

	//rating
	router.Post("/createRating", handler.CreateRating)
	router.Get("/positiveByBlog/{blogId}", handler.GetPositiveByBlog)
	router.Get("/negativeByBlog/{blogId}", handler.GetNegativeByBlog)

	return router
}

func (handler *UserBlogHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var userBlog model.Blog
	err := json.NewDecoder(req.Body).Decode(&userBlog)
	if err != nil {
		log.Println("Error while parsing json:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	createdBlog, err := handler.UserBlogService.Create(&userBlog)
	if err != nil {
		log.Println("Error while creating a new user blog:", err)
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(createdBlog)
}

func (handler *UserBlogHandler) GetByBlog(writer http.ResponseWriter, reader *http.Request) {
	var blogId, _ = strconv.Atoi(chi.URLParam(reader, "blogId"))
	comments, err := handler.BlogCommentService.GetByBlog(blogId)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(comments)
}

func (handler *UserBlogHandler) GetPositiveByBlog(writer http.ResponseWriter, reader *http.Request) {
	var blogId, _ = strconv.Atoi(chi.URLParam(reader, "blogId"))
	ratings, err := handler.RatingService.GetPositiveByBlog(blogId)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(ratings)
}
func (handler *UserBlogHandler) GetNegativeByBlog(writer http.ResponseWriter, reader *http.Request) {
	var blogId, _ = strconv.Atoi(chi.URLParam(reader, "blogId"))
	ratings, err := handler.RatingService.GetNegativeByBlog(blogId)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(ratings)
}

func (handler *UserBlogHandler) CreateComment(writer http.ResponseWriter, req *http.Request) {
	var userBlog model.BlogComment
	err := json.NewDecoder(req.Body).Decode(&userBlog)
	if err != nil {
		log.Println("Error while parsing json:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	createdBlog, err := handler.BlogCommentService.Create(&userBlog)
	if err != nil {
		log.Println("Error while creating a new user blog:", err)
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(createdBlog)
}

func (handler *UserBlogHandler) CreateRating(writer http.ResponseWriter, req *http.Request) {
	var rating model.Rating
	err := json.NewDecoder(req.Body).Decode(&rating)
	if err != nil {
		log.Println("Error while parsing json:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	createdComment, err := handler.RatingService.Create(&rating)

	if err != nil {
		log.Println("Error while creating a new user blog:", err)
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(createdComment)
}

func (handler *UserBlogHandler) GetAll(writer http.ResponseWriter, reader *http.Request) {
	blogs, err := handler.UserBlogService.GetAll()
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(blogs)
}

func (handler *UserBlogHandler) GetAllComments(writer http.ResponseWriter, reader *http.Request) {
	blogs, err := handler.BlogCommentService.GetAll()
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(blogs)
}

//	func (handler *UserBlogHandler) Get(writer http.ResponseWriter, req *http.Request) {
//		id := mux.Vars(req)["id"]
//		log.Printf("UserBlog with id: %s", id)
//		// userBlog, err := handler.UserBlogService.FindUserBlog(id)
//		// writer.Header().Set("Content-Type", "application/json")
//		// if err != nil {
//		// 	writer.WriteHeader(http.StatusNotFound)
//		// 	return
//		// }
//		writer.WriteHeader(http.StatusOK)
//		// json.NewEncoder(writer).Encode(userBlog)
//	}

func (handler *UserBlogHandler) Get(writer http.ResponseWriter, reader *http.Request) {
	log.Println("ulazi u GET")
	var id, _ = strconv.Atoi(chi.URLParam(reader, "id"))
	blog, err := handler.UserBlogService.Get(id)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	if blog.Id != 0 { // find better way, lazy
		json.NewEncoder(writer).Encode(blog)
	}
}

func (handler *UserBlogHandler) GetByUser(writer http.ResponseWriter, reader *http.Request) {
	var userId, _ = strconv.Atoi(chi.URLParam(reader, "userId"))
	blogs, err := handler.UserBlogService.GetByUser(userId)
	log.Println(blogs)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(blogs)
}

func (handler *UserBlogHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	log.Println("ulazi u blog delete")
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		log.Println("Error while parsing ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.UserBlogService.Delete(id)
	if err != nil {
		log.Println("Error while deleting user blog:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]string{"message": "Blog deleted successfully"})
}

func (handler *UserBlogHandler) Update(writer http.ResponseWriter, req *http.Request) {
	log.Println("ulazi u blog update")
	var userBlog model.Blog
	err := json.NewDecoder(req.Body).Decode(&userBlog)
	if err != nil {
		log.Println("Error while parsing json:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extracting ID from URL
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		log.Println("Error while parsing ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	// Setting the ID of the user blog from the URL
	userBlog.Id = id

	err = handler.UserBlogService.Update(&userBlog)
	if err != nil {
		log.Println("Error while updating user blog:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(userBlog)
}

func (handler *UserBlogHandler) DeleteComment(writer http.ResponseWriter, req *http.Request) {
	log.Println("ulazi u comment delete")
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		log.Println("Error while parsing ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.BlogCommentService.Delete(id)
	if err != nil {
		log.Println("Error while deleting blog comment:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]string{"message": "Comment deleted successfully"})
}

func (handler *UserBlogHandler) UpdateComment(writer http.ResponseWriter, req *http.Request) {
	log.Println("ulazi u comment update")
	var blogComment model.BlogComment
	err := json.NewDecoder(req.Body).Decode(&blogComment)
	if err != nil {
		log.Println("Error while parsing json:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extracting ID from URL
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		log.Println("Error while parsing ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	// Setting the ID of the user blog from the URL
	blogComment.Id = id

	err = handler.BlogCommentService.Update(&blogComment)
	if err != nil {
		log.Println("Error while updating blog comment:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(blogComment)
}

// func (handler *UserBlogHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
// 	userBlogs, err := handler.UserBlogService.FindAllUserBlogs()
// 	writer.Header().Set("Content-Type", "application/json")
// 	if err != nil {
// 		log.Println("Error while fetching all user blogs:", err)
// 		writer.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	writer.WriteHeader(http.StatusOK)
// 	json.NewEncoder(writer).Encode(userBlogs)
// }

// func (handler *UserBlogHandler) Update(writer http.ResponseWriter, req *http.Request) {
// 	var userBlog model.UserBlog
// 	err := json.NewDecoder(req.Body).Decode(&userBlog)
// 	if err != nil {
// 		log.Println("Error while parsing json:", err)
// 		writer.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	err = handler.UserBlogService.UpdateUserBlog(&userBlog)
// 	if err != nil {
// 		log.Println("Error while updating user blog:", err)
// 		writer.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	writer.WriteHeader(http.StatusOK)
// }

// func (handler *UserBlogHandler) Delete(writer http.ResponseWriter, req *http.Request) {
// 	id := mux.Vars(req)["id"]
// 	log.Printf("Deleting UserBlog with id: %s", id)
// 	err := handler.UserBlogService.DeleteUserBlog(id)
// 	if err != nil {
// 		log.Println("Error while deleting user blog:", err)
// 		writer.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	writer.WriteHeader(http.StatusOK)
// }
