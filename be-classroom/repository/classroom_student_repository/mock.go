package classroom_student_repository

import "be-classroom/pkg/errs"

var (
	FindAllClassroomStudents func(classroomId int) ([]ClassroomStudentWithStudentMapped, errs.Error)
	StoreClassroomStudent    func(classroomId int, studentId []int) errs.Error
	DestroyClassroomStudent  func(classroomStudentId int) errs.Error
)

type ClassroomStudentRepositoryMock struct{}

func NewClassroomStudentRepositoryMock() ClassroomStudentRepository {
	return &ClassroomStudentRepositoryMock{}
}

func (csm *ClassroomStudentRepositoryMock) FindAllClassroomStudents(classroomId int) ([]ClassroomStudentWithStudentMapped, errs.Error) {
	return FindAllClassroomStudents(classroomId)
}

func (csm *ClassroomStudentRepositoryMock) StoreClassroomStudent(classroomId int, studentId []int) errs.Error {
	return StoreClassroomStudent(classroomId, studentId)
}

func (csm *ClassroomStudentRepositoryMock) DestroyClassroomStudent(classroomStudentId int) errs.Error {
	return DestroyClassroomStudent(classroomStudentId)
}
