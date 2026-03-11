package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/ragnacron/chirpy/internal/auth"
	"github.com/ragnacron/chirpy/internal/database"
)

func (cfg *apiConfig) revokeHandler(w http.ResponseWriter, r *http.Request) {
	refreshTokenString, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find refresh token", err)
		return
	}

	err = cfg.db.RevokeToken(r.Context(), database.RevokeTokenParams{
		Token:     refreshTokenString,
		RevokedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't revoke token", err)
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}
