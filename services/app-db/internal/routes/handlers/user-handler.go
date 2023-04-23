package handlers

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/models"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Handler) CreateUser(context context.Context, userReq *db.CreateUserReq) (*db.CreateUserRes, error) {
	if userReq == nil {
		return nil, ErrEmptyRequest
	}
	user := &models.User{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Phone:    userReq.PhoneNumber,
		Password: userReq.PasswordHash,
	}

	createdUser, err := s.userRepo.CreateUser(context, user)
	if err != nil {
		return nil, err
	}

	res := &db.CreateUserRes{
		Name:        createdUser.Name,
		Email:       createdUser.Email,
		PhoneNumber: createdUser.Phone,
		CreatedAt:   timestamppb.New(createdUser.CreatedAt),
		UpdatedAt:   timestamppb.New(createdUser.UpdatedAt),
		DeletedAt:   timestamppb.New(createdUser.DeletedAt.Time),
	}

	return res, nil
}

func (s *Handler) UpdateUser(context context.Context, userReq *db.UpdateUserReq) (*db.UpdateUserRes, error) {
	if userReq == nil {
		return nil, ErrEmptyRequest
	}
	user := &models.User{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Phone:    userReq.PhoneNumber,
		Password: userReq.PasswordHash,
	}

	createdUser, err := s.userRepo.UpdateUser(context, user)
	if err != nil {
		return nil, err
	}

	res := &db.UpdateUserRes{
		UserID:      createdUser.ID.String(),
		Name:        createdUser.Name,
		Email:       createdUser.Email,
		PhoneNumber: createdUser.Phone,
		Verified:    createdUser.Verified,
		CreatedAt:   timestamppb.New(createdUser.CreatedAt),
		UpdatedAt:   timestamppb.New(createdUser.UpdatedAt),
		DeletedAt:   timestamppb.New(createdUser.DeletedAt.Time),
	}

	return res, nil
}

func (s *Handler) GetPagedUsers(ctx context.Context, page *db.GetPagedUsersReq) (*db.GetPagedUsersRes, error) {
	if page == nil {
		return nil, ErrEmptyRequest
	}

	pageDto := &models.Paging{
		Offset: int(page.Offset),
		Limit:  int(page.Limit),
	}

	pagedUsers, err := s.userRepo.GetPagedUsers(ctx, pageDto)
	if err != nil {
		return nil, err
	}

	res := &db.GetPagedUsersRes{Offset: page.Offset, Limit: page.Limit, Users: make([]*db.User, 0)}
	for _, usr := range pagedUsers {
		res.Users = append(res.Users, &db.User{
			UserID:      usr.ID.String(),
			Verified:    usr.Verified,
			Name:        usr.Name,
			Email:       usr.Email,
			PhoneNumber: usr.Phone,
			CreatedAt:   timestamppb.New(usr.CreatedAt),
			UpdatedAt:   timestamppb.New(usr.UpdatedAt),
			DeletedAt:   timestamppb.New(usr.DeletedAt.Time),
		})
	}

	return res, nil
}

func (s *Handler) GetUserByField(ctx context.Context, filter *db.GetByfieldReq) (*db.GetByfieldRes, error) {
	if filter == nil {
		return nil, ErrEmptyRequest
	}
	var filterMap map[string]interface{} = make(map[string]interface{})

	// fill the filterMap with the filter values
	for k, f := range filter.Filter {
		filterMap[k] = f.Value
	}

	// page of a single user
	page := models.Paging{
		Offset: 0,
		Limit:  1,
	}

	pagedUsers, err := s.userRepo.GetUserByField(ctx, &filterMap, page)
	if err != nil {
		return nil, err
	}

	res := &db.GetPagedUsersRes{Offset: 0, Limit: 1, Users: make([]*db.User, 0)}
	for _, usr := range pagedUsers {
		res.Users = append(res.Users, &db.User{
			UserID:      usr.ID.String(),
			Verified:    usr.Verified,
			Name:        usr.Name,
			Email:       usr.Email,
			PhoneNumber: usr.Phone,
			CreatedAt:   timestamppb.New(usr.CreatedAt),
			UpdatedAt:   timestamppb.New(usr.UpdatedAt),
			DeletedAt:   timestamppb.New(usr.DeletedAt.Time),
		})
	}

	return nil, ErrEmptyRequest
}
