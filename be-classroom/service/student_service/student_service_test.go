package student_service

import (
	"be-classroom/dto"
	"be-classroom/entity"
	"be-classroom/pkg/errs"
	"be-classroom/repository/student_repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Delete_Success(t *testing.T) {
	repo := student_repository.NewStudentRepositoryMock()

	service := NewStudentService(repo)

	student_repository.DestroyStudent = func(studentId int) errs.Error {
		return nil
	}

	err := service.DeleteStudent(1)

	assert.Nil(t, err)
}

func Test_Delete_Failed(t *testing.T) {
	repo := student_repository.NewStudentRepositoryMock()

	service := NewStudentService(repo)

	student_repository.DestroyStudent = func(studentId int) errs.Error {
		return errs.NewNotFoundError("student not found")
	}

	err := service.DeleteStudent(1)

	assert.NotNil(t, err)
}

func Test_Create_Success(t *testing.T) {
	repo := student_repository.NewStudentRepositoryMock()

	service := NewStudentService(repo)

	student_repository.StoreStudent = func(studentPayload *entity.Student) (*dto.StudentResponse, errs.Error) {
		return &dto.StudentResponse{
			Id:         1,
			Name:       "Student 1",
			Nis:        "123456",
			PlaceBirth: "Jakarta",
			DateBirth:  "2000-01-01",
			CreatedAt:  "2021-01-01",
			UpdatedAt:  "2021-01-01",
		}, nil
	}

	payload := &dto.StudentCreateRequest{
		Name:       "Student 1",
		Nis:        "123456",
		PlaceBirth: "Jakarta",
		DateBirth:  "2000-01-01",
	}

	result, err := service.CreateStudent(payload)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.Id)
	assert.Equal(t, "Student 1", result.Name)
	assert.Equal(t, "123456", result.Nis)
	assert.Equal(t, "Jakarta", result.PlaceBirth)
	assert.Equal(t, "2000-01-01", result.DateBirth)
	assert.Equal(t, "2021-01-01", result.CreatedAt)
	assert.Equal(t, "2021-01-01", result.UpdatedAt)
}

