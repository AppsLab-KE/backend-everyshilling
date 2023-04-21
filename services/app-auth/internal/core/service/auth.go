package service

import (
	"context"
	"errors"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/config"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/pkg/hash"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/pkg/tokens"
	log "github.com/sirupsen/logrus"
)

type AuthService struct {
	repo   adapters.AuthRepo
	config config.Jwt
}

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserExists        = errors.New("user exists")
	ErrUserCreation      = errors.New("io error while creating user. see server logs")
	ErrValidationError   = errors.New("validation error")
	ErrRequestValidation = errors.New("validation  error")
)

func (d AuthService) SendResetOTP(request dto.OtpGenReq) (*dto.OtpGenRes, error) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	// check if user with phone number exists
	user, _ := d.repo.GetUserByPhone(ctx, request.Phone)
	if user != nil {
		return nil, ErrUserNotFound
	}

	// if user exists, send otp
	otpRes, err := d.repo.CreateOtpCode(ctx, request)
	if user != nil {
		return nil, err
	}

	return otpRes, nil
}

func (d AuthService) CreateUser(registerRequest dto.RegisterRequest) (*dto.UserRegistrationRes, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// check if registerRequest exists by phone
	user, _ := d.repo.GetUserByEmail(ctx, registerRequest.Email)
	if user != nil {
		return nil, ErrUserExists
	}

	user, _ = d.repo.GetUserByPhone(ctx, registerRequest.PhoneNumber)
	if user != nil {
		return nil, ErrUserExists
	}

	// create password hash
	hashedPassword, err := hash.GenerateHash(registerRequest.Password)
	if err != nil {
		log.Errorf("password hashing error: %s", err)
		return nil, ErrUserCreation
	}

	registerRequest.Password = hashedPassword

	// create user
	createdUser, err := d.repo.CreateUser(ctx, registerRequest)
	if err != nil {
		log.Errorf("user creation error: %s", err)
		return nil, ErrUserCreation
	}

	// generate jwt
	jwtToken, err := tokens.GenerateToken(createdUser.UserId, d.config.Secret, d.config.ExpiryMinutes)
	if err != nil {
		log.Errorf("jwt generation error: %s", err)
		return nil, ErrUserCreation
	}

	res := dto.UserRegistrationRes{
		User:  *createdUser,
		Token: jwtToken,
	}

	return &res, nil
}

func (d AuthService) VerifyResetOTP(request dto.OtpVerificationReq) (*dto.OtpVerificationRes, error) {
	//TODO implement me
	panic("implement me")
}

func (d AuthService) ChangePassword(request dto.RequestResetCredentials) (*dto.ResetRes, error) {
	//TODO implement me
	panic("implement me")
}

func (d AuthService) SendLoginOtp(request dto.LoginInitReq) (*dto.LoginInitRes, error) {
	//TODO implement me
	panic("implement me")
}

func (d AuthService) VerifyLoginOtp(request dto.OtpVerificationReq) (*dto.OtpVerificationRes, error) {
	//TODO implement me
	panic("implement me")
}

func (d AuthService) ResendLoginOTP(request dto.ResendLoginOTPReq) (*dto.ResendLoginOTPRes, error) {
	//TODO implement me
	panic("implement me")
}

func (d AuthService) SendVerifyPhoneOTP(verificationRequest dto.OtpVerificationReq) (*dto.OtpVerificationRes, error) {
	//TODO implement me
	panic("implement me")
}

func (d AuthService) VerifyPhoneOTP(verificationRequest dto.OtpVerificationReq) (*dto.OtpVerificationRes, error) {
	//TODO implement me
	panic("implement me")
}

func (d AuthService) ResendVerifyPhoneOTP(request dto.ResendLoginOTPReq) (*dto.ResendLoginOTPRes, error) {
	//TODO implement me
	panic("implement me")
}

func NewDefaultAuthService(jwtConfig config.Jwt, repo adapters.AuthRepo) adapters.AuthService {
	return &AuthService{repo: repo, config: jwtConfig}
}
