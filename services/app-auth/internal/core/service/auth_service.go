package service

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/config"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/errors"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/pkg/hash"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/pkg/tokens"
	log "github.com/sirupsen/logrus"
)

type AuthService struct {
	repo   adapters.AuthRepo
	config config.Config
}

func (d AuthService) CreateUser(registerRequest dto.RegisterRequest) (*dto.UserRegistrationRes, error) {
	// check if registerRequest exists by phone
	user, _ := d.repo.GetUserByEmail(registerRequest.Email)
	if user != nil {
		return nil, errors.ErrUserExists
	}

	user, _ = d.repo.GetUserByPhone(registerRequest.PhoneNumber)
	if user != nil {
		return nil, errors.ErrUserExists
	}

	// create password hash
	hashedPassword, err := hash.GenerateHash(registerRequest.Password)
	if err != nil {
		log.Errorf("password hashing error: %s", err)
		return nil, errors.ErrUserCreation
	}

	registerRequest.Password = hashedPassword

	// create user
	createdUser, err := d.repo.CreateUser(registerRequest)
	if err != nil {
		log.Errorf("user creation error: %s", err)
		return nil, errors.ErrUserCreation
	}

	// generate jwt
	jwtToken, err := tokens.GenerateToken(createdUser.UserId, d.config.Jwt.Secret, d.config.Jwt.ExpiryMinutes)
	if err != nil {
		log.Errorf("jwt generation error: %s", err)
		return nil, errors.ErrUserCreation
	}

	res := dto.UserRegistrationRes{
		User:  *createdUser,
		Token: jwtToken,
	}

	return &res, nil
}

func (d AuthService) RequestOtp(request dto.OtpReq) dto.DefaultRes[interface{}] {
	//TODO implement me
	panic("implement me")
}

func NewDefaultAuthService(repo adapters.AuthRepo) adapters.AuthService {
	return &AuthService{repo: repo}
}
