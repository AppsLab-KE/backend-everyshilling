package entity

type User struct {
	UserId      string `json:"user_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Hash        string `json:"-"`
	Verified    bool   `json:"verified"`
}

type Otp struct {
	PhoneNumber string
	Code        string
}
