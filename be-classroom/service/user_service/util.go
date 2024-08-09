package user_service

import (
	"be-classroom/pkg/errs"
	"net/http"
)

func (u *userServiceImpl) validateDuplicateEmail(email string) errs.Error {
	usr, err := u.ur.FetchByEmail(email)

	if err != nil && err.Status() != http.StatusNotFound {
		return err
	}

	if usr != nil {
		return errs.NewConflictError("email has been used")
	}

	if usr != nil {
		return errs.NewConflictError("username has been used")
	}
	return nil
}
