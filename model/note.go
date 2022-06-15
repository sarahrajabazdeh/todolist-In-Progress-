package model

type Note struct {
	ID          int
	Description string
	Title       string
	UserId      User
}
