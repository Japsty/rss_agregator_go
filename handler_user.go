package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/japsty/rssagg"
	"github.com/japsty/rssagg"
	"github.com/japsty/rssagg/internal/database"
	"net/http"
	"time"
)

func (apiCfg *main.apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		main.respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	location, _ := time.LoadLocation("Europe/Moscow")
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC().In(location),
		UpdatedAt: time.Now().UTC().In(location),
		Name:      params.Name,
	})
	if err != nil {
		main.respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	main.respondWithJSON(w, 201, main.databaseUserToUser(user))
}

func (apiCfg *main.apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	main.respondWithJSON(w, 200, main.databaseUserToUser(user))
}

func (apiCfg *main.apiConfig) handlerGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := apiCfg.DB.GetAllUsers(r.Context())
	if err != nil {
		main.respondWithError(w, 400, fmt.Sprintf("Couldn't get users: %v", err))
		return
	}

	main.respondWithJSON(w, 201, users)
}
