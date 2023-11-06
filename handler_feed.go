package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/japsty/rssagg"
	"github.com/japsty/rssagg/internal/database"
	"net/http"
	"time"
)

func (apiCfg *main.apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		main.respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	location, _ := time.LoadLocation("Europe/Moscow")
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC().In(location),
		UpdatedAt: time.Now().UTC().In(location),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		main.respondWithError(w, 400, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	main.respondWithJSON(w, 201, main.databaseFeedToFeed(feed))
}

func (apiCfg *main.apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		main.respondWithError(w, 400, fmt.Sprintf("Couldn't get feed: %v", err))
		return
	}

	main.respondWithJSON(w, 201, feeds)
}