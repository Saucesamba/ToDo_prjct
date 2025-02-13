package models

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

type UserJSON struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginJSON struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponseJSON struct {
	Id       int          `json:"id"`
	Name     string       `json:"name"`
	Email    string       `json:"email"`
	TaskStat UserTaskInfo `json:"task_stat"`
}

type UserTaskInfo struct {
	TaskCount      int `json:"task_count"`
	CompletedCount int `json:"completed_count"`
}
