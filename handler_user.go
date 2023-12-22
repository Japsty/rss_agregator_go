package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/japsty/rssagg/internal/cmd"
	"github.com/japsty/rssagg/internal/database"
	"github.com/japsty/rssagg/internal/middleware"
	"github.com/japsty/rssagg/internal/models"
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
		middleware.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
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
		middleware.RespondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	middleware.RespondWithJSON(w, 201, models.DatabaseUserToUser(user))
}

func (apiCfg *main.apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	middleware.RespondWithJSON(w, 200, models.DatabaseUserToUser(user))
}

func (apiCfg *main.apiConfig) handlerGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := apiCfg.DB.GetAllUsers(r.Context())
	if err != nil {
		middleware.RespondWithError(w, 400, fmt.Sprintf("Couldn't get users: %v", err))
		return
	}

	middleware.RespondWithJSON(w, 201, users)
}

func (apiCfg *main.apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		middleware.RespondWithError(w, 400, fmt.Sprintf("Couldn't get posts: %v", err))
		return
	}

	middleware.RespondWithJSON(w, 200, models.DatabasePostsToPosts(posts))
}
