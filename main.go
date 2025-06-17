package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jarednil/go-art/internal/database"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load() // Reading environment
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT string is not be empty")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("dbURL string is not be empty")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Cannot create database.", err)
	}

	apiCfg := apiConfig{
		DB: database.New((conn)),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1 := chi.NewRouter()
	v1.Get("/ready", handlerReadiness)
	v1.Get("/error", handlerError)
	v1.Post("/users", apiCfg.handlerCreateUser)
	v1.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))

	v1.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))

	router.Mount("/v1", v1)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	fmt.Printf("Сервер стартует на %v порту", port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
