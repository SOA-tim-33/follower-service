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

type UserHandler struct {
	UserService           service.IUserService
	ProfileService        service.IProfileService
	FollowService         service.IFollowService
	TourPreferenceService service.ITourPreferenceService
}

func (handler *UserHandler) InitRouter(userService service.IUserService, profileService service.IProfileService,
	followService service.IFollowService, tourPreferenceService service.ITourPreferenceService) *chi.Mux {
	handler.UserService = userService
	handler.ProfileService = profileService
	handler.FollowService = followService
	handler.TourPreferenceService = tourPreferenceService

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	//user
	router.Get("/{id}", handler.Get)
	router.Get("/all", handler.GetAll)
	router.Post("/", handler.Create)
	router.Delete("/{id}", handler.Delete)
	router.Put("/{id}", handler.Update)

	//profile
	router.Get("/{id}", handler.GetProfile)
	router.Get("/allProfiles", handler.GetAllProfiles)
	router.Post("/createProfile", handler.CreateProfile)
	router.Delete("/deleteProfile/{id}", handler.DeleteProfile)
	router.Put("/updateProfile/{id}", handler.UpdateProfile)

	//follow
	router.Get("/{id}", handler.GetFollow)
	router.Get("/allFollows", handler.GetAllFollows)
	router.Post("/createFollow", handler.CreateFollow)
	router.Delete("/deleteFollow/{id}", handler.DeleteFollow)
	router.Put("/updateFollow/{id}", handler.UpdateFollow)

	//tour preference
	router.Get("/{id}", handler.GetTourPreference)
	router.Get("/allProfiles", handler.GetAllTourPreferences)
	router.Post("/createProfile", handler.CreateTourPreference)
	router.Delete("/deleteProfile/{id}", handler.DeleteTourPreference)
	router.Put("/updateProfile/{id}", handler.UpdateTourPreference)

	return router
}

