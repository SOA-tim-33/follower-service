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

// func initDB() *gorm.DB {
// 	connectionStr := "root:root@tcp(localhost:3306)/students?charset=utf8mb4&parseTime=True&loc=Local"
// 	database, err := gorm.Open(mysql.Open(connectionStr), &gorm.Config{})
// 	if err != nil {
// 		print(err)
// 		return nil
// 	}

// 	database.AutoMigrate(&model.BlogComment{})
// 	database.AutoMigrate(&model.Rating{})
// 	database.AutoMigrate(&model.UserBlog{})
// 	database.AutoMigrate(&model.UserBlogTourReport{})

// 	//database.Exec("INSERT IGNORE INTO students VALUES ('aec7e123-233d-4a09-a289-75308ea5b7e6', 'Marko Markovic', 'Graficki dizajn')")

// 	return database
// }

//func startServer(blog_comment_handler *handler.BlogCommentHandler, rating_handler *handler.RatingHandler, user_blog_handler *handler.UserBlogHandler, user_blog_tour_report_handler *handler.UserBlogTourReportHandler) {
// 	router := mux.NewRouter().StrictSlash(true)

// 	router.HandleFunc("/blogcomments", blog_comment_handler.Create).Methods("POST")
// 	router.HandleFunc("/blogcomments/{id}", blog_comment_handler.Get).Methods("GET")
// 	router.HandleFunc("/blogcomments/all", blog_comment_handler.GetAll).Methods("GET")
// 	router.HandleFunc("/blogcomments", blog_comment_handler.Update).Methods("PUT")
// 	router.HandleFunc("/blogcomments", blog_comment_handler.Delete).Methods("DELETE")

// 	router.HandleFunc("/ratings", rating_handler.Create).Methods("POST")
// 	router.HandleFunc("/ratings/{id}", rating_handler.Get).Methods("GET")
// 	router.HandleFunc("/ratings/all", rating_handler.GetAll).Methods("GET")
// 	router.HandleFunc("/ratings", rating_handler.Update).Methods("PUT")
// 	router.HandleFunc("/ratings", rating_handler.Delete).Methods("DELETE")

// 	router.HandleFunc("/userblogs", user_blog_handler.Create).Methods("POST")
// 	router.HandleFunc("/userblogs/{id}", user_blog_handler.Get).Methods("GET")
// 	// router.HandleFunc("/userblogs/all", user_blog_handler.GetAll).Methods("GET")
// 	// router.HandleFunc("/userblogs", user_blog_handler.Update).Methods("PUT")
// 	// router.HandleFunc("/userblogs", user_blog_handler.Delete).Methods("DELETE")

// 	router.HandleFunc("/userblogtourreports", user_blog_tour_report_handler.Create).Methods("POST")
// 	router.HandleFunc("/userblogtourreports/{id}", user_blog_tour_report_handler.Get).Methods("GET")
// 	router.HandleFunc("/userblogtourreports/all", user_blog_tour_report_handler.GetAll).Methods("GET")
// 	router.HandleFunc("/userblogtourreports", user_blog_tour_report_handler.Update).Methods("PUT")
// 	router.HandleFunc("/userblogtourreports", user_blog_tour_report_handler.Delete).Methods("DELETE")

// 	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
// 	println("Server starting")
// 	log.Fatal(http.ListenAndServe(":8080", router))
// }

// func main() {
// 	database := initDB()
// 	if database == nil {
// 		print("FAILED TO CONNECT TO DB")
// 		return
// 	}

// 	blog_comment_repo := &repo.BlogCommentRepository{DatabaseConnection: database}
// 	blog_comment_service := &service.BlogCommentService{BlogCommentRepo: blog_comment_repo}
// 	blog_comment_handler := &handler.BlogCommentHandler{BlogCommentService: blog_comment_service}

// 	rating_repo := &repo.RatingRepository{DatabaseConnection: database}
// 	rating_service := &service.RatingService{RatingRepo: rating_repo}
// 	rating_handler := &handler.RatingHandler{RatingService: rating_service}

// 	user_blog_repo := &repo.UserBlogRepository{DatabaseConnection: database}
// 	user_blog_service := &service.UserBlogService{UserBlogRepo: user_blog_repo}
// 	user_blog_handler := &handler.UserBlogHandler{UserBlogService: user_blog_service}

// 	user_blog_tour_report_repo := &repo.UserBlogTourReportRepository{DatabaseConnection: database}
// 	user_blog_tour_report_service := &service.UserBlogTourReportService{UserBlogTourReportRepo: user_blog_tour_report_repo}
// 	user_blog_tour_report_handler := &handler.UserBlogTourReportHandler{UserBlogTourReportService: user_blog_tour_report_service}

//		startServer(blog_comment_handler, rating_handler, user_blog_handler, user_blog_tour_report_handler)
//	}
func main() {
	loadConfig()

	database := initDB()

	blogRepository := repo.UserBlogRepository{}
	blogRepository.Init(database)
	blogCommentRepository := repo.BlogCommentRepository{}
	blogCommentRepository.Init(database)
	ratingRepository := repo.RatingRepository{}
	ratingRepository.Init(database)

	blogService := service.UserBlogService{}
	blogService.Init(&blogRepository)
	blogCommentService := service.BlogCommentService{}
	blogCommentService.Init(&blogCommentRepository)
	ratingService := service.RatingService{}
	ratingService.Init(&ratingRepository)

	blogHandler := handler.UserBlogHandler{}

	router := blogHandler.InitRouter(&blogService, &blogCommentService, &ratingService)
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
