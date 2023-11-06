package handlers

import (
	"github.com/japsty/rssagg"
	"net/http"
)

func handlerErr(w http.ResponseWriter, r *http.Request) {
	main.respondWithError(w, 400, "Something went wrong")
}
