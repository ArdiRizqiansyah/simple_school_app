package auth_service

import (
	"be-classroom/entity"
	"be-classroom/pkg/errs"
	"be-classroom/repository/user_repository"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
}

type authServiceImpl struct {
	ur user_repository.UserRepository
}

func NewAuthService(userRepo user_repository.UserRepository) AuthService {
	return &authServiceImpl{
		ur: userRepo,
	}
}

func (a *authServiceImpl) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		invalidToken := errs.NewUnauthenticatedError("Invalid token")
		bearerToken := ctx.GetHeader("Authorization")

		user := entity.User{}

		err := user.ValidateToken(bearerToken)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		_, err = a.ur.FetchByEmail(user.Email)

		if err != nil {
			ctx.AbortWithStatusJSON(invalidToken.Status(), invalidToken)
			return
		}

		ctx.Set("userData", user)
		ctx.Next()
	}
}
