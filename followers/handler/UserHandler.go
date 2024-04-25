package handler

import (
	"database-example/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	FollowerService *service.FollowService
}

func (handler *UserHandler) Follow(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id1, err1 := strconv.Atoi(vars["id1"])
	id2, err2 := strconv.Atoi(vars["id2"])
	if err1 != nil || err2 != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	var err = handler.FollowerService.Follow(id1, id2)
	if err == nil {
		writer.WriteHeader(http.StatusOK)
		return
	}
	fmt.Println(err)
	writer.WriteHeader(http.StatusInternalServerError)

}

func (handler *UserHandler) CheckIfFollowing(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id1, err1 := strconv.Atoi(vars["id1"])
	id2, err2 := strconv.Atoi(vars["id2"])
	if err1 != nil || err2 != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := handler.FollowerService.CheckFollowing(id1, id2)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := []byte(strconv.FormatBool(result))
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func (handler *UserHandler) GetRecommendation(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	result, error := handler.FollowerService.GetRecommendations(id)

	if error != nil {
		fmt.Println(error)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}
