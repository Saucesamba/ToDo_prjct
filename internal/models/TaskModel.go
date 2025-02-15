package models

import "time"

type Task struct {
	Id          int
	Name        string
	Description string
	Completed   bool
	UserId      int
	CreatedAt   time.Time
} //Структура для взаимодействия с БД

type OneTaskResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}
type UsertasksResp struct {
	Tasks []OneTaskResponse `json:"tasks"`
}

type TaskReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
