package main

import (
	"net/http"
	"time"

	"github.com/ragnacron/chirpy/internal/auth"
)

func (cfg *apiConfig) refreshHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Token string `json:"token"`
	}

	refreshTokenString, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find refresh token", err)
		return
	}

	rt, err := cfg.db.GetRefreshToken(r.Context(), refreshTokenString)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't get refresh token", err)
		return
	}
	if rt.RevokedAt.Valid {
		respondWithError(w, http.StatusUnauthorized, "Token was revoked", err)
		return
	}
	if rt.ExpiresAt.After(time.Now().UTC()) {
		respondWithError(w, http.StatusUnauthorized, "Token has expired", err)
		return
	}

	token, err := auth.MakeJWT(rt.UserID, cfg.secret, time.Duration(time.Hour))
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't create access JWT", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{Token: token})

}
