package main

import (
	"am.ca-server/controllers"
	"am.ca-server/middleware"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
)

const (
	// PORT : port that server is hosted under
	PORT = ":8088"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.Use(middleware.LoggingMiddleware)

	router.HandleFunc("/", controllers.BaseURL).Methods("GET")
	router.HandleFunc("/email", controllers.EmailService).Methods("POST")
	router.HandleFunc("/resume", controllers.ResumeJSON).Methods("GET")

	// CORS middleware
	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(PORT, handler))
}

func main() {
	fmt.Println("Running server on port", PORT)

	// load dotenv variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	handleRequests()
}
