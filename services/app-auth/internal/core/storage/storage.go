package storage

type Db interface {
	CreateOtp(data any) error
}
