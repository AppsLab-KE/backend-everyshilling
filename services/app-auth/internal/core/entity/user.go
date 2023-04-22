package entity

type User struct {
	UserId      string
	Name        string
	Email       string
	PhoneNumber string
	Hash        string `json:"-"`
	Verified    bool
}

type Otp struct {
	PhoneNumber string
	Code        string
}
