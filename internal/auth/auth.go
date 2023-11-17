package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey достает API ключ из хедера HTTP запроса
// Authorization: ApiKey {apikey}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("incorrect auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed auth header")
	}
	return vals[1], nil
}
