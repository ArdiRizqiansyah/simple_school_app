package classroom_student_repository

import "be-classroom/entity"

type ClassroomStudentWithStudent struct {
	ClassroomStudent entity.ClassroomStudent
	Student          entity.Student
}

type ClassroomStudentWithStudentMapped struct {
	Id          int            `json:"id"`
	ClassroomId int            `json:"classroom_id"`
	StudentId   int            `json:"student_id"`
	CreatedAt   string         `json:"created_at"`
	UpdatedAt   string         `json:"updated_at"`
	Student     entity.Student `json:"student"`
}

func (csws *ClassroomStudentWithStudentMapped) HandleMappingClassroomStudentWithStudent(classroomStudent []ClassroomStudentWithStudent) []ClassroomStudentWithStudentMapped {
	classroomStudentMapped := []ClassroomStudentWithStudentMapped{}

	for _, cs := range classroomStudent {
		classroomStudentMapped = append(classroomStudentMapped, ClassroomStudentWithStudentMapped{
			Id:          cs.ClassroomStudent.Id,
			ClassroomId: cs.ClassroomStudent.ClassroomId,
			StudentId:   cs.ClassroomStudent.StudentId,
			CreatedAt:   cs.ClassroomStudent.CreatedAt,
			UpdatedAt:   cs.ClassroomStudent.UpdatedAt,
			Student: entity.Student{
				Id:         cs.Student.Id,
				Name:       cs.Student.Name,
				Nis:        cs.Student.Nis,
				PlaceBirth: cs.Student.PlaceBirth,
				DateBirth:  cs.Student.DateBirth,
				CreatedAt:  cs.Student.CreatedAt,
				UpdatedAt:  cs.Student.UpdatedAt,
			},
		})
	}

	return classroomStudentMapped
}

func (csws *ClassroomStudentWithStudentMapped) HandleMappingClassroomStudentWithStudentById(classroomStudent ClassroomStudentWithStudent) *ClassroomStudentWithStudentMapped {
	return &ClassroomStudentWithStudentMapped{
		Id:          classroomStudent.ClassroomStudent.Id,
		ClassroomId: classroomStudent.ClassroomStudent.ClassroomId,
		StudentId:   classroomStudent.ClassroomStudent.StudentId,
		CreatedAt:   classroomStudent.ClassroomStudent.CreatedAt,
		UpdatedAt:   classroomStudent.ClassroomStudent.UpdatedAt,
		Student: entity.Student{
			Id:         classroomStudent.Student.Id,
			Name:       classroomStudent.Student.Name,
			Nis:        classroomStudent.Student.Nis,
			PlaceBirth: classroomStudent.Student.PlaceBirth,
			DateBirth:  classroomStudent.Student.DateBirth,
			CreatedAt:  classroomStudent.Student.CreatedAt,
			UpdatedAt:  classroomStudent.Student.UpdatedAt,
		},
	}
}
