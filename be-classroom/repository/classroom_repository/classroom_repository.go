package classroom_repository

import (
	"be-classroom/dto"
	"be-classroom/entity"
	"be-classroom/pkg/errs"
)

type ClassroomRepository interface {
	FindAllClassrooms() ([]dto.ClassroomResponse, errs.Error)
	FindClassroomById(classroomId int) (*dto.ClassroomResponse, errs.Error)
	CountClassrooms() (int, errs.Error)
	FindAllClassroomsWithPagination(page, perPage int) (*dto.ClassroomWithPagination, errs.Error)
	StoreClassroom(classroomPayload *entity.Classroom) (*dto.ClassroomResponse, errs.Error)
	UpdateClassroom(classroomId int, classroomPayload *entity.Classroom) (*dto.ClassroomResponse, errs.Error)
	DestroyClassroom(classroomId int) errs.Error
}
