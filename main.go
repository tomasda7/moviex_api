package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/tomasda7/moviex_api/controllers"
	"github.com/tomasda7/moviex_api/services"
)

func main() {
	// db conn
	err := services.EstablishDbConn()

	if err != nil {
		log.Fatal(err)
	}

	// router
	router := mux.NewRouter()

	// movies routes
	router.HandleFunc("/api/v1/moviex/detail/{movie_id:[0-9]+}", controllers.GetMovieDetail).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/moviex/most_viewed", controllers.GetMostViewedMovies).Methods(http.MethodGet)

	// comments routes
	router.HandleFunc("/api/v1/moviex/add_comment/{movie_id:[0-9]+}", controllers.AddComment).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/moviex/remove_comment/{user_id:[0-9]+}/{comment_id:[0-9]+}", controllers.RemoveComment).Methods(http.MethodDelete)
	router.HandleFunc("/api/v1/moviex/edit_comment/{user_id:[0-9]+}/{comment_id:[0-9]+}", controllers.EditComment).Methods(http.MethodPut)

	//users routes
	router.HandleFunc("/api/v1/moviex/update_user/{user_id:[0-9]+}", controllers.UpdateUserInfo).Methods(http.MethodPut)

	// CORS config
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5433"},
		AllowedMethods: []string{"GET", "POST"},
	})

	// launching server
	handler := corsOptions.Handler(router)
	PORT := ":9090"

	if err := startServer(PORT, handler); err != nil {
		log.Fatalf("Error while trying to launch server \n%v", err)
	}
}

func startServer(port string, router http.Handler) error {
	server := &http.Server{
		Handler: router,
		Addr: port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	fmt.Println("Server started at <http://localhost" + port + ">" + " ...")

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
