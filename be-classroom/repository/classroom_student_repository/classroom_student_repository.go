package classroom_student_repository

import (
	"be-classroom/pkg/errs"
)

type ClassroomStudentRepository interface {
	FindAllClassroomStudents(classroomId int) ([]ClassroomStudentWithStudentMapped, errs.Error)
	StoreClassroomStudent(classroomId int, studentId []int) errs.Error
	DestroyClassroomStudent(classroomStudentId int) errs.Error
}
