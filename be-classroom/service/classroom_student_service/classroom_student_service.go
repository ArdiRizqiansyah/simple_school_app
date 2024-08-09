package classroom_student_service

import (
	"be-classroom/dto"
	"be-classroom/pkg/errs"
	"be-classroom/pkg/helper"
	"be-classroom/repository/classroom_repository"
	"be-classroom/repository/classroom_student_repository"
)

type ClassroomStudentService interface {
	GetAllClassroomStudents(classroomId int) ([]dto.ClassroomStudentResponse, errs.Error)
	CreateClassroomStudent(classroom_id int, classroomStudentPayload *dto.ClassroomStudentCreateRequest) errs.Error
	DeleteClassroomStudent(classroomStudentId int) errs.Error
}

type classroomStudentServiceImpl struct {
	csr classroom_student_repository.ClassroomStudentRepository
	cr  classroom_repository.ClassroomRepository
}

func NewClassroomStudentService(classroomStudentRepo classroom_student_repository.ClassroomStudentRepository, classroomRepo classroom_repository.ClassroomRepository) ClassroomStudentService {
	return &classroomStudentServiceImpl{
		csr: classroomStudentRepo,
		cr:  classroomRepo,
	}
}

func (css *classroomStudentServiceImpl) GetAllClassroomStudents(classroomId int) ([]dto.ClassroomStudentResponse, errs.Error) {
	classroomStudents, err := css.csr.FindAllClassroomStudents(classroomId)

	if err != nil {
		return nil, err
	}

	classroomStudentResponses := []dto.ClassroomStudentResponse{}

	for _, classroomStudent := range classroomStudents {
		classroomStudentResponses = append(classroomStudentResponses, dto.ClassroomStudentResponse{
			Id:          classroomStudent.Id,
			ClassroomId: classroomStudent.ClassroomId,
			StudentId:   classroomStudent.StudentId,
			CreatedAt:   classroomStudent.CreatedAt,
			UpdatedAt:   classroomStudent.UpdatedAt,
			Student: dto.StudentResponse{
				Id:         classroomStudent.Student.Id,
				Name:       classroomStudent.Student.Name,
				Nis:        classroomStudent.Student.Nis,
				PlaceBirth: classroomStudent.Student.PlaceBirth,
				DateBirth:  classroomStudent.Student.DateBirth,
				CreatedAt:  classroomStudent.Student.CreatedAt,
				UpdatedAt:  classroomStudent.Student.UpdatedAt,
			},
		})
	}

	return classroomStudentResponses, nil
}

func (css *classroomStudentServiceImpl) CreateClassroomStudent(clasroomId int, classroomStudentPayload *dto.ClassroomStudentCreateRequest) errs.Error {
	err := helper.ValidateStruct(classroomStudentPayload)

	if err != nil {
		return err
	}

	// cek apakah classroom_id sudah ada di database
	classroom, err := css.cr.FindClassroomById(clasroomId)

	if err != nil {
		return err
	}

	err = css.csr.StoreClassroomStudent(classroom.Id, classroomStudentPayload.StudentId)

	if err != nil {
		return err
	}

	return nil
}

func (css *classroomStudentServiceImpl) DeleteClassroomStudent(classroomStudentId int) errs.Error {
	err := css.csr.DestroyClassroomStudent(classroomStudentId)

	if err != nil {
		return err
	}

	return nil
}
