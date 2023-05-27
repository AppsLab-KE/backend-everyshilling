package models

type User struct {
	BaseModel
	Name     string
	Email    string
	Phone    string
	Password string
	Verified bool
}
