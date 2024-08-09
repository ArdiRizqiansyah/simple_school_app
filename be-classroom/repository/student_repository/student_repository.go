package student_repository

import (
	"be-classroom/dto"
	"be-classroom/entity"
	"be-classroom/pkg/errs"
)

type StudentRepository interface {
	FindAllStudents() ([]dto.StudentResponse, errs.Error)
	FindStudentById(studentId int) (*dto.StudentResponse, errs.Error)
	FindStudentDontHaveInClassroom(classroomId int) ([]dto.StudentResponse, errs.Error)
	CountStudents() (int, errs.Error)
	FindAllStudentsWithPagination(page, perPage int) (*dto.StudentWithPagination, errs.Error)
	StoreStudent(studentPayload *entity.Student) (*dto.StudentResponse, errs.Error)
	UpdateStudent(studentId int, studentPayload *entity.Student) (*dto.StudentResponse, errs.Error)
	DestroyStudent(studentId int) errs.Error
}
