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
	"net/http"
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
	ErrCacheSave         = errors.New("error occured while saving cache")
	ErrCacheFetch        = errors.New("error occured while fetching cache")
	ErrPasswordNotMatch  = errors.New("passwords dont match")
	ErrIncorrectPassword = errors.New("incorrect password/username")
	ErrIncorrectOTP      = errors.New("incorrect otp")
	ErrUserLogin         = errors.New("system error occured while login in")
	ErrOTPGeneration     = errors.New("otp generation error")
)

func (d AuthService) SendResetOTP(request dto.OtpGenReq) (*dto.OtpGenRes, error) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	// check if user with phone number exists
	user, err := d.repo.GetUserByPhone(ctx, request.Phone)
	if user == nil || err != nil {
		return nil, ErrUserNotFound
	}

	// if user exists, send otp
	otpRes, err := d.repo.CreateOtpCode(ctx, request)
	if user != nil {
		return nil, err
	}

	if otpRes.StatusCode != http.StatusOK {
		return nil, ErrOTPGeneration
	}
	// save trackerID
	err = d.repo.SavePhoneFromResetOTP(ctx, otpRes.TrackingUuid, request.Phone)
	if err != nil {
		return nil, ErrCacheSave
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
	user, err := d.repo.GetUserByPhone(ctx, request.Phone)
	if user == nil || err != nil {
		return nil, ErrUserNotFound
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

	if otpRes.StatusCode != http.StatusOK {
		return nil, ErrOTPGeneration
	}

	// save trackerID
	err = d.repo.SavePhoneFromLoginOTP(ctx, otpRes.TrackingUuid, request.Phone)
	if err != nil {
		return nil, ErrCacheSave
	}

	// generate response
	loginRes := &dto.LoginInitRes{
		StatusCode:   otpRes.StatusCode,
		Message:      otpRes.Message,
		TrackingUuid: otpRes.TrackingUuid,
	}

	return loginRes, nil
}

func (d AuthService) VerifyLoginOtp(request dto.OtpVerificationReq) (*dto.LoginRes, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get phone from DB
	phone, err := d.repo.GetPhoneFromLoginOTP(ctx, request.TrackingUID)
	if err != nil {
		return nil, ErrCacheFetch
	}

	// check if user with phone number exists
	user, err := d.repo.GetUserByPhone(ctx, phone)
	if user == nil || err != nil {
		return nil, ErrIncorrectOTP
	}

	// compare with the otp service
	otpVerificationRes, err := d.repo.VerifyOtpCode(ctx, request)
	if err != nil {
		return nil, ErrIncorrectOTP
	}

	if otpVerificationRes.StatusCode != http.StatusOK {
		return nil, ErrIncorrectOTP
	}

	// generate jwt
	jwtToken, err := tokens.GenerateToken(user.UserId, d.config.Secret, d.config.ExpiryMinutes)
	if err != nil {
		log.Errorf("jwt generation error: %s", err)
		return nil, ErrUserLogin
	}

	// return response
	loginRes := &dto.LoginRes{
		StatusCode: otpVerificationRes.StatusCode,
		Message:    otpVerificationRes.Message,
		Token:      jwtToken,
		User:       user,
	}
	return loginRes, nil
}

func (d AuthService) ResendLoginOTP(request dto.ResendOTPReq) (*dto.ResendOTPRes, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get phone from DB
	phone, err := d.repo.GetPhoneFromLoginOTP(ctx, request.TrackingUID)
	if err != nil {
		return nil, ErrCacheFetch
	}

	resendOTPRes, err := d.repo.ResendOtpCode(ctx, request)
	if err != nil {
		return nil, ErrCacheFetch
	}

	if resendOTPRes.StatusCode != http.StatusOK {
		return nil, ErrOTPGeneration
	}

	// cache tracking uuid
	err = d.repo.SavePhoneFromLoginOTP(ctx, resendOTPRes.TrackingUuid, phone)
	if err != nil {
		return nil, ErrCacheSave
	}

	return resendOTPRes, nil
}

func (d AuthService) SendVerifyPhoneOTP(request dto.OtpGenReq) (*dto.OtpGenRes, error) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	otpReq := dto.OtpGenReq{Phone: request.Phone}

	// check if user with phone number exists
	user, err := d.repo.GetUserByPhone(ctx, request.Phone)
	if user == nil || err != nil {
		return nil, ErrUserNotFound
	}

	// if user exists, send otp
	otpRes, err := d.repo.CreateOtpCode(ctx, otpReq)
	if user != nil {
		return nil, err
	}

	if otpRes.StatusCode != http.StatusOK {
		return nil, ErrOTPGeneration
	}

	// save trackerID
	err = d.repo.SavePhoneFromVerificationOTP(ctx, otpRes.TrackingUuid, request.Phone)
	if err != nil {
		return nil, ErrCacheSave
	}

	// generate response
	return otpRes, nil
}

