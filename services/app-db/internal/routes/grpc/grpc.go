package grpc

import "appslab.co.ke/everyshilling/app-db/internal/core/ports"

type Serve struct {
	useController ports.UserRepo
}
