//models layer

package models

type Task struct {
	Id          int
	Name        string
	Description string
	Completed   bool
	UserId      int
} //Структура для взаимодействия с БД

type TaskResponse struct {
	Id          int
	Name        string
	Description string
	Completed   bool
}

type TaskReq struct {
	Name        string
	Description string
}