func (d AuthService) VerifyPhoneOTP(verificationRequest dto.OtpVerificationReq) (*dto.OtpVerificationRes, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get phone from DB
	phone, err := d.repo.GetPhoneFromVerificationOTP(ctx, verificationRequest.TrackingUID)
	if err != nil {
		return nil, ErrCacheFetch
	}

	// check if user with phone number exists
	user, err := d.repo.GetUserByPhone(ctx, phone)
	if user != nil || err != nil {
		return nil, ErrIncorrectOTP
	}

	// compare with the otp service
	otpVerificationRes, err := d.repo.VerifyOtpCode(ctx, verificationRequest)
	if err != nil {
		return nil, ErrIncorrectOTP
	}

	if otpVerificationRes.StatusCode != http.StatusOK {
		return nil, ErrIncorrectOTP
	}

	// update user as verified
	user.Verified = true

	_, err = d.repo.UpdateUser(ctx, *user)
	if err != nil {
		return nil, ErrIncorrectOTP
	}

	// return response
	return otpVerificationRes, nil
}

func (d AuthService) ResendVerifyPhoneOTP(request dto.ResendOTPReq) (*dto.ResendOTPRes, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get phone from DB
	phone, err := d.repo.GetPhoneFromVerificationOTP(ctx, request.TrackingUID)
	if err != nil {
		return nil, ErrCacheFetch
	}

	resendOTPRes, err := d.repo.ResendOtpCode(ctx, request)
	if err != nil {
		return nil, ErrCacheFetch
	}

	if resendOTPRes.StatusCode != http.StatusOK {
		return nil, ErrOTPGeneration
	}

	// cache tracking uuid
	err = d.repo.SavePhoneFromVerificationOTP(ctx, resendOTPRes.TrackingUuid, phone)
	if err != nil {
		return nil, ErrCacheSave
	}

	// return response
	return resendOTPRes, nil
}

func (d AuthService) ResendResetOTP(request dto.ResendOTPReq) (*dto.ResendOTPRes, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get phone from DB
	phone, err := d.repo.GetPhoneFromResetOTP(ctx, request.TrackingUID)
	if err != nil {
		return nil, ErrCacheFetch
	}

	resendOTPRes, err := d.repo.ResendOtpCode(ctx, request)
	if err != nil {
		return nil, ErrCacheFetch
	}

	if resendOTPRes.StatusCode != http.StatusOK {
		return nil, ErrOTPGeneration
	}

	// cache tracking uuid
	err = d.repo.SavePhoneFromResetOTP(ctx, resendOTPRes.TrackingUuid, phone)
	if err != nil {
		return nil, ErrCacheSave
	}

	// return response
	return resendOTPRes, nil
}

func NewDefaultAuthService(jwtConfig config.Jwt, repo adapters.AuthRepo) adapters.AuthService {
	return &AuthService{repo: repo, config: jwtConfig}
}
