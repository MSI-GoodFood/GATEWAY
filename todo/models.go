package todo

import "github.com/gofrs/uuid"

type User struct {
	ID uuid.UUID `json:"id"`
	Email string `json:"email"`
}

type Users []User

type SignUp struct {
	User  *User   `json:"user"`
	Token string `json:"token"`
}

type Todo struct {
	ID uuid.UUID `json:"id"`
	Text string `json:"text"`
	Done int `json:"done"`
	UserId uuid.UUID  `json:"user_id"`
}

type Todos []Todo
