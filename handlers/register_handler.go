package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/lifebee66-afk/NarusheniamNET/dto"
	"github.com/lifebee66-afk/NarusheniamNET/pkg/httphelper"
	"github.com/lifebee66-afk/NarusheniamNET/storage"
)

type Handler struct {
	DB *sql.DB
}

func NewHandler(DB *sql.DB) *Handler {
	return &Handler{DB: DB}
}

func (h *Handler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {

	validate := validator.New()
	var RegDto *dto.RegisterUserDTO
	if err := json.NewDecoder(r.Body).Decode(&RegDto); err != nil {
		http.Error(w, "failed decode json", http.StatusInternalServerError)
		return
	}

	if err := validate.Struct(RegDto); err != nil {
		http.Error(w, "data not complate validate "+err.Error(), http.StatusBadRequest)
		return
	}

	save := storage.NewRegisterUser(h.DB)
	if err := save.RegisterUser(r.Context(), RegDto); err != nil {
		http.Error(w, "failed save user in database", http.StatusInternalServerError)
		return
	}

	if err := httphelper.ResponseJSON(w, http.StatusCreated, map[string]string{"succes": "ok"}); err != nil {
		http.Error(w, "failed get data", http.StatusInternalServerError)
		return
	}

}
