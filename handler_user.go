package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jarednil/go-art/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parametres struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parametres{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing to JSON %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error of creating user %v", err))
		return
	}

	respondWithJSON(w, 201, databaseUsertoUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, databaseUsertoUser(user))
}
