//models layer

package models

type Task struct {
	Id          int
	Name        string
	Description string
	Completed   bool
	UserId      int
} //Структура для взаимодействия с БД

type OneTaskResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
type UsertasksResp struct {
	Tasks []OneTaskResponse `json:"tasks"`
}

type TaskReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
