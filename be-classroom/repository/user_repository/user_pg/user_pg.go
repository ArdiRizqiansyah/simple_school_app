package user_pg

import (
	"be-classroom/dto"
	"be-classroom/entity"
	"be-classroom/pkg/errs"
	"be-classroom/repository/user_repository"
	"database/sql"
)

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) user_repository.UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (userRepo *userRepositoryImpl) FetchByEmail(email string) (*entity.User, errs.Error) {

	user := entity.User{}
	err := userRepo.db.QueryRow(fetchByEmailQuery, email).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}

func (ur *userRepositoryImpl) FetchById(userId int) (*entity.User, errs.Error) {
	user := entity.User{}

	err := ur.db.QueryRow(fetchByIdQuery, userId).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &user, nil
}

// Create ...
func (ur *userRepositoryImpl) Store(userPayload *entity.User) (*dto.UserResponse, errs.Error) {
	var user dto.UserResponse

	err := ur.db.QueryRow(
		createUserQuery,
		userPayload.Name,
		userPayload.Email,
		userPayload.Password,
	).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &user, nil
}
