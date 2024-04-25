package main

import (
	"context"
	"database-example/handler"
	"database-example/repo"
	"database-example/service"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func startServer(handler *handler.UserHandler) {

	router := mux.NewRouter().StrictSlash(true)
	url := "/followers/"
	router.HandleFunc(url+"if-following/{id1}/{id2}", handler.CheckIfFollowing).Methods("GET")
	router.HandleFunc(url+"follow/{id1}/{id2}", handler.Follow).Methods("POST")
	router.HandleFunc(url+"get-recommendations/{id}", handler.GetRecommendation).Methods("GET")
	log.Fatal(http.ListenAndServe(":7007", router))
}
func initDB() neo4j.DriverWithContext {

	uri := "bolt://localhost:7687"
	user := "neo4j"
	pass := "123456789"
	auth := neo4j.BasicAuth(user, pass, "")

	driver, err := neo4j.NewDriverWithContext(uri, auth)
	if err != nil {
		fmt.Println(err)
		return nil

	}

	return driver

}
func CheckConnection(driver neo4j.DriverWithContext) {
	ctx := context.Background()
	err := driver.VerifyConnectivity(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func cleanData(driver neo4j.DriverWithContext) error {

	ctx := context.Background()
	session := driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			_, err := transaction.Run(
				ctx,
				`MATCH (n) DETACH DELETE n`,
				nil)
			return nil, err
		})
	return err
}

func migrateData(driver neo4j.DriverWithContext) error {

	ctx := context.Background()
	session := driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)
	fmt.Printf("Neo %s\n", session)

	_, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			_, err := transaction.Run(
				ctx,
				`MERGE (u1: User {id: 1})
				 MERGE (u2: User {id: 2})
				 MERGE (u3: User {id: 3})
				 MERGE (u4: User {id: 4})
				 MERGE (u5: User {id: 5})
				 MERGE (u6: User {id: 6})
				 CREATE (u1) -[:Following]->(u2)
				 CREATE (u1) -[:Following]->(u5)
				 CREATE (u2) -[:Following]->(u6)
				 CREATE (u3) -[:Following]->(u2)
				 CREATE (u3) -[:Following]->(u5)
				 CREATE (u4) -[:Following]->(u1)
				 `,
				nil)
			return nil, err
		})
	return err
}

func main() {
	timeoutContext, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	driver := initDB()

	defer driver.Close(timeoutContext)
	CheckConnection(driver)

	if driver == nil {
		return
	}

	cleanData(driver)
	migrateData(driver)

	followerRepository := &repo.FollowRepository{driver}
	followerService := &service.FollowService{followerRepository}
	followerHandler := &handler.UserHandler{followerService}

	startServer(followerHandler)
}
