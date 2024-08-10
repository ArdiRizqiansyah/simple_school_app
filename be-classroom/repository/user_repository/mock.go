package user_repository

import (
	"be-classroom/dto"
	"be-classroom/entity"
	"be-classroom/pkg/errs"
)

var (
	FetchByEmail func(email string) (*entity.User, errs.Error)
	FetchById    func(userId int) (*entity.User, errs.Error)
	Store        func(userPayload *entity.User) (*dto.UserResponse, errs.Error)
)

type UserRepositoryMock struct{}

func NewUserRepositoryMock() UserRepository {
	return &UserRepositoryMock{}
}

func (um *UserRepositoryMock) FetchByEmail(email string) (*entity.User, errs.Error) {
	return FetchByEmail(email)
}

func (um *UserRepositoryMock) FetchById(userId int) (*entity.User, errs.Error) {
	return FetchById(userId)
}

func (um *UserRepositoryMock) Store(userPayload *entity.User) (*dto.UserResponse, errs.Error) {
	return Store(userPayload)
}
