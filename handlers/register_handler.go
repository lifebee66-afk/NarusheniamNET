package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lifebee66-afk/NarusheniamNET/dto"
)

type Handler struct {
}

func NewHandler() {

}

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var RegDto *dto.RegisterUserDTO
	if err := json.NewDecoder(r.Body).Decode(&RegDto); err != nil {
		http.Error(w, "failed decode json", http.StatusInternalServerError)
		return
	}

}