func (handler *UserHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		log.Println("Error while parsing json:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	createdUser, err := handler.UserService.Create(&user)
	if err != nil {
		log.Println("Error while creating a new user:", err)
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(createdUser)
}

func (handler *UserHandler) CreateProfile(writer http.ResponseWriter, req *http.Request) {
	var profile model.Profile
	err := json.NewDecoder(req.Body).Decode(&profile)
	if err != nil {
		log.Println("Error while parsing json:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	createdProfile, err := handler.ProfileService.Create(&profile)
	if err != nil {
		log.Println("Error while creating a new profile:", err)
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(createdProfile)
}

func (handler *UserHandler) CreateFollow(writer http.ResponseWriter, req *http.Request) {
	var follow model.Follow
	err := json.NewDecoder(req.Body).Decode(&follow)
	if err != nil {
		log.Println("Error while parsing json:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	createdFollow, err := handler.FollowService.Create(&follow)
	if err != nil {
		log.Println("Error while creating a new follow:", err)
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(createdFollow)
}

func (handler *UserHandler) CreateTourPreference(writer http.ResponseWriter, req *http.Request) {
	var tourPreference model.TourPreference
	err := json.NewDecoder(req.Body).Decode(&tourPreference)
	if err != nil {
		log.Println("Error while parsing json:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	createdTourPreference, err := handler.TourPreferenceService.Create(&tourPreference)
	if err != nil {
		log.Println("Error while creating a new tour preference:", err)
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(createdTourPreference)
}

func (handler *UserHandler) Get(writer http.ResponseWriter, reader *http.Request) {
	var id, _ = strconv.Atoi(chi.URLParam(reader, "id"))
	user, err := handler.UserService.Get(id)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	if user.Id != 0 {
		json.NewEncoder(writer).Encode(user)
	}
}

func (handler *UserHandler) GetProfile(writer http.ResponseWriter, reader *http.Request) {
	var id, _ = strconv.Atoi(chi.URLParam(reader, "id"))
	profile, err := handler.ProfileService.Get(id)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	if profile.Id != 0 {
		json.NewEncoder(writer).Encode(profile)
	}
}

func (handler *UserHandler) GetFollow(writer http.ResponseWriter, reader *http.Request) {
	var id, _ = strconv.Atoi(chi.URLParam(reader, "id"))
	follow, err := handler.FollowService.Get(id)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	if follow.Id != 0 {
		json.NewEncoder(writer).Encode(follow)
	}
}

func (handler *UserHandler) GetTourPreference(writer http.ResponseWriter, reader *http.Request) {
	var id, _ = strconv.Atoi(chi.URLParam(reader, "id"))
	tourPreference, err := handler.TourPreferenceService.Get(id)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	if tourPreference.Id != 0 {
		json.NewEncoder(writer).Encode(tourPreference)
	}
}

func (handler *UserHandler) GetAll(writer http.ResponseWriter, reader *http.Request) {
	users, err := handler.UserService.GetAll()
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(users)
}

func (handler *UserHandler) GetAllProfiles(writer http.ResponseWriter, reader *http.Request) {
	profiles, err := handler.ProfileService.GetAll()
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(profiles)
}

func (handler *UserHandler) GetAllFollows(writer http.ResponseWriter, reader *http.Request) {
	follows, err := handler.FollowService.GetAll()
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(follows)
}

func (handler *UserHandler) GetAllTourPreferences(writer http.ResponseWriter, reader *http.Request) {
	tourPreferences, err := handler.TourPreferenceService.GetAll()
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tourPreferences)
}

func (handler *UserHandler) Update(writer http.ResponseWriter, req *http.Request) {
	log.Println("Entering user update")
	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		log.Println("Error while parsing JSON:", err)
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
	// Setting the ID of the user from the URL
	user.Id = id

	err = handler.UserService.Update(&user)
	if err != nil {
		log.Println("Error while updating user:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(user)
}

func (handler *UserHandler) UpdateProfile(writer http.ResponseWriter, req *http.Request) {
	log.Println("Entering profile update")
	var profile model.Profile
	err := json.NewDecoder(req.Body).Decode(&profile)
	if err != nil {
		log.Println("Error while parsing JSON:", err)
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
	// Setting the ID of the profile from the URL
	profile.Id = id

	err = handler.ProfileService.Update(&profile)
	if err != nil {
		log.Println("Error while updating profile:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(profile)
}

func (handler *UserHandler) UpdateFollow(writer http.ResponseWriter, req *http.Request) {
	log.Println("Entering follow update")
	var follow model.Follow
	err := json.NewDecoder(req.Body).Decode(&follow)
	if err != nil {
		log.Println("Error while parsing JSON:", err)
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
	// Setting the ID of the follow from the URL
	follow.Id = id

	err = handler.FollowService.Update(&follow)
	if err != nil {
		log.Println("Error while updating follow:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(follow)
}

func (handler *UserHandler) UpdateTourPreference(writer http.ResponseWriter, req *http.Request) {
	log.Println("Entering tour preference update")
	var tourPreference model.TourPreference
	err := json.NewDecoder(req.Body).Decode(&tourPreference)
	if err != nil {
		log.Println("Error while parsing JSON:", err)
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
	// Setting the ID of the tour preference from the URL
	tourPreference.Id = id

	err = handler.TourPreferenceService.Update(&tourPreference)
	if err != nil {
		log.Println("Error while updating tour preference:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tourPreference)
}

func (handler *UserHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	log.Println("Entering user delete")
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		log.Println("Error while parsing ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.UserService.Delete(id)
	if err != nil {
		log.Println("Error while deleting user:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]string{"message": "User deleted successfully"})
}

func (handler *UserHandler) DeleteProfile(writer http.ResponseWriter, req *http.Request) {
	log.Println("Entering profile delete")
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		log.Println("Error while parsing ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.ProfileService.Delete(id)
	if err != nil {
		log.Println("Error while deleting profile:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]string{"message": "Profile deleted successfully"})
}

func (handler *UserHandler) DeleteFollow(writer http.ResponseWriter, req *http.Request) {
	log.Println("Entering follow delete")
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		log.Println("Error while parsing ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.FollowService.Delete(id)
	if err != nil {
		log.Println("Error while deleting follow:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]string{"message": "Follow deleted successfully"})
}

func (handler *UserHandler) DeleteTourPreference(writer http.ResponseWriter, req *http.Request) {
	log.Println("Entering tour preference delete")
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		log.Println("Error while parsing ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.TourPreferenceService.Delete(id)
	if err != nil {
		log.Println("Error while deleting tour preference:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]string{"message": "Tour preference deleted successfully"})
}
