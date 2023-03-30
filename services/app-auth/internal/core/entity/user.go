package entity

type User struct {
	Name        string
	Email       string
	PhoneNumber string
	Hash        string
}

type Otp struct {
	PhoneNumber string
	Code        string
}
