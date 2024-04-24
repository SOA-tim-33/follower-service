package handler

import "database-example/service"

type BlogCommentHandler struct {
	BlogCommentService *service.BlogCommentService
}

// func (handler *BlogCommentHandler) Create(writer http.ResponseWriter, req *http.Request) {
// 	var blogComment model.BlogComment
// 	err := json.NewDecoder(req.Body).Decode(&blogComment)
// 	if err != nil {
// 		println("Error while parsing json")
// 		writer.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	err = handler.BlogCommentService.CreateBlogComment(&blogComment)
// 	if err != nil {
// 		println("Error while creating a new blog comment")
// 		writer.WriteHeader(http.StatusExpectationFailed)
// 		return
// 	}
// 	writer.WriteHeader(http.StatusCreated)
// 	writer.Header().Set("Content-Type", "application/json")
// }

// func (handler *BlogCommentHandler) Get(writer http.ResponseWriter, req *http.Request) {
// 	id := mux.Vars(req)["id"]
// 	log.Printf("BlogComment sa id-em %s", id)
// 	// blogComment, err := handler.BlogCommentService.FindBlogComment(id)
// 	// writer.Header().Set("Content-Type", "application/json")
// 	// if err != nil {
// 	// 	writer.WriteHeader(http.StatusNotFound)
// 	// 	return
// 	// }
// 	writer.WriteHeader(http.StatusOK)
// 	// json.NewEncoder(writer).Encode(blogComment)
// }

// func (handler *BlogCommentHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
// 	blogComments, err := handler.BlogCommentService.FindAllBlogComments()
// 	writer.Header().Set("Content-Type", "application/json")
// 	if err != nil {
// 		log.Println("Error while fetching all blog comments:", err)
// 		writer.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	writer.WriteHeader(http.StatusOK)
// 	json.NewEncoder(writer).Encode(blogComments)
// }

// func (handler *BlogCommentHandler) Update(writer http.ResponseWriter, req *http.Request) {
// 	var blogComment model.BlogComment
// 	err := json.NewDecoder(req.Body).Decode(&blogComment)
// 	if err != nil {
// 		log.Println("Error while parsing json:", err)
// 		writer.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	err = handler.BlogCommentService.UpdateBlogComment(&blogComment)
// 	if err != nil {
// 		log.Println("Error while updating blog comment:", err)
// 		writer.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	writer.WriteHeader(http.StatusOK)
// }

// func (handler *BlogCommentHandler) Delete(writer http.ResponseWriter, req *http.Request) {
// 	id := mux.Vars(req)["id"]
// 	log.Printf("Deleting BlogComment with id: %s", id)
// 	err := handler.BlogCommentService.DeleteBlogComment(id)
// 	if err != nil {
// 		log.Println("Error while deleting blog comment:", err)
// 		writer.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	writer.WriteHeader(http.StatusOK)
// }
