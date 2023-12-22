package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/japsty/rssagg/internal/database"
	"github.com/japsty/rssagg/internal/middleware"
	"github.com/japsty/rssagg/internal/models"
	"net/http"
	"time"
)

func (apiCfg apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		middleware.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	location, _ := time.LoadLocation("Europe/Moscow")
	feedFollow, err := apiCfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC().In(location),
		UpdatedAt: time.Now().UTC().In(location),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		middleware.RespondWithError(w, 400, fmt.Sprintf("Couldn't create feed follow: %v", err))
		return
	}

	middleware.RespondWithJSON(w, 201, models.DatabaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		middleware.RespondWithError(w, 400, fmt.Sprintf("Couldn't get feed follows: %v", err))
		return
	}

	middleware.RespondWithJSON(w, 201, feedFollows)
}

func (apiCfg apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowId, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		middleware.RespondWithError(w, 400, fmt.Sprintf("Couldn't get feed follow id: %v", err))
		return
	}
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})
	if err != nil {
		middleware.RespondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}
	middleware.RespondWithJSON(w, 200, struct{}{})
}
