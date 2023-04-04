package storage

type Db interface {
	CreateOtp(data any) error
	CreateUser(user interface{}) (interface{}, error)
}
