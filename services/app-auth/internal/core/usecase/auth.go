package usecase

type AuthUsecase interface {
	Login() error
	Logout() error
	Register() error
}
