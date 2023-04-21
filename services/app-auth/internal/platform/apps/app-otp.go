package apps

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/config"
	"github.com/AppsLab-KE/be-go-gen-grpc/otp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewOTPServiceClient(config config.OtpService) (otp.OtpServiceClient, error) {
	conn, err := grpc.Dial(config.Host+":"+config.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := otp.NewOtpServiceClient(conn)
	return client, nil
}
