package service

import (
	"context"
	"errors"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/config"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/pkg/hash"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/pkg/tokens"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type AuthService struct {
	repo   adapters.AuthRepo
	config config.Jwt
}

var (
	ErrUserNotFound             = errors.New("user not found")
	ErrUserExists               = errors.New("user exists")
	ErrDatabaseWrite            = errors.New("io error while creating user. see server logs")
	ErrValidationError          = errors.New("validation error")
	ErrRequestValidation        = errors.New("validation  error")
	ErrCacheSave                = errors.New("error occured while saving cache")
	ErrCacheFetch               = errors.New("error occured while fetching cache")
	ErrCacheDelete              = errors.New("error occured while deleting cache")
	ErrPasswordNotMatch         = errors.New("passwords dont match")
	ErrIncorrectPassword        = errors.New("incorrect password/username")
	ErrIncorrectOTP             = errors.New("incorrect otp")
	ErrTokenGeneration          = errors.New("io error while generating token")
	ErrHashGeneration           = errors.New("io error in generating password hash")
	ErrOTPNotInitialied         = errors.New("otp tracking uuid missing from cache")
	ErrTokenBlacklisted         = errors.New("token expired/invalid")
	ErrTokenInvalid             = errors.New("token invalid or expired")
	ErrVerificationOnWrongPhone = errors.New("verification on wrong phone number")
	ErrUserLoggedOut            = errors.New("user logged out")
)

func (d AuthService) SendResetOTP(request dto.OtpGenReq) (*dto.OtpGenRes, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// check if user with phone number exists
	user, err := d.repo.GetUserByPhone(ctx, request.Phone)
	if user == nil || err != nil {
		return nil, ErrUserNotFound
	}

	// if user exists, send otp
	otpRes, err := d.repo.CreateOtpCode(ctx, request)
	if err != nil {
		return nil, err
	}

	if otpRes.StatusCode != http.StatusOK {
		return nil, ErrHashGeneration
	}
	// save trackerID
	err = d.repo.SavePhoneFromResetOTP(ctx, otpRes.TrackingUuid, request.Phone)
	if err != nil {
		log.Errorf("error while caching %v", err)
		return nil, ErrCacheSave
	}

	return otpRes, nil
}

func (d AuthService) CreateUser(registerRequest dto.RegisterReq) (*dto.UserRegistrationRes, error) {

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
		return nil, ErrHashGeneration
	}

	registerRequest.Password = hashedPassword

	// create user
	createdUser, err := d.repo.CreateUser(ctx, registerRequest)
	if err != nil {
		log.Errorf("user creation error: %s", err)
		return nil, ErrDatabaseWrite
	}

	// generate jwt
	jwtToken, refreshToken, err := tokens.GenerateToken(createdUser.UserId, d.config.ExpiryMinutes, d.config.RefreshExpiryDays)
	if err != nil {
		log.Errorf("jwt generation error: %s", err)
		return nil, ErrTokenGeneration
	}

	res := dto.UserRegistrationRes{
		User:         *createdUser,
		Token:        jwtToken,
		RefreshToken: refreshToken,
	}

	return &res, nil
}

