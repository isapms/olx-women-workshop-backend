package main

import (
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"

	"olx-women-workshop-2022-backend/database"
	"olx-women-workshop-2022-backend/handlers"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	database.CreateConn()

	router := mux.NewRouter()
	router.HandleFunc("/api/ads", handlers.List).Methods(http.MethodGet)
	router.HandleFunc("/api/ads", handlers.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/ads/{id}", handlers.Delete).Methods(http.MethodDelete)

	router.
		PathPrefix("/static/").
		Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static/"))))

	allowedMethods := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodDelete,
		http.MethodPatch,
		http.MethodOptions,
	}

	handler := cors.
		New(cors.Options{AllowedMethods: allowedMethods}).
		Handler(router)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), handler)
	if err != nil {
		log.Fatal(err)
	}
}
