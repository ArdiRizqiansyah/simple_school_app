package handler

import (
	"be-classroom/dto"
	"be-classroom/service/classroom_service"
	"be-classroom/service/student_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type dashboardHandler struct {
	cs classroom_service.ClassroomService
	ss student_service.StudentService
}

func NewDashboardHandler(classroomService classroom_service.ClassroomService, studentService student_service.StudentService) *dashboardHandler {
	return &dashboardHandler{
		cs: classroomService,
		ss: studentService,
	}
}

func (dh *dashboardHandler) Dashboard(ctx *gin.Context) {
	// count total classrooms
	totalClassrooms, err := dh.cs.TotalClassrooms()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	// count total students
	totalStudents, err := dh.ss.TotalStudents()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	// get classroom data filter by pagination 5 data
	classrooms, err := dh.cs.GetAllClassroomsWithPagination(1, 5)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	// get student data filter by pagination 5 data
	students, err := dh.ss.GetAllStudentsWithPagination(1, 5)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, dto.WebResponse{
		StatusCode: http.StatusOK,
		Message:    "Success get dashboard data",
		Data: gin.H{
			"total_classrooms": totalClassrooms,
			"total_students":   totalStudents,
			"classrooms":       classrooms.Data,
			"students":         students.Data,
		},
	})
}
