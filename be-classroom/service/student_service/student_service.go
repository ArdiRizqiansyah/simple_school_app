package student_service

import (
	"be-classroom/dto"
	"be-classroom/entity"
	"be-classroom/pkg/errs"
	"be-classroom/pkg/helper"
	"be-classroom/repository/student_repository"
)

type StudentService interface {
	GetAllStudents() ([]dto.StudentResponse, errs.Error)
	GetStudentById(studentId int) (*dto.StudentResponse, errs.Error)
	GetStudentDontHaveInClassroom(classroomId int) ([]dto.StudentResponse, errs.Error)
	TotalStudents() (int, errs.Error)
	GetAllStudentsWithPagination(page, perPage int) (*dto.StudentWithPagination, errs.Error)
	CreateStudent(studentPayload *dto.StudentCreateRequest) (*dto.StudentResponse, errs.Error)
	UpdateStudent(studentId int, studentPayload *dto.StudentUpdateRequest) (*dto.StudentResponse, errs.Error)
	DeleteStudent(studentId int) errs.Error
}

type studentServiceImpl struct {
	sr student_repository.StudentRepository
}

func NewStudentService(studentRepo student_repository.StudentRepository) StudentService {
	return &studentServiceImpl{
		sr: studentRepo,
	}
}

func (ss *studentServiceImpl) GetAllStudents() ([]dto.StudentResponse, errs.Error) {
	students, err := ss.sr.FindAllStudents()

	if err != nil {
		return nil, err
	}

	return students, nil
}

func (ss *studentServiceImpl) GetStudentById(studentId int) (*dto.StudentResponse, errs.Error) {
	student, err := ss.sr.FindStudentById(studentId)

	if err != nil {
		return nil, err
	}

	return student, nil
}

func (ss *studentServiceImpl) GetStudentDontHaveInClassroom(classroomId int) ([]dto.StudentResponse, errs.Error) {
	students, err := ss.sr.FindStudentDontHaveInClassroom(classroomId)

	if err != nil {
		return nil, err
	}

	return students, nil
}

func (ss *studentServiceImpl) CreateStudent(studentPayload *dto.StudentCreateRequest) (*dto.StudentResponse, errs.Error) {
	err := helper.ValidateStruct(studentPayload)

	if err != nil {
		return nil, err
	}

	student := &entity.Student{
		Name:       studentPayload.Name,
		Nis:        studentPayload.Nis,
		PlaceBirth: studentPayload.PlaceBirth,
		DateBirth:  studentPayload.DateBirth,
	}

	createdStudent, err := ss.sr.StoreStudent(student)

	if err != nil {
		return nil, err
	}

	return createdStudent, nil
}

func (ss *studentServiceImpl) UpdateStudent(studentId int, studentPayload *dto.StudentUpdateRequest) (*dto.StudentResponse, errs.Error) {
	err := helper.ValidateStruct(studentPayload)

	if err != nil {
		return nil, err
	}

	student := &entity.Student{
		Name:       studentPayload.Name,
		Nis:        studentPayload.Nis,
		PlaceBirth: studentPayload.PlaceBirth,
		DateBirth:  studentPayload.DateBirth,
	}

	updatedStudent, err := ss.sr.UpdateStudent(studentId, student)

	if err != nil {
		return nil, err
	}

	return updatedStudent, nil
}

func (ss *studentServiceImpl) DeleteStudent(studentId int) errs.Error {
	err := ss.sr.DestroyStudent(studentId)

	if err != nil {
		return err
	}

	return nil
}

func (ss *studentServiceImpl) TotalStudents() (int, errs.Error) {
	total, err := ss.sr.CountStudents()

	if err != nil {
		return 0, err
	}

	return total, nil
}

func (ss *studentServiceImpl) GetAllStudentsWithPagination(page, perPage int) (*dto.StudentWithPagination, errs.Error) {
	students, err := ss.sr.FindAllStudentsWithPagination(page, perPage)

	if err != nil {
		return nil, err
	}

	return students, nil
}
