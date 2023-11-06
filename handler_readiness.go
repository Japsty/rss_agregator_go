package handlers

import (
	"github.com/japsty/rssagg"
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	main.respondWithJSON(w, 200, struct{}{})
}
