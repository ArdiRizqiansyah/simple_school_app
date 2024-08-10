package classroom_repository

import (
	"be-classroom/dto"
	"be-classroom/entity"
	"be-classroom/pkg/errs"
)

var (
	FindAllClassrooms               func() ([]dto.ClassroomResponse, errs.Error)
	FindClassroomById               func(classroomId int) (*dto.ClassroomResponse, errs.Error)
	CountClassrooms                 func() (int, errs.Error)
	FindAllClassroomsWithPagination func(page, perPage int) (*dto.ClassroomWithPagination, errs.Error)
	StoreClassroom                  func(classroomPayload *entity.Classroom) (*dto.ClassroomResponse, errs.Error)
	UpdateClassroom                 func(classroomId int, classroomPayload *entity.Classroom) (*dto.ClassroomResponse, errs.Error)
	DestroyClassroom                func(classroomId int) errs.Error
)

type classroomRepositoryMock struct{}

func NewClassroomRepositoryMock() ClassroomRepository {
	return &classroomRepositoryMock{}
}

func (cm *classroomRepositoryMock) FindAllClassrooms() ([]dto.ClassroomResponse, errs.Error) {
	return FindAllClassrooms()
}

func (cm *classroomRepositoryMock) FindClassroomById(classroomId int) (*dto.ClassroomResponse, errs.Error) {
	return FindClassroomById(classroomId)
}

func (cm *classroomRepositoryMock) CountClassrooms() (int, errs.Error) {
	return CountClassrooms()
}

func (cm *classroomRepositoryMock) FindAllClassroomsWithPagination(page, perPage int) (*dto.ClassroomWithPagination, errs.Error) {
	return FindAllClassroomsWithPagination(page, perPage)
}

func (cm *classroomRepositoryMock) StoreClassroom(classroomPayload *entity.Classroom) (*dto.ClassroomResponse, errs.Error) {
	return StoreClassroom(classroomPayload)
}

func (cm *classroomRepositoryMock) UpdateClassroom(classroomId int, classroomPayload *entity.Classroom) (*dto.ClassroomResponse, errs.Error) {
	return UpdateClassroom(classroomId, classroomPayload)
}

func (cm *classroomRepositoryMock) DestroyClassroom(classroomId int) errs.Error {
	return DestroyClassroom(classroomId)
}
