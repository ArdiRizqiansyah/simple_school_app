package user_service

import (
	"be-classroom/dto"
	"be-classroom/entity"
	"be-classroom/pkg/errs"
	"be-classroom/pkg/helper"
	"be-classroom/repository/user_repository"
	"net/http"
)

type UserService interface {
	Login(userPayload *dto.UserLoginRequest) (*dto.TokenResponse, errs.Error)
	CreateUser(userPayload *dto.UserRegisterRequest) (*dto.WebResponse, errs.Error)
	CheckProfile(email string) (*dto.UserResponse, errs.Error)
}

type userServiceImpl struct {
	ur user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userServiceImpl{
		ur: userRepo,
	}
}

// Login ...
func (us *userServiceImpl) Login(userPayload *dto.UserLoginRequest) (*dto.TokenResponse, errs.Error) {
	err := helper.ValidateStruct(userPayload)

	if err != nil {
		return nil, err
	}

	user, err := us.ur.FetchByEmail(userPayload.Email)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewBadRequestError("Invalid email/password")
		}

		return nil, err
	}

	isValidPassword := user.ComparePassword(userPayload.Password)

	if !isValidPassword {
		return nil, errs.NewBadRequestError("Invalid email/password")
	}

	token := user.GenerateToken()

	return &dto.TokenResponse{
		Token: token,
	}, nil
}

// CreateUser ...
func (us *userServiceImpl) CreateUser(userPayload *dto.UserRegisterRequest) (*dto.WebResponse, errs.Error) {
	err := helper.ValidateStruct(userPayload)

	if err != nil {
		return nil, err
	}

	err = us.validateDuplicateEmail(userPayload.Email)

	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Name:     userPayload.Name,
		Email:    userPayload.Email,
		Password: userPayload.Password,
	}

	err = user.HashPassword()

	if err != nil {
		return nil, err
	}

	createdUser, err := us.ur.Store(user)

	if err != nil {
		return nil, err
	}

	// get token
	token := user.GenerateToken()

	// data response to client return user and token
	data := map[string]interface{}{
		"user":  createdUser,
		"token": token,
	}

	return &dto.WebResponse{
		StatusCode: http.StatusCreated,
		Message:    "User created",
		Data:       data,
	}, nil
}

// CheckProfile ...
func (us *userServiceImpl) CheckProfile(email string) (*dto.UserResponse, errs.Error) {
	user, err := us.ur.FetchByEmail(email)

	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
