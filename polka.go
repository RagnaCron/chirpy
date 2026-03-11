package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/ragnacron/chirpy/internal/database"
)

func (cfg *apiConfig) polkaHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID string `json:"user_id"`
		} `json:"data"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if params.Event != "user.upgraded" {
		respondWithJSON(w, http.StatusNoContent, nil)
		return
	}

	userID, err := uuid.Parse(params.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Invalid user id", err)
		return
	}

	err = cfg.db.UpgradeUserToChirpyRedByID(r.Context(), database.UpgradeUserToChirpyRedByIDParams{
		ID:          userID,
		IsChirpyRed: true,
	})
	if err != nil {
		respondWithError(w, http.StatusNotFound, "User not found", err)
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}
