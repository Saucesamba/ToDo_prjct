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
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdatePasswordJSON struct {
	Password string `json:"password"`
}
