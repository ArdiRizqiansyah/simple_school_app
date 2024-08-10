package classroom_service

import (
	"be-classroom/dto"
	"be-classroom/entity"
	"be-classroom/pkg/errs"
	"be-classroom/repository/classroom_repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Delete_Success(t *testing.T) {
	repo := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomService(repo)

	classroom_repository.DestroyClassroom = func(classroomId int) errs.Error {
		return nil
	}

	err := service.DeleteClassroom(1)

	assert.Nil(t, err)
}

func Test_Delete_Failed(t *testing.T) {
	repo := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomService(repo)

	classroom_repository.DestroyClassroom = func(classroomId int) errs.Error {
		return errs.NewNotFoundError("classroom not found")
	}

	err := service.DeleteClassroom(1)

	assert.NotNil(t, err)
}

func Test_Create_Success(t *testing.T) {
	repo := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomService(repo)

	classroom_repository.StoreClassroom = func(classroomPayload *entity.Classroom) (*dto.ClassroomResponse, errs.Error) {
		return &dto.ClassroomResponse{
			Id:   1,
			Name: "Classroom 1",
		}, nil
	}

	payload := &dto.ClassroomCreateRequest{
		Name: "Classroom 1",
	}

	result, err := service.CreateClassroom(payload)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.Id)
	assert.Equal(t, "Classroom 1", result.Name)
}

func Test_Create_Failed(t *testing.T) {
	repo := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomService(repo)

	classroom_repository.StoreClassroom = func(classroomPayload *entity.Classroom) (*dto.ClassroomResponse, errs.Error) {
		return nil, errs.NewBadRequestError("bad request")
	}

	payload := &dto.ClassroomCreateRequest{
		Name: "Classroom 1",
	}

	result, err := service.CreateClassroom(payload)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_Create_InvalidValidateData(t *testing.T) {
	repo := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomService(repo)

	payload := &dto.ClassroomCreateRequest{
		Name: "",
	}

	result, err := service.CreateClassroom(payload)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_Edit_Success(t *testing.T) {
	repo := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomService(repo)

	classroom_repository.UpdateClassroom = func(classroomId int, classroomPayload *entity.Classroom) (*dto.ClassroomResponse, errs.Error) {
		return &dto.ClassroomResponse{
			Id:   1,
			Name: "Classroom 1",
		}, nil
	}

	payload := &dto.ClassroomUpdateRequest{
		Name: "Classroom 1",
	}

	result, err := service.EditClassroom(1, payload)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.Id)
	assert.Equal(t, "Classroom 1", result.Name)
}

func Test_Edit_Failed(t *testing.T) {
	repo := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomService(repo)

	classroom_repository.UpdateClassroom = func(classroomId int, classroomPayload *entity.Classroom) (*dto.ClassroomResponse, errs.Error) {
		return nil, errs.NewBadRequestError("bad request")
	}

	payload := &dto.ClassroomUpdateRequest{
		Name: "Classroom 1",
	}

	result, err := service.EditClassroom(1, payload)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_Edit_InvalidValidateData(t *testing.T) {
	repo := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomService(repo)

	payload := &dto.ClassroomUpdateRequest{
		Name: "",
	}

	result, err := service.EditClassroom(1, payload)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_GetAll_Success(t *testing.T) {
	repo := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomService(repo)

	classroom_repository.FindAllClassrooms = func() ([]dto.ClassroomResponse, errs.Error) {
		return []dto.ClassroomResponse{
			{
				Id:   1,
				Name: "Classroom 1",
			},
			{
				Id:   2,
				Name: "Classroom 2",
			},
		}, nil
	}

	result, err := service.GetAllClassrooms()

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(result))
}

func Test_GetAll_Failed(t *testing.T) {
	repo := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomService(repo)

	classroom_repository.FindAllClassrooms = func() ([]dto.ClassroomResponse, errs.Error) {
		return nil, errs.NewNotFoundError("classroom not found")
	}

	result, err := service.GetAllClassrooms()

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_GetById_Success(t *testing.T) {
	repo := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomService(repo)

	classroom_repository.FindClassroomById = func(classroomId int) (*dto.ClassroomResponse, errs.Error) {
		return &dto.ClassroomResponse{
			Id:   1,
			Name: "Classroom 1",
		}, nil
	}

	result, err := service.GetClassroomById(1)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.Id)
	assert.Equal(t, "Classroom 1", result.Name)
}

func Test_GetById_Failed(t *testing.T) {
	repo := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomService(repo)

	classroom_repository.FindClassroomById = func(classroomId int) (*dto.ClassroomResponse, errs.Error) {
		return nil, errs.NewNotFoundError("classroom not found")
	}

	result, err := service.GetClassroomById(1)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_GetById_InvalidId(t *testing.T) {
	repo := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomService(repo)

	result, err := service.GetClassroomById(0)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_GetAllWithPagination_Success(t *testing.T) {
	repo := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomService(repo)

	classroom_repository.FindAllClassroomsWithPagination = func(page, perPage int) (*dto.ClassroomWithPagination, errs.Error) {
		return &dto.ClassroomWithPagination{
			Data: []dto.ClassroomResponse{
				{
					Id:   1,
					Name: "Classroom 1",
				},
				{
					Id:   2,
					Name: "Classroom 2",
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

	result, err := service.GetAllClassroomsWithPagination(1, 10)

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
	repo := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomService(repo)

	classroom_repository.FindAllClassroomsWithPagination = func(page, perPage int) (*dto.ClassroomWithPagination, errs.Error) {
		return nil, errs.NewNotFoundError("classroom not found")
	}

	result, err := service.GetAllClassroomsWithPagination(1, 10)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_Total_Success(t *testing.T) {
	repo := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomService(repo)

	classroom_repository.CountClassrooms = func() (int, errs.Error) {
		return 10, nil
	}

	result, err := service.TotalClassrooms()

	assert.Nil(t, err)
	assert.Equal(t, 10, result)
}

func Test_Total_Failed(t *testing.T) {
	repo := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomService(repo)

	classroom_repository.CountClassrooms = func() (int, errs.Error) {
		return 0, errs.NewNotFoundError("classroom not found")
	}

	result, err := service.TotalClassrooms()

	assert.NotNil(t, err)
	assert.Equal(t, 0, result)
}

func Test_Total_InvalidId(t *testing.T) {
	repo := classroom_repository.NewClassroomRepositoryMock()

	service := NewClassroomService(repo)

	result, err := service.TotalClassrooms()

	assert.NotNil(t, err)
	assert.Equal(t, 0, result)
}
