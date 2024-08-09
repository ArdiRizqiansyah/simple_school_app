package user_repository

import (
	"be-classroom/dto"
	"be-classroom/entity"
	"be-classroom/pkg/errs"
)

type UserRepository interface {
	FetchByEmail(email string) (*entity.User, errs.Error)
	FetchById(userId int) (*entity.User, errs.Error)
	Store(userPayload *entity.User) (*dto.UserResponse, errs.Error)
}
