package handler

import (
	"be-classroom/dto"
	"be-classroom/pkg/errs"
	"be-classroom/pkg/helper"
	"be-classroom/service/classroom_student_service"
	"be-classroom/service/student_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type classroomStudentHandler struct {
	css classroom_student_service.ClassroomStudentService
	ss  student_service.StudentService
}

func NewClassroomStudentHandler(classroomStudentService classroom_student_service.ClassroomStudentService, studentService student_service.StudentService) *classroomStudentHandler {
	return &classroomStudentHandler{
		css: classroomStudentService,
		ss:  studentService,
	}
}

func (csh *classroomStudentHandler) GetAllClassroomStudents(ctx *gin.Context) {
	classroomId, errParam := helper.GetParamId(ctx, "classroomId")

	if errParam != nil {
		ctx.AbortWithStatusJSON(errParam.Status(), errParam)
		return
	}

	// Get all classroom students
	classroomStudents, err := csh.css.GetAllClassroomStudents(classroomId)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	// Get students that don't have in classroom
	students, err := csh.ss.GetStudentDontHaveInClassroom(classroomId)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, dto.WebResponse{
		StatusCode: http.StatusOK,
		Message:    "Classroom students data",
		Data: map[string]interface{}{
			"students":           students,
			"classroom_students": classroomStudents,
		},
	})
}

func (csh *classroomStudentHandler) AddStudentToClassroom(ctx *gin.Context) {
	classroomId, errParam := helper.GetParamId(ctx, "classroomId")

	if errParam != nil {
		ctx.AbortWithStatusJSON(errParam.Status(), errParam)
		return
	}

	classroomStudentPayload := &dto.ClassroomStudentCreateRequest{}

	if err := ctx.ShouldBindJSON(classroomStudentPayload); err != nil {
		errBindJson := errs.NewBadRequestError("Invalid JSON body")
		ctx.AbortWithStatusJSON(errBindJson.Status(), gin.H{"error": errBindJson.Message()})
		return
	}

	err := csh.css.CreateClassroomStudent(classroomId, classroomStudentPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.WebResponse{
		StatusCode: http.StatusCreated,
		Message:    "Classroom student added successfully",
		Data:       nil,
	})
}

func (csh *classroomStudentHandler) DeleteClassroomStudent(ctx *gin.Context) {
	classroomStudentId, errParam := helper.GetParamId(ctx, "classroomStudentId")

	if errParam != nil {
		ctx.AbortWithStatusJSON(errParam.Status(), errParam)
		return
	}

	err := csh.css.DeleteClassroomStudent(classroomStudentId)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, dto.WebResponse{
		StatusCode: http.StatusOK,
		Message:    "Classroom student deleted successfully",
		Data:       nil,
	})
}
