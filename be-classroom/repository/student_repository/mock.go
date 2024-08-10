package student_repository

import (
	"be-classroom/dto"
	"be-classroom/entity"
	"be-classroom/pkg/errs"
)

var (
	FindAllStudents                func() ([]dto.StudentResponse, errs.Error)
	FindStudentById                func(studentId int) (*dto.StudentResponse, errs.Error)
	FindStudentDontHaveInClassroom func(classroomId int) ([]dto.StudentResponse, errs.Error)
	CountStudents                  func() (int, errs.Error)
	FindAllStudentsWithPagination  func(page, perPage int) (*dto.StudentWithPagination, errs.Error)
	StoreStudent                   func(studentPayload *entity.Student) (*dto.StudentResponse, errs.Error)
	UpdateStudent                  func(studentId int, studentPayload *entity.Student) (*dto.StudentResponse, errs.Error)
	DestroyStudent                 func(studentId int) errs.Error
)

type StudentRepositoryMock struct{}

func NewStudentRepositoryMock() StudentRepository {
	return &StudentRepositoryMock{}
}

func (sm *StudentRepositoryMock) FindAllStudents() ([]dto.StudentResponse, errs.Error) {
	return FindAllStudents()
}

func (sm *StudentRepositoryMock) FindStudentById(studentId int) (*dto.StudentResponse, errs.Error) {
	return FindStudentById(studentId)
}

func (sm *StudentRepositoryMock) FindStudentDontHaveInClassroom(classroomId int) ([]dto.StudentResponse, errs.Error) {
	return FindStudentDontHaveInClassroom(classroomId)
}

func (sm *StudentRepositoryMock) CountStudents() (int, errs.Error) {
	return CountStudents()
}

func (sm *StudentRepositoryMock) FindAllStudentsWithPagination(page, perPage int) (*dto.StudentWithPagination, errs.Error) {
	return FindAllStudentsWithPagination(page, perPage)
}

func (sm *StudentRepositoryMock) StoreStudent(studentPayload *entity.Student) (*dto.StudentResponse, errs.Error) {
	return StoreStudent(studentPayload)
}

func (sm *StudentRepositoryMock) UpdateStudent(studentId int, studentPayload *entity.Student) (*dto.StudentResponse, errs.Error) {
	return UpdateStudent(studentId, studentPayload)
}

func (sm *StudentRepositoryMock) DestroyStudent(studentId int) errs.Error {
	return DestroyStudent(studentId)
}
