package handler

import (
	"be-classroom/dto"
	"be-classroom/pkg/errs"
	"be-classroom/pkg/helper"
	"be-classroom/service/student_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	ss student_service.StudentService
}

func NewStudentHandler(studentService student_service.StudentService) *StudentHandler {
	return &StudentHandler{
		ss: studentService,
	}
}

func (sh *StudentHandler) GetAllStudents(ctx *gin.Context) {
	page, perPage, err := helper.GetQueryPagination(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	students, err := sh.ss.GetAllStudentsWithPagination(page, perPage)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, students)
}

func (sh *StudentHandler) GetStudentById(ctx *gin.Context) {
	studentId, errParam := helper.GetParamId(ctx, "studentId")

	if errParam != nil {
		ctx.AbortWithStatusJSON(errParam.Status(), errParam)
		return
	}

	student, err := sh.ss.GetStudentById(studentId)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, student)
}

func (sh *StudentHandler) CreateStudent(ctx *gin.Context) {
	studentPayload := &dto.StudentCreateRequest{}

	if err := ctx.ShouldBindJSON(studentPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError(err.Error())
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	student, err := sh.ss.CreateStudent(studentPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusCreated, student)
}

func (sh *StudentHandler) EditStudent(ctx *gin.Context) {
	studentId, errParam := helper.GetParamId(ctx, "studentId")

	if errParam != nil {
		ctx.AbortWithStatusJSON(errParam.Status(), errParam)
		return
	}

	studentPayload := &dto.StudentUpdateRequest{}

	if err := ctx.ShouldBindJSON(studentPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError(err.Error())
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	student, err := sh.ss.UpdateStudent(studentId, studentPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, student)
}

func (sh *StudentHandler) DeleteStudent(ctx *gin.Context) {
	studentId, errParam := helper.GetParamId(ctx, "studentId")

	if errParam != nil {
		ctx.AbortWithStatusJSON(errParam.Status(), errParam)
		return
	}

	err := sh.ss.DeleteStudent(studentId)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, dto.WebResponse{
		StatusCode: http.StatusOK,
		Message:    "Student deleted successfully",
		Data:       nil,
	})
}
