package entity

type User struct {
	Name        string
	Email       string
	PhoneNumber string
}

type Otp struct {
	PhoneNumber string
	Code        string
}