func (d AuthService) VerifyResetOTP(request dto.OtpVerificationReq) (*dto.OtpVerificationRes, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Save some IO time by checking redis
	_, err := d.repo.GetPhoneFromResetOTP(ctx, request.TrackingUID)
	if err != nil {
		if err == redis.Nil {
			return nil, ErrOTPNotInitialied
		}

		log.Errorf("error fetching from cache: %v", err)
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
	defer cancel()

	if request.ConfirmPassword != request.Password {
		return nil, ErrPasswordNotMatch
	}
	phone, err := d.repo.GetPhoneFromResetOTP(ctx, request.TrackerUUID)
	if err != nil {
		if err == redis.Nil {
			return nil, ErrOTPNotInitialied
		}

		log.Errorf("error fetching from cache: %v", err)
		return nil, ErrCacheFetch
	}

	user, err := d.repo.GetUserByPhone(ctx, phone)
	if err != nil {
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

	// delete otp tracker
	err = d.repo.InvalidateResetTracker(ctx, request.TrackerUUID)
	if err != nil {
		log.Errorf("error deleting from cache: %v", err)
		return nil, ErrCacheDelete
	}

	return &dto.ResetRes{}, nil
}

func (d AuthService) SendLoginOtp(request dto.LoginInitReq) (*dto.LoginInitRes, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	otpReq := dto.OtpGenReq{Phone: request.PhoneNumber}

	// check if user with phone number exists
	user, err := d.repo.GetUserByPhone(ctx, request.PhoneNumber)
	if user == nil || err != nil {
		return nil, ErrUserNotFound
	}

	// verify login password
	if !hash.CompareHash(request.Password, user.Hash) {
		return nil, ErrIncorrectPassword
	}

	// if user exists, send otp
	otpRes, err := d.repo.CreateOtpCode(ctx, otpReq)
	if err != nil {
		return nil, err
	}

	if otpRes.StatusCode != http.StatusOK {
		return nil, ErrHashGeneration
	}

	// save trackerID
	err = d.repo.SavePhoneFromLoginOTP(ctx, otpRes.TrackingUuid, request.PhoneNumber)
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
		if err == redis.Nil {
			return nil, ErrOTPNotInitialied
		}

		log.Errorf("error fetching from cache: %v", err)
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
	authToken, refreshToken, err := tokens.GenerateToken(user.UserId, d.config.ExpiryMinutes, d.config.RefreshExpiryDays)
	if err != nil {
		log.Errorf("jwt generation error: %s", err)
		return nil, ErrTokenGeneration
	}

	// delete otp tracker
	err = d.repo.InvalidateLoginTracker(ctx, request.TrackingUID)
	if err != nil {
		log.Errorf("error deleting from cache: %v", err)
		return nil, ErrCacheDelete
	}

	// delete user frp, blacklist
	err = d.repo.UnBlacklistToken(ctx, user.UserId)
	if err != nil {
		log.Errorf("error deleting from cache: %v", err)
		return nil, ErrCacheDelete
	}

	// return response
	loginRes := &dto.LoginRes{
		StatusCode:   otpVerificationRes.StatusCode,
		Message:      otpVerificationRes.Message,
		Token:        authToken,
		User:         user,
		RefreshToken: refreshToken,
	}
	return loginRes, nil
}

func (d AuthService) ResendLoginOTP(request dto.ResendOTPReq) (*dto.ResendOTPRes, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get phone from DB
	phone, err := d.repo.GetPhoneFromLoginOTP(ctx, request.TrackingUID)
	if err != nil {
		if err == redis.Nil {
			return nil, ErrOTPNotInitialied
		}

		log.Errorf("error fetching from cache: %v", err)
		return nil, ErrCacheFetch
	}

	resendOTPRes, err := d.repo.ResendOtpCode(ctx, request)
	if err != nil {
		return nil, ErrHashGeneration
	}

	if resendOTPRes.StatusCode != http.StatusOK {
		return nil, ErrHashGeneration
	}

	// cache tracking uuid
	err = d.repo.SavePhoneFromLoginOTP(ctx, resendOTPRes.TrackingUuid, phone)
	if err != nil {
		return nil, ErrCacheSave
	}

	// invalidate previous otp
	err = d.repo.InvalidateLoginTracker(ctx, request.TrackingUID)
	if err != nil {
		log.Errorf("error deleting from cache: %v", err)
		return nil, ErrCacheDelete
	}

	return resendOTPRes, nil
}

func (d AuthService) SendVerifyPhoneOTP(request dto.AccountVerificationOTPGenReq) (*dto.OtpGenRes, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	otpReq := dto.OtpGenReq{Phone: request.Phone}

	// check if user with phone number exists
	user, err := d.repo.GetUserByPhone(ctx, request.Phone)
	if user == nil || err != nil {
		log.Error(err)
		return nil, ErrUserNotFound
	}

	if user.UserId != request.UserUUID {
		return nil, ErrVerificationOnWrongPhone
	}

	// if user exists, send otp
	otpRes, err := d.repo.CreateOtpCode(ctx, otpReq)
	if err != nil {
		log.Errorf("otp creation error: %v", err)
		return nil, ErrHashGeneration
	}

	if otpRes.StatusCode != http.StatusOK {
		return nil, ErrHashGeneration
	}

	// save trackerID
	err = d.repo.SavePhoneFromVerificationOTP(ctx, otpRes.TrackingUuid, request.Phone)
	if err != nil {
		log.Errorf("error while caching %v", err)
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
		if err == redis.Nil {
			return nil, ErrOTPNotInitialied
		}
		log.Errorf("error fetching from cache: %v", err)
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
		log.Errorf("error while verifying: %v", err)
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

	// invalidate cache
	err = d.repo.InvalidateVerificationTracker(ctx, verificationRequest.TrackingUID)
	if err != nil {
		log.Errorf("error while invalidating cache: %v", err)
		return nil, ErrCacheSave
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
		if err == redis.Nil {
			return nil, ErrOTPNotInitialied
		}
		log.Errorf("error fetching from cache: %v", err)
		return nil, ErrCacheFetch
	}

	resendOTPRes, err := d.repo.ResendOtpCode(ctx, request)
	if err != nil {
		return nil, ErrCacheFetch
	}

	if resendOTPRes.StatusCode != http.StatusOK {
		return nil, ErrHashGeneration
	}

	// cache tracking uuid
	err = d.repo.SavePhoneFromVerificationOTP(ctx, resendOTPRes.TrackingUuid, phone)
	if err != nil {
		log.Errorf("error while caching %v", err)
		return nil, ErrCacheSave
	}

	// invalidate previous otp
	err = d.repo.InvalidateVerificationTracker(ctx, request.TrackingUID)
	if err != nil {
		log.Errorf("error deleting from cache: %v", err)
		return nil, ErrCacheDelete
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
		if err == redis.Nil {
			return nil, ErrOTPNotInitialied
		}
		log.Errorf("error fetching from cache: %v", err)
		return nil, ErrCacheFetch
	}

	resendOTPRes, err := d.repo.ResendOtpCode(ctx, request)
	if err != nil {
		return nil, ErrCacheFetch
	}

	if resendOTPRes.StatusCode != http.StatusOK {
		return nil, ErrHashGeneration
	}

	// cache tracking uuid
	err = d.repo.SavePhoneFromResetOTP(ctx, resendOTPRes.TrackingUuid, phone)
	if err != nil {
		log.Errorf("error while caching %v", err)
		return nil, ErrCacheSave
	}

	// invalidate previous otp
	err = d.repo.InvalidateResetTracker(ctx, request.TrackingUID)
	if err != nil {
		log.Errorf("error deleting from cache: %v", err)
		return nil, ErrCacheDelete
	}

	// return response
	return resendOTPRes, nil
}

func (d AuthService) RefreshToken(request dto.RefreshTokenReq) (*dto.RefreshTokenRes, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// verify token
	userUUID, err := tokens.VerifyToken(request.RefreshToken, true)
	if err != nil {
		return nil, ErrTokenInvalid
	}

	// check if user is logged out
	blacklisted, _ := d.repo.IsTokenBlacklisted(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	if blacklisted {
		return nil, ErrTokenBlacklisted
	}

	// generate new token
	token, refreshToken, err := tokens.GenerateToken(userUUID, d.config.ExpiryMinutes, d.config.RefreshExpiryDays)
	if err != nil {
		return nil, err
	}

	//return response
	return &dto.RefreshTokenRes{
		BearerToken:  token,
		RefreshToken: refreshToken,
	}, nil
}

func (d AuthService) Logout(userUUID string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// add uuid to blacklist
	err := d.repo.BlacklistToken(ctx, userUUID)
	if err != nil {
		return err
	}

	return nil
}

func (d AuthService) VerifyAccessToken(token string) (string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// verify token
	userUUID, err := tokens.VerifyToken(token, false)
	if err != nil {
		log.Errorf("error while verifying token: %v", err)
		return "", ErrTokenInvalid
	}

	// check if user is logged i
	blacklisted, err := d.repo.IsTokenBlacklisted(ctx, userUUID)
	if err != nil {
		return "", ErrUserLoggedOut
	}

	if blacklisted {
		return "", ErrTokenBlacklisted
	}

	return userUUID, nil
}

func NewDefaultAuthService(jwtConfig config.Jwt, repo adapters.AuthRepo) adapters.AuthService {
	return &AuthService{repo: repo, config: jwtConfig}
}
