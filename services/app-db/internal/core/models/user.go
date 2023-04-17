package models

type User struct {
	Model
	Name     string
	Email    string
	Phone    string
	Password string
}
