// hhtp-request handlers
package handlers

import (
	"MyProject/internal/app"
	"MyProject/internal/models"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
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
	log.Println("Method: ", r.Method, " Url: ", r.URL, " UserRegister")

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
		Id:       createdUser.Id,
		Name:     createdUser.Name,
		Email:    createdUser.Email,
		TaskStat: models.UserTaskInfo{CompletedCount: 0, TaskCount: 0},
	})
}
func (h *Handler) HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	log.Println("Method: ", r.Method, " Url: ", r.URL, " UserLogin")

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
	} else {
		taskStat, _ := app.GetTaskStat(&h.Repo, loginUser.Id)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.UserResponseJSON{
			Id:       loginUser.Id,
			Name:     loginUser.Name,
			Email:    loginUser.Email,
			TaskStat: taskStat,
		})
	}
}
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request, id int) {
	log.Println("Method: ", r.Method, " Url: ", r.URL, " updateInfo")
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
	err = app.UpdateUser(&h.Repo, user)
	if err != nil {
		http.Error(w, "Unable to update user", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	taskStat, _ := app.GetTaskStat(&h.Repo, id)
	json.NewEncoder(w).Encode(models.UserResponseJSON{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		TaskStat: taskStat,
	})
}
func (h *Handler) GetInfo(w http.ResponseWriter, r *http.Request, id int) {
	log.Println("Method: ", r.Method, " Url: ", r.URL, " getInfo")
	user, err := app.GetInfoUser(&h.Repo, id)
	if err != nil {
		http.Error(w, "Unable to find user", http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	taskStat, _ := app.GetTaskStat(&h.Repo, id)
	json.NewEncoder(w).Encode(models.UserResponseJSON{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		TaskStat: taskStat,
	})
}
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

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request, params url.Values) {
	//на будущее тут будет фильрация и пагинация тасков
	log.Println("Method: ", r.Method, " Url: ", r.URL, "GetAllTasks")
	userId := params.Get("userId")
	if userId == "" {
		http.Error(w, "Missing user id", http.StatusBadRequest)
		return
	}
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	//filter := params.Get("filter")
	//log.Println(filter, userIdInt)

	tasks, err := app.GetUserTasks(&h.Repo, userIdInt)
	if len(tasks) == 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.UsertasksResp{})
	} else {
		var taskResp []models.OneTaskResponse
		for _, task := range tasks {
			taskResp = append(taskResp, models.OneTaskResponse{Id: task.Id, Description: task.Description, Name: task.Name, Completed: task.Completed})
		}
		if err != nil {
			http.Error(w, "Unable to find tasks", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(taskResp)
	}
}
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request, userId int) {

	log.Println("Method: ", r.Method, " Url: ", r.URL, " CreateTask")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
	}
	var task models.TaskReq
	err = json.Unmarshal(body, &task)
	if err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
	}

	var newTask models.Task
	newTask.Description = task.Description
	newTask.Name = task.Name
	newTask.UserId = userId
	newTask.Completed = false

	err = app.CreateTask(&h.Repo, newTask, userId)
	if err != nil {
		http.Error(w, "Unable to create task", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var respTask = models.OneTaskResponse{Id: newTask.Id, Description: newTask.Description, Name: newTask.Name, Completed: newTask.Completed}
	json.NewEncoder(w).Encode(respTask)
}
func (h *Handler) TasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		return
	}

	params := r.URL.Query()
	userId := params.Get("userId")
	if userId == "" {
		http.Error(w, "Missing user id", http.StatusBadRequest)
		return
	}
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	taskId := params.Get("taskId")
	taskIdInt, _ := strconv.Atoi(taskId)
	switch r.Method {
	case http.MethodGet:
		h.GetTasks(w, r, params)
	case http.MethodPost:
		h.CreateTask(w, r, userIdInt)
	case http.MethodPut:
		log.Println("Updating task", taskIdInt, userIdInt)
		app.UpdateTaskStatus(&h.Repo, taskIdInt)
	case http.MethodDelete:
		log.Println("Deleted task", taskIdInt, userIdInt)
		app.DeleteTask(&h.Repo, taskIdInt)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
