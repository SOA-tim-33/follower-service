package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type UserBlogTourReportHandler struct {
	UserBlogTourReportService *service.UserBlogTourReportService
}

func (handler *UserBlogTourReportHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var userBlogTourReport model.UserBlogTourReport
	err := json.NewDecoder(req.Body).Decode(&userBlogTourReport)
	if err != nil {
		log.Println("Error while parsing json:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.UserBlogTourReportService.CreateUserBlogTourReport(&userBlogTourReport)
	if err != nil {
		log.Println("Error while creating a new user blog tour report:", err)
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *UserBlogTourReportHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("UserBlogTourReport with id: %s", id)
	// userBlogTourReport, err := handler.UserBlogTourReportService.FindUserBlogTourReport(id)
	// writer.Header().Set("Content-Type", "application/json")
	// if err != nil {
	// 	writer.WriteHeader(http.StatusNotFound)
	// 	return
	// }
	writer.WriteHeader(http.StatusOK)
	// json.NewEncoder(writer).Encode(userBlogTourReport)
}

func (handler *UserBlogTourReportHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	userBlogTourReports, err := handler.UserBlogTourReportService.FindAllUserBlogTourReports()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.Println("Error while fetching all user blog tour reports:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(userBlogTourReports)
}

func (handler *UserBlogTourReportHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var userBlogTourReport model.UserBlogTourReport
	err := json.NewDecoder(req.Body).Decode(&userBlogTourReport)
	if err != nil {
		log.Println("Error while parsing json:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.UserBlogTourReportService.UpdateUserBlogTourReport(&userBlogTourReport)
	if err != nil {
		log.Println("Error while updating user blog tour report:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (handler *UserBlogTourReportHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Deleting UserBlogTourReport with id: %s", id)
	err := handler.UserBlogTourReportService.DeleteUserBlogTourReport(id)
	if err != nil {
		log.Println("Error while deleting user blog tour report:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
