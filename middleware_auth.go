package main

import (
	"fmt"
	"net/http"

	"github.com/isaquefranklin/rss-scraper/auth"
	"github.com/isaquefranklin/rss-scraper/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %w", err))
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apikey)
		if err != nil {
			respondWithError(w, 404, fmt.Sprintf("Couldn't get user: %w", err))
			return
		}

		handler(w, r, user)

	}
}
