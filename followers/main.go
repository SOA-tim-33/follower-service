package main

import (
	"database-example/handler"
	"database-example/repo"
	"database-example/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	loadConfig()

	database := initDB()

	userRepository := repo.UserRepository{}
	userRepository.Init(database)
	profileRepository := repo.ProfileRepository{}
	profileRepository.Init(database)
	followRepository := repo.FollowRepository{}
	followRepository.Init(database)
	tourPreferenceRepository := repo.TourPreferenceRepository{}
	tourPreferenceRepository.Init(database)

	userService := service.UserService{}
	userService.Init(&userRepository)
	profileService := service.ProfileService{}
	profileService.Init(&profileRepository)
	followService := service.FollowService{}
	followService.Init(&followRepository)
	tourPreferenceService := service.TourPreferenceService{}
	tourPreferenceService.Init(&tourPreferenceRepository)

	userHandler := handler.UserHandler{}

	router := userHandler.InitRouter(&userService, &profileService, &followService, &tourPreferenceService)
	fmt.Println("Encounters micro-service running")
	http.ListenAndServe(":7007", router)

}
func loadConfig() {
	envErr := godotenv.Load("config/.env")

	if envErr != nil {
		log.Fatalf(envErr.Error())
	}
}

/* not sure where to put this*/
func initDB() *gorm.DB {
	/* TODO: lazy to think of something easier: */
	dbType := os.Getenv("DATABASE_TYPE")
	dbUser := os.Getenv("DATABASE_USER")
	dbSecret := os.Getenv("DATABASE_SECRET")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")
	connectionUrl := fmt.Sprintf("%s://%s:%s@%s:%s/%s", dbType, dbUser, dbSecret, dbHost, dbPort, dbName)
	database, databaseErr := gorm.Open(postgres.Open(connectionUrl), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		NoLowerCase: true,
	}})
	if databaseErr != nil {
		log.Fatalf(databaseErr.Error())
		return nil
	}
	return database
}
