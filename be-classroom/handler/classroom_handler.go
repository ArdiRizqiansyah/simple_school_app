package handler

import (
	"be-classroom/dto"
	"be-classroom/pkg/errs"
	"be-classroom/pkg/helper"
	"be-classroom/service/classroom_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type classroomHandler struct {
	cs classroom_service.ClassroomService
}

func NewClassroomHandler(classroomService classroom_service.ClassroomService) *classroomHandler {
	return &classroomHandler{
		cs: classroomService,
	}
}

func (ch *classroomHandler) GetAllClassrooms(ctx *gin.Context) {
	page, perPage, err := helper.GetQueryPagination(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	classrooms, err := ch.cs.GetAllClassroomsWithPagination(page, perPage)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, classrooms)
}

func (ch *classroomHandler) GetClassroomById(ctx *gin.Context) {
	classroomId, errParam := helper.GetParamId(ctx, "classroomId")

	if errParam != nil {
		ctx.AbortWithStatusJSON(errParam.Status(), errParam)
		return
	}

	classroom, err := ch.cs.GetClassroomById(classroomId)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, classroom)
}

func (ch *classroomHandler) CreateClassroom(ctx *gin.Context) {
	classroomPayload := &dto.ClassroomCreateRequest{}

	if err := ctx.ShouldBindJSON(classroomPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError(err.Error())
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := ch.cs.CreateClassroom(classroomPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (ch *classroomHandler) EditClassroom(ctx *gin.Context) {
	classroomId, errParam := helper.GetParamId(ctx, "classroomId")

	if errParam != nil {
		ctx.AbortWithStatusJSON(errParam.Status(), errParam)
		return
	}

	classroomPayload := &dto.ClassroomUpdateRequest{}

	if err := ctx.ShouldBindJSON(classroomPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError(err.Error())
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := ch.cs.EditClassroom(classroomId, classroomPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (ch *classroomHandler) DeleteClassroom(ctx *gin.Context) {
	classroomId, errParam := helper.GetParamId(ctx, "classroomId")

	if errParam != nil {
		ctx.AbortWithStatusJSON(errParam.Status(), errParam)
		return
	}

	err := ch.cs.DeleteClassroom(classroomId)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, dto.WebResponse{
		StatusCode: http.StatusOK,
		Message:    "Classroom deleted successfully",
		Data:       nil,
	})
}
