package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cyub/go-kit-demo/account/model"
	"github.com/gorilla/mux"
)

// RegisterRequest define register request struct
type RegisterRequest struct {
	Name   string `json:"name"`
	Gender int    `json:"gender"`
}

// Register is registr handler
func Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid  params"})
		return
	}

	var (
		name   = req.Name
		gender = req.Gender
	)

	if len(name) <= 0 || len(name) > 20 {
		RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid name param"})
		return
	}

	account := &model.Account{}
	if err := account.Create(name, gender); err != nil {
		RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]uint{"id": account.ID})
}

// Show is show user info struct
func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])
	account := &model.Account{}

	if err := account.Find(userID); err != nil {
		RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	RespondWithJSON(w, http.StatusBadRequest, account)
}
