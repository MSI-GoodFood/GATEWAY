package todo

import "github.com/gofrs/uuid"

type User struct {
	ID uuid.UUID `json:"id"`
	Email string `json:"email" default:""`
}

type Users []User

type Todo struct {
	ID uuid.UUID `json:"id"`
	Text string `json:"text" default:""`
	Done int `json:"done" default:"0"`
	UserId uuid.UUID  `json:"user_id"`
}

type Todos []Todo

type Account struct {
	User  *User   `json:"user"`
	Token string `json:"token"`
}

type TodoUpdate struct {
	Text string `json:"text" default:""`
	Done int `json:"done" default:"-1"`
}

type UserToken struct {
	User uuid.UUID `json:"id"`
	Token string `json:"token"`
}
