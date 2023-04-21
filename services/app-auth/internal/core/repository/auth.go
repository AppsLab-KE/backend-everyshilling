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
	dbStorage  adapters.DBStorage
	otpStorage adapters.OTPStorage
}

func (a authRepo) ResendOtpCode(ctx context.Context, data dto.ResendLoginOTPReq) (*dto.ResendLoginOTPRes, error) {
	resendOtpReq := &otp.ResendOTPRed{
		TrackingId: data.TrackingUID,
	}

	createOtpRes, err := a.otpStorage.ResendOTP(ctx, resendOtpReq)
	if err != nil {
		return nil, err
	}

	otpRes := &dto.ResendLoginOTPRes{
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

func (a authRepo) CreateUser(ctx context.Context, registerRequest dto.RegisterRequest) (*entity.User, error) {
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
		Limit:  0,
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
		Limit:  0,
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
	}

	return &updatedUser, nil
}

func NewAuthRepo(dbStorage adapters.DBStorage, otpStorage adapters.OTPStorage) adapters.AuthRepo {
	return &authRepo{
		dbStorage:  dbStorage,
		otpStorage: otpStorage,
	}
}