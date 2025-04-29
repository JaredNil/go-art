package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load() // Reading environment
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT string is not be empty")
	}

	router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	fmt.Printf("Сервер стартует на %s порту", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	// http.ListenAndServe(port, router)

}
