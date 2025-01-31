// hhtp-request handlers
package handlers

import (
	"MyProject/internal/app"
	"MyProject/internal/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	Repo sql.DB
}

func NewHandler(repo sql.DB) *Handler {
	return &Handler{Repo: repo}
}

func (h *Handler) HandleUserRegister(w http.ResponseWriter, r *http.Request) {
	log.Println("Method", r.Method)
	log.Println("Url", r.URL)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	var user models.UserJSON
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
	log.Println("Method", r.Method)
	log.Println("Url", r.URL)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPut {
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

// хэндлер для обновлении информации о пользователе, возвращает измененные данные если все прошло хорошо

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request, id int) {
	log.Println("Method", r.Method)
	fmt.Println("Url", r.URL, "updateInfo")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
	}

	var userReq models.UserJSON
	err = json.Unmarshal(body, &userReq)
	if err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
	}
	var user models.User = models.User{Email: userReq.Email, Name: userReq.Name, Password: userReq.Password, Id: id}
	log.Println(user)
	err = app.UpdateUser(&h.Repo, user)
	if err != nil {
		http.Error(w, "Unable to update user", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.UserResponseJSON{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	})
}

func (h *Handler) GetInfo(w http.ResponseWriter, r *http.Request, id int) {
	log.Println("Method", r.Method)
	log.Println("Url", r.URL, "getInfo")

	user, err := app.GetInfoUser(&h.Repo, id)
	log.Println(user)
	if err != nil {
		http.Error(w, "Unable to find user", http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.UserResponseJSON{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	})
}

// хендлер для изменения и получения инфы о пользователе
func (h *Handler) UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if parts[1] != "users" {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		h.GetInfo(w, r, id)
	case http.MethodPut:
		h.UpdateUser(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