func Test_Create_Failed(t *testing.T) {
	repo := student_repository.NewStudentRepositoryMock()

	service := NewStudentService(repo)

	student_repository.StoreStudent = func(studentPayload *entity.Student) (*dto.StudentResponse, errs.Error) {
		return nil, errs.NewInternalServerError("internal server error")
	}

	payload := &dto.StudentCreateRequest{
		Name:       "Student 1",
		Nis:        "123456",
		PlaceBirth: "Jakarta",
		DateBirth:  "2000-01-01",
	}

	result, err := service.CreateStudent(payload)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_Create_InvalidValidateData(t *testing.T) {
	repo := student_repository.NewStudentRepositoryMock()

	service := NewStudentService(repo)

	payload := &dto.StudentCreateRequest{
		Name: "",
	}

	result, err := service.CreateStudent(payload)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_Update_Success(t *testing.T) {
	repo := student_repository.NewStudentRepositoryMock()

	service := NewStudentService(repo)

	student_repository.UpdateStudent = func(studentId int, studentPayload *entity.Student) (*dto.StudentResponse, errs.Error) {
		return &dto.StudentResponse{
			Id:         1,
			Name:       "Student 1",
			Nis:        "123456",
			PlaceBirth: "Jakarta",
			DateBirth:  "2000-01-01",
			CreatedAt:  "2021-01-01",
			UpdatedAt:  "2021-01-01",
		}, nil
	}

	payload := &dto.StudentUpdateRequest{
		Name:       "Student 1",
		Nis:        "123456",
		PlaceBirth: "Jakarta",
		DateBirth:  "2000-01-01",
	}

	result, err := service.UpdateStudent(1, payload)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.Id)
	assert.Equal(t, "Student 1", result.Name)
	assert.Equal(t, "123456", result.Nis)
	assert.Equal(t, "Jakarta", result.PlaceBirth)
	assert.Equal(t, "2000-01-01", result.DateBirth)
	assert.Equal(t, "2021-01-01", result.CreatedAt)
	assert.Equal(t, "2021-01-01", result.UpdatedAt)
}

func Test_Update_Failed(t *testing.T) {
	repo := student_repository.NewStudentRepositoryMock()

	service := NewStudentService(repo)

	student_repository.UpdateStudent = func(studentId int, studentPayload *entity.Student) (*dto.StudentResponse, errs.Error) {
		return nil, errs.NewInternalServerError("internal server error")
	}

	payload := &dto.StudentUpdateRequest{
		Name:       "Student 1",
		Nis:        "123456",
		PlaceBirth: "Jakarta",
		DateBirth:  "2000-01-01",
	}

	result, err := service.UpdateStudent(1, payload)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_Update_InvalidValidateData(t *testing.T) {
	repo := student_repository.NewStudentRepositoryMock()

	service := NewStudentService(repo)

	payload := &dto.StudentUpdateRequest{
		Name: "",
	}

	result, err := service.UpdateStudent(1, payload)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_GetAll_Success(t *testing.T) {
	repo := student_repository.NewStudentRepositoryMock()

	service := NewStudentService(repo)

	student_repository.FindAllStudents = func() ([]dto.StudentResponse, errs.Error) {
		return []dto.StudentResponse{
			{
				Id:   1,
				Name: "Student 1",
			},
			{
				Id:   2,
				Name: "Student 2",
			},
		}, nil
	}

	result, err := service.GetAllStudents()

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(result))
}

func Test_GetAll_Failed(t *testing.T) {
	repo := student_repository.NewStudentRepositoryMock()

	service := NewStudentService(repo)

	student_repository.FindAllStudents = func() ([]dto.StudentResponse, errs.Error) {
		return nil, errs.NewNotFoundError("student not found")
	}

	result, err := service.GetAllStudents()

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_GetById_Success(t *testing.T) {
	repo := student_repository.NewStudentRepositoryMock()

	service := NewStudentService(repo)

	student_repository.FindStudentById = func(studentId int) (*dto.StudentResponse, errs.Error) {
		return &dto.StudentResponse{
			Id:   1,
			Name: "Student 1",
		}, nil
	}

	result, err := service.GetStudentById(1)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func Test_GetById_Failed(t *testing.T) {
	repo := student_repository.NewStudentRepositoryMock()

	service := NewStudentService(repo)

	student_repository.FindStudentById = func(studentId int) (*dto.StudentResponse, errs.Error) {
		return nil, errs.NewNotFoundError("student not found")
	}

	result, err := service.GetStudentById(1)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_GetStudentDontHaveInClassroom_Success(t *testing.T) {
	repo := student_repository.NewStudentRepositoryMock()

	service := NewStudentService(repo)

	student_repository.FindStudentDontHaveInClassroom = func(classroomId int) ([]dto.StudentResponse, errs.Error) {
		return []dto.StudentResponse{
			{
				Id:   1,
				Name: "Student 1",
			},
			{
				Id:   2,
				Name: "Student 2",
			},
		}, nil
	}

	result, err := service.GetStudentDontHaveInClassroom(1)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(result))
}

func Test_GetStudentDontHaveInClassroom_Failed(t *testing.T) {
	repo := student_repository.NewStudentRepositoryMock()

	service := NewStudentService(repo)

	student_repository.FindStudentDontHaveInClassroom = func(classroomId int) ([]dto.StudentResponse, errs.Error) {
		return nil, errs.NewNotFoundError("student not found")
	}

	result, err := service.GetStudentDontHaveInClassroom(1)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_Count_Success(t *testing.T) {
	repo := student_repository.NewStudentRepositoryMock()

	service := NewStudentService(repo)

	student_repository.CountStudents = func() (int, errs.Error) {
		return 2, nil
	}

	result, err := service.TotalStudents()

	assert.Nil(t, err)
	assert.Equal(t, 2, result)
}

func Test_Count_Failed(t *testing.T) {
	repo := student_repository.NewStudentRepositoryMock()

	service := NewStudentService(repo)

	student_repository.CountStudents = func() (int, errs.Error) {
		return 0, errs.NewNotFoundError("student not found")
	}

	result, err := service.TotalStudents()

	assert.NotNil(t, err)
	assert.Equal(t, 0, result)
}

func Test_GetAllWithPagination_Success(t *testing.T) {
	repo := student_repository.NewStudentRepositoryMock()

	service := NewStudentService(repo)

	student_repository.FindAllStudentsWithPagination = func(page, perPage int) (*dto.StudentWithPagination, errs.Error) {
		return &dto.StudentWithPagination{
			Data: []dto.StudentResponse{
				{
					Id:   1,
					Name: "Student 1",
				},
				{
					Id:   2,
					Name: "Student 2",
				},
			},
			Page: dto.Page{
				CurrentPage: 1,
				PerPage:     10,
				From:        1,
				To:          2,
				TotalData:   2,
				TotalPage:   1,
				LastPage:    1,
			},
		}, nil
	}

	result, err := service.GetAllStudentsWithPagination(1, 10)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(result.Data))
	assert.Equal(t, 1, result.Page.CurrentPage)
	assert.Equal(t, 10, result.Page.PerPage)
	assert.Equal(t, 1, result.Page.From)
	assert.Equal(t, 2, result.Page.To)
	assert.Equal(t, 2, result.Page.TotalData)
	assert.Equal(t, 1, result.Page.TotalPage)
	assert.Equal(t, 1, result.Page.LastPage)
}

func Test_GetAllWithPagination_Failed(t *testing.T) {
	repo := student_repository.NewStudentRepositoryMock()

	service := NewStudentService(repo)

	student_repository.FindAllStudentsWithPagination = func(page, perPage int) (*dto.StudentWithPagination, errs.Error) {
		return nil, errs.NewNotFoundError("student not found")
	}

	result, err := service.GetAllStudentsWithPagination(1, 10)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_GetAllWithPagination_Failed_InvalidPage(t *testing.T) {
	repo := student_repository.NewStudentRepositoryMock()

	service := NewStudentService(repo)

	result, err := service.GetAllStudentsWithPagination(0, 10)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_GetAllWithPagination_Failed_InvalidPerPage(t *testing.T) {
	repo := student_repository.NewStudentRepositoryMock()

	service := NewStudentService(repo)

	result, err := service.GetAllStudentsWithPagination(1, 0)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}
