package classroom_service

import (
	"be-classroom/dto"
	"be-classroom/entity"
	"be-classroom/pkg/errs"
	"be-classroom/pkg/helper"
	"be-classroom/repository/classroom_repository"
)

type ClassroomService interface {
	GetAllClassrooms() ([]dto.ClassroomResponse, errs.Error)
	GetClassroomById(classroomId int) (*dto.ClassroomResponse, errs.Error)
	TotalClassrooms() (int, errs.Error)
	GetAllClassroomsWithPagination(page, perPage int) (*dto.ClassroomWithPagination, errs.Error)
	CreateClassroom(classroomPayload *dto.ClassroomCreateRequest) (*dto.ClassroomResponse, errs.Error)
	EditClassroom(classroomId int, classroomPayload *dto.ClassroomUpdateRequest) (*dto.ClassroomResponse, errs.Error)
	DeleteClassroom(classroomId int) errs.Error
}

type classroomServiceImpl struct {
	cr classroom_repository.ClassroomRepository
}

func NewClassroomService(classroomRepo classroom_repository.ClassroomRepository) ClassroomService {
	return &classroomServiceImpl{
		cr: classroomRepo,
	}
}

func (cs *classroomServiceImpl) GetAllClassrooms() ([]dto.ClassroomResponse, errs.Error) {
	classrooms, err := cs.cr.FindAllClassrooms()

	if err != nil {
		return nil, err
	}

	return classrooms, nil
}

func (cs *classroomServiceImpl) GetClassroomById(classroomId int) (*dto.ClassroomResponse, errs.Error) {
	classroom, err := cs.cr.FindClassroomById(classroomId)

	if err != nil {
		return nil, err
	}

	return classroom, nil
}

func (cs *classroomServiceImpl) CreateClassroom(classroomPayload *dto.ClassroomCreateRequest) (*dto.ClassroomResponse, errs.Error) {
	err := helper.ValidateStruct(classroomPayload)

	if err != nil {
		return nil, err
	}

	classroom := &entity.Classroom{
		Name: classroomPayload.Name,
	}

	createdClassroom, err := cs.cr.StoreClassroom(classroom)

	if err != nil {
		return nil, err
	}

	return createdClassroom, nil
}

func (cs *classroomServiceImpl) EditClassroom(classroomId int, classroomPayload *dto.ClassroomUpdateRequest) (*dto.ClassroomResponse, errs.Error) {
	err := helper.ValidateStruct(classroomPayload)

	if err != nil {
		return nil, err
	}

	classroom := &entity.Classroom{
		Name: classroomPayload.Name,
	}

	updatedClassroom, err := cs.cr.UpdateClassroom(classroomId, classroom)

	if err != nil {
		return nil, err
	}

	return updatedClassroom, nil
}

func (cs *classroomServiceImpl) DeleteClassroom(classroomId int) errs.Error {
	err := cs.cr.DestroyClassroom(classroomId)

	if err != nil {
		return err
	}

	return nil
}

func (cs *classroomServiceImpl) TotalClassrooms() (int, errs.Error) {
	total, err := cs.cr.CountClassrooms()

	if err != nil {
		return 0, err
	}

	return total, nil
}

func (cs *classroomServiceImpl) GetAllClassroomsWithPagination(page, perPage int) (*dto.ClassroomWithPagination, errs.Error) {
	classrooms, err := cs.cr.FindAllClassroomsWithPagination(page, perPage)

	if err != nil {
		return nil, err
	}

	return classrooms, nil
}
