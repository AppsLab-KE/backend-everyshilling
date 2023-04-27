package repository

import (
	"context"
	"errors"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/entity"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
	"github.com/AppsLab-KE/be-go-gen-grpc/otp"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type authRepo struct {
	dbStorage    adapters.DBStorage
	otpStorage   adapters.OTPStorage
	cacheStorage adapters.CacheStorage
}

func (a authRepo) InvalidateSession() error {
	return nil
}

func (a authRepo) ResendOtpCode(ctx context.Context, data dto.ResendOTPReq) (*dto.ResendOTPRes, error) {
	resendOtpReq := &otp.ResendOTPReq{
		TrackingId: data.TrackingUID,
	}

	createOtpRes, err := a.otpStorage.ResendOTP(ctx, resendOtpReq)
	if err != nil {
		return nil, err
	}

	otpRes := &dto.ResendOTPRes{
		StatusCode:   createOtpRes.StatusCode,
		Message:      createOtpRes.Message,
		TrackingUuid: createOtpRes.TrackingUuid,
	}

	return otpRes, nil
}

func (a authRepo) CreateOtpCode(ctx context.Context, data dto.OtpGenReq) (*dto.OtpGenRes, error) {
	otpReq := &otp.CreateAndSendOtpReq{
		PhoneNumber: data.Phone,
	}
	createOtpRes, err := a.otpStorage.CreateAndSendOtp(ctx, otpReq)
	if err != nil {
		return nil, err
	}

	otpRes := &dto.OtpGenRes{
		StatusCode:   createOtpRes.StatusCode,
		Message:      createOtpRes.Message,
		TrackingUuid: createOtpRes.TrackingUuid,
	}

	return otpRes, nil
}

func (a authRepo) VerifyOtpCode(ctx context.Context, data dto.OtpVerificationReq) (*dto.OtpVerificationRes, error) {
	otpReq := &otp.VerifyOTPReq{
		OtpCode:      data.OtpCode,
		TrackingUuid: data.TrackingUID,
	}
	createOtpRes, err := a.otpStorage.VerifyOtp(ctx, otpReq)
	if err != nil {
		return nil, err
	}

	otpRes := &dto.OtpVerificationRes{
		StatusCode: createOtpRes.StatusCode,
		Message:    createOtpRes.Message,
	}

	return otpRes, nil
}

func (a authRepo) CreateUser(ctx context.Context, registerRequest dto.RegisterReq) (*entity.User, error) {
	userReq := db.CreateUserReq{
		Name:         registerRequest.Name,
		Email:        registerRequest.Email,
		PhoneNumber:  registerRequest.PhoneNumber,
		PasswordHash: registerRequest.Password,
	}
	userRes, err := a.dbStorage.CreateUser(ctx, &userReq)
	if err != nil {
		return nil, err
	}

	user := entity.User{
		UserId:      userRes.UserID,
		Name:        userRes.Name,
		Email:       userRes.Email,
		PhoneNumber: userRes.PhoneNumber,
	}
	return &user, nil
}

func (a authRepo) GetUserByPhone(ctx context.Context, phone string) (*entity.User, error) {
	phoneFilter := anypb.Any{
		Value: []byte(phone),
	}
	keyValueReq := &db.GetByfieldReq{
		Filter: map[string]*anypb.Any{
			"phone": &phoneFilter,
		},
		Offset: 0,
		Limit:  1,
	}
	userRes, err := a.dbStorage.GetUserByField(ctx, keyValueReq)
	if err != nil {
		return nil, err
	}

	if userRes == nil {
		return nil, ErrUserNotFound
	}

	if len(userRes.Users) > 0 {
		user := &entity.User{
			UserId:      userRes.Users[0].UserID,
			Name:        userRes.Users[0].Name,
			Email:       userRes.Users[0].Email,
			PhoneNumber: userRes.Users[0].PhoneNumber,
			Verified:    userRes.Users[0].Verified,
			Hash:        userRes.Users[0].Hash,
		}

		return user, nil
	}

	return nil, ErrUserNotFound
}

func (a authRepo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	emailFilter := anypb.Any{
		Value: []byte(email),
	}
	keyValueReq := &db.GetByfieldReq{
		Filter: map[string]*anypb.Any{
			"email": &emailFilter,
		},
		Offset: 0,
		Limit:  1,
	}
	userRes, err := a.dbStorage.GetUserByField(ctx, keyValueReq)
	if err != nil {
		return nil, err
	}

	if userRes == nil {
		return nil, ErrUserNotFound
	}

	if len(userRes.Users) > 0 {
		user := &entity.User{
			UserId:      userRes.Users[0].UserID,
			Name:        userRes.Users[0].Name,
			Email:       userRes.Users[0].Email,
			PhoneNumber: userRes.Users[0].PhoneNumber,
			Verified:    userRes.Users[0].Verified,
			Hash:        userRes.Users[0].Hash,
		}

		return user, nil
	}

	return nil, ErrUserNotFound
}

func (a authRepo) UpdateUser(ctx context.Context, user entity.User) (*entity.User, error) {
	updateUserReq := &db.UpdateUserReq{
		Name:         user.Name,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		PasswordHash: user.Hash,
		UserID:       user.UserId,
		Verified:     user.Verified,
	}

	userRes, err := a.dbStorage.UpdateUser(ctx, updateUserReq)
	if err != nil {
		return nil, err
	}

	updatedUser := entity.User{
		UserId:      userRes.UserID,
		Name:        userRes.Name,
		Email:       userRes.Email,
		PhoneNumber: userRes.PhoneNumber,
		Verified:    userRes.Verified,
	}

	return &updatedUser, nil
}

func (a authRepo) SavePhoneFromLoginOTP(ctx context.Context, trackerUUID, phone string) error {
	return a.cacheStorage.SavePhoneFromLoginOTP(ctx, trackerUUID, phone)
}

func (a authRepo) GetPhoneFromLoginOTP(ctx context.Context, trackerUUID string) (string, error) {
	return a.cacheStorage.GetPhoneFromLoginOTP(ctx, trackerUUID)
}

func (a authRepo) SavePhoneFromResetOTP(ctx context.Context, trackerUUID, phone string) error {
	return a.cacheStorage.SavePhoneFromResetOTP(ctx, trackerUUID, phone)
}

func (a authRepo) GetPhoneFromResetOTP(ctx context.Context, trackerUUID string) (string, error) {
	return a.cacheStorage.GetPhoneFromResetOTP(ctx, trackerUUID)
}

func (a authRepo) SavePhoneFromVerificationOTP(ctx context.Context, trackerUUID, phone string) error {
	return a.cacheStorage.SavePhoneFromVerificationOTP(ctx, trackerUUID, phone)
}

func (a authRepo) GetPhoneFromVerificationOTP(ctx context.Context, trackerUUID string) (string, error) {
	return a.cacheStorage.GetPhoneFromVerificationOTP(ctx, trackerUUID)
}

func (a authRepo) InvalidateLoginTracker(ctx context.Context, trackerUUID string) error {
	return a.cacheStorage.InvalidateLoginTracker(ctx, trackerUUID)
}

func (a authRepo) InvalidateResetTracker(ctx context.Context, trackerUID string) error {
	return a.cacheStorage.InvalidateResetTracker(ctx, trackerUID)
}

func (a authRepo) InvalidateVerificationTracker(ctx context.Context, trackerUUID string) error {
	return a.cacheStorage.InvalidateVerificationTracker(ctx, trackerUUID)
}

func (a authRepo) BlacklistToken(ctx context.Context, userUUID string) error {
	return a.cacheStorage.BlacklistToken(ctx, userUUID)
}

func (a authRepo) IsTokenBlacklisted(ctx context.Context, userUUID string) (bool, error) {
	return a.cacheStorage.IsTokenBlacklisted(ctx, userUUID)
}

func (a authRepo) UnBlacklistToken(ctx context.Context, userUUID string) error {
	return a.cacheStorage.UnBlacklistToken(ctx, userUUID)
}

func NewAuthRepo(cacheStorage adapters.CacheStorage, dbStorage adapters.DBStorage, otpStorage adapters.OTPStorage) adapters.AuthRepo {
	return &authRepo{
		dbStorage:    dbStorage,
		otpStorage:   otpStorage,
		cacheStorage: cacheStorage,
	}
}
