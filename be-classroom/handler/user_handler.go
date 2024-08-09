package handler

import (
	"be-classroom/dto"
	"be-classroom/entity"
	"be-classroom/pkg/errs"
	"be-classroom/service/user_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	Profile(ctx *gin.Context)
}

type userHandlerImpl struct {
	us user_service.UserService
}

func NewUserHandler(userService user_service.UserService) UserHandler {
	return &userHandlerImpl{
		us: userService,
	}
}

func (u *userHandlerImpl) Login(ctx *gin.Context) {
	userPayload := &dto.UserLoginRequest{}

	if err := ctx.ShouldBindJSON(userPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := u.us.Login(userPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (u *userHandlerImpl) Register(ctx *gin.Context) {
	userPayload := &dto.UserRegisterRequest{}

	if err := ctx.ShouldBindJSON(userPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := u.us.CreateUser(userPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

func (u *userHandlerImpl) Profile(ctx *gin.Context) {
	user, ok := ctx.MustGet("userData").(entity.User)

	if !ok {
		internalServerErr := errs.NewInternalServerError("something went wrong")
		ctx.AbortWithStatusJSON(internalServerErr.Status(), internalServerErr)
		return
	}

	response, err := u.us.CheckProfile(user.Email)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
