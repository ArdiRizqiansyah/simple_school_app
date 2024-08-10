package classroom_student_service

import (
	"be-classroom/dto"
	"be-classroom/entity"
	"be-classroom/pkg/errs"
	"be-classroom/repository/classroom_repository"
	"be-classroom/repository/classroom_student_repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Delete_Success(t *testing.T) {
	repo := classroom_student_repository.NewClassroomStudentRepositoryMock()
	repoClassroom := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomStudentService(repo, repoClassroom)

	classroom_student_repository.DestroyClassroomStudent = func(classroomStudentId int) errs.Error {
		return nil
	}

	err := service.DeleteClassroomStudent(1)

	assert.Nil(t, err)
}

func Test_Delete_Failed(t *testing.T) {
	repo := classroom_student_repository.NewClassroomStudentRepositoryMock()
	repoClassroom := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomStudentService(repo, repoClassroom)

	classroom_student_repository.DestroyClassroomStudent = func(classroomStudentId int) errs.Error {
		return errs.NewNotFoundError("classroom student not found")
	}

	err := service.DeleteClassroomStudent(1)

	assert.NotNil(t, err)
}

func Test_Create_Success(t *testing.T) {
	repo := classroom_student_repository.NewClassroomStudentRepositoryMock()
	repoClassroom := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomStudentService(repo, repoClassroom)

	classroom_repository.FindClassroomById = func(classroomId int) (*dto.ClassroomResponse, errs.Error) {
		return &dto.ClassroomResponse{
			Id: 1,
		}, nil
	}

	classroom_student_repository.StoreClassroomStudent = func(classroomId int, studentId []int) errs.Error {
		return nil
	}

	payload := &dto.ClassroomStudentCreateRequest{
		StudentId: []int{1},
	}

	err := service.CreateClassroomStudent(1, payload)

	assert.Nil(t, err)
}

func Test_Create_FailGetClassroom(t *testing.T) {
	repo := classroom_student_repository.NewClassroomStudentRepositoryMock()
	repoClassroom := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomStudentService(repo, repoClassroom)

	classroom_repository.FindClassroomById = func(classroomId int) (*dto.ClassroomResponse, errs.Error) {
		return nil, errs.NewNotFoundError("classroom not found")
	}

	payload := &dto.ClassroomStudentCreateRequest{
		StudentId: []int{1},
	}

	err := service.CreateClassroomStudent(1, payload)

	assert.NotNil(t, err)
}

func Test_Create_FailInvalidData(t *testing.T) {
	repo := classroom_student_repository.NewClassroomStudentRepositoryMock()
	repoClassroom := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomStudentService(repo, repoClassroom)

	classroom_repository.FindClassroomById = func(classroomId int) (*dto.ClassroomResponse, errs.Error) {
		return &dto.ClassroomResponse{
			Id: 1,
		}, nil
	}

	payload := &dto.ClassroomStudentCreateRequest{
		StudentId: []int{},
	}

	err := service.CreateClassroomStudent(1, payload)

	assert.NotNil(t, err)
}

func Test_Create_FailStoreData(t *testing.T) {
	repo := classroom_student_repository.NewClassroomStudentRepositoryMock()
	repoClassroom := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomStudentService(repo, repoClassroom)

	classroom_repository.FindClassroomById = func(classroomId int) (*dto.ClassroomResponse, errs.Error) {
		return &dto.ClassroomResponse{
			Id: 1,
		}, nil
	}

	classroom_student_repository.StoreClassroomStudent = func(classroomId int, studentId []int) errs.Error {
		return errs.NewInternalServerError("error store classroom student")
	}

	payload := &dto.ClassroomStudentCreateRequest{
		StudentId: []int{1},
	}

	err := service.CreateClassroomStudent(1, payload)

	assert.NotNil(t, err)
}

func Test_GetAll_Success(t *testing.T) {
	repo := classroom_student_repository.NewClassroomStudentRepositoryMock()
	repoClassroom := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomStudentService(repo, repoClassroom)

	classroom_student_repository.FindAllClassroomStudents = func(classroomId int) ([]classroom_student_repository.ClassroomStudentWithStudentMapped, errs.Error) {
		return []classroom_student_repository.ClassroomStudentWithStudentMapped{
			{
				Id:          1,
				ClassroomId: 1,
				StudentId:   1,
				CreatedAt:   "2021-01-01",
				UpdatedAt:   "2021-01-01",
				Student: entity.Student{
					Id:         1,
					Name:       "Student 1",
					Nis:        "123",
					PlaceBirth: "Jakarta",
					DateBirth:  "2021-01-01",
					CreatedAt:  "2021-01-01",
					UpdatedAt:  "2021-01-01",
				},
			},
		}, nil
	}

	result, err := service.GetAllClassroomStudents(1)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result))
}

func Test_GetAll_Failed(t *testing.T) {
	repo := classroom_student_repository.NewClassroomStudentRepositoryMock()
	repoClassroom := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomStudentService(repo, repoClassroom)

	classroom_student_repository.FindAllClassroomStudents = func(classroomId int) ([]classroom_student_repository.ClassroomStudentWithStudentMapped, errs.Error) {
		return nil, errs.NewNotFoundError("classroom student not found")
	}

	result, err := service.GetAllClassroomStudents(1)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}
