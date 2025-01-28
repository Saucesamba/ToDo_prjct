// hhtp-request handlers
package handlers

import (
	"MyProject/internal/app"
	"MyProject/internal/models"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
)

type Handler struct {
	Repo sql.DB
}

func NewHandler(repo sql.DB) *Handler {
	return &Handler{Repo: repo}
}

func (h *Handler) HandleUserRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	var user models.UserRegisterJSON
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
		return
	}
	createdUser, err := app.RegistrUser(&h.Repo, user.Email, user.Name, user.Password)
	if err != nil {
		http.Error(w, "Unable to register user", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.UserResponseJSON{
		Id:    createdUser.Id,
		Name:  createdUser.Name,
		Email: createdUser.Email,
	})
}

func (h *Handler) HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
	}
	var user models.UserLoginJSON
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
	}

	loginUser, err := app.AuthUser(&h.Repo, user.Email, user.Password)
	if err != nil {
		http.Error(w, "Unable to login", http.StatusUnauthorized)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.UserResponseJSON{
		Id:    loginUser.Id,
		Name:  loginUser.Name,
		Email: loginUser.Email,
	})
}
