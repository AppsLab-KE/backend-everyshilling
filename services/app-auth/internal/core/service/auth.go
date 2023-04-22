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
	ErrCache             = errors.New("error occured while saving cache")
	ErrCacheFetch        = errors.New("error occured while fetching cache")
	ErrPasswordNotMatch  = errors.New("passwords dont match")
	ErrIncorrectPassword = errors.New("incorrect password/username")
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

	// save trackerID
	err = d.repo.SavePhoneFromResetOTP(ctx, otpRes.TrackingUuid, request.Phone)
	if err != nil {
		return nil, ErrCache
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
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	// Save some IO time by checking redis
	_, err := d.repo.GetPhoneFromResetOTP(ctx, request.TrackingUID)
	if err != nil {
		return nil, ErrCacheFetch
	}

	verificationRes, err := d.repo.VerifyOtpCode(ctx, request)
	if err != nil {
		return nil, err
	}

	return verificationRes, nil
}

func (d AuthService) ChangePassword(request dto.ResetReq) (*dto.ResetRes, error) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if request.ConfirmPassword != request.Password {
		return nil, ErrPasswordNotMatch
	}
	phone, err := d.repo.GetPhoneFromResetOTP(ctx, request.TrackerUUID)
	if err != nil {
		return nil, ErrCacheFetch
	}

	user, err := d.repo.GetUserByPhone(ctx, phone)
	if user != nil {
		return nil, ErrUserExists
	}

	// create password hash
	hashedPassword, err := hash.GenerateHash(request.Password)
	if err != nil {
		log.Errorf("password hashing error: %s", err)
		return nil, err
	}

	user.Hash = hashedPassword

	// update user
	_, err = d.repo.UpdateUser(ctx, *user)
	if err != nil {
		return nil, err
	}

	return &dto.ResetRes{}, nil
}

func (d AuthService) SendLoginOtp(request dto.LoginInitReq) (*dto.LoginInitRes, error) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	otpReq := dto.OtpGenReq{Phone: request.Phone}

	// check if user with phone number exists
	user, _ := d.repo.GetUserByPhone(ctx, request.Phone)
	if user != nil {
		return nil, ErrIncorrectPassword
	}

	// verify login password
	if !hash.CompareHash(request.Password, user.Hash) {
		return nil, ErrIncorrectPassword
	}

	// if user exists, send otp
	otpRes, err := d.repo.CreateOtpCode(ctx, otpReq)
	if user != nil {
		return nil, err
	}

	// save trackerID
	err = d.repo.SavePhoneFromLoginOTP(ctx, otpRes.TrackingUuid, request.Phone)
	if err != nil {
		return nil, ErrCache
	}

	// generate response
	loginRes := &dto.LoginInitRes{
		StatusCode:   otpRes.StatusCode,
		Message:      otpRes.Message,
		TrackingUuid: otpRes.TrackingUuid,
	}

	return loginRes, nil
}

func (d AuthService) VerifyLoginOtp(request dto.OtpVerificationReq) (*dto.OtpVerificationRes, error) {

}

func (d AuthService) ResendLoginOTP(request dto.ResendLoginOTPReq) (*dto.ResendLoginOTPRes, error) {

}

func (d AuthService) SendVerifyPhoneOTP(verificationRequest dto.OtpVerificationReq) (*dto.OtpVerificationRes, error) {

}

func (d AuthService) VerifyPhoneOTP(verificationRequest dto.OtpVerificationReq) (*dto.OtpVerificationRes, error) {

}

func (d AuthService) ResendVerifyPhoneOTP(request dto.ResendLoginOTPReq) (*dto.ResendLoginOTPRes, error) {

}

func NewDefaultAuthService(jwtConfig config.Jwt, repo adapters.AuthRepo) adapters.AuthService {
	return &AuthService{repo: repo, config: jwtConfig}
}
