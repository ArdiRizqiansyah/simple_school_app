package classroom_student_pg

import (
	"be-classroom/pkg/errs"
	"be-classroom/repository/classroom_student_repository"
	"database/sql"
)

type classroomRepositoryImpl struct {
	db *sql.DB
}

func NewClassroomStudentRepository(db *sql.DB) classroom_student_repository.ClassroomStudentRepository {
	return &classroomRepositoryImpl{
		db: db,
	}
}

func (csr *classroomRepositoryImpl) FindAllClassroomStudents(classroomId int) ([]classroom_student_repository.ClassroomStudentWithStudentMapped, errs.Error) {
	classroomStudents := []classroom_student_repository.ClassroomStudentWithStudent{}
	rows, err := csr.db.Query(GetAllInClassroomWithStudent, classroomId)

	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	for rows.Next() {
		classroomStudent := classroom_student_repository.ClassroomStudentWithStudent{}

		err = rows.Scan(
			&classroomStudent.ClassroomStudent.Id,
			&classroomStudent.ClassroomStudent.ClassroomId,
			&classroomStudent.ClassroomStudent.StudentId,
			&classroomStudent.ClassroomStudent.CreatedAt,
			&classroomStudent.ClassroomStudent.UpdatedAt,
			&classroomStudent.Student.Id,
			&classroomStudent.Student.Name,
			&classroomStudent.Student.Nis,
			&classroomStudent.Student.PlaceBirth,
			&classroomStudent.Student.DateBirth,
			&classroomStudent.Student.CreatedAt,
			&classroomStudent.Student.UpdatedAt,
		)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errs.NewNotFoundError("Classroom Student Not Found")
			}

			return nil, errs.NewInternalServerError(err.Error())
		}

		classroomStudents = append(classroomStudents, classroomStudent)
	}

	result := classroom_student_repository.ClassroomStudentWithStudentMapped{}
	return result.HandleMappingClassroomStudentWithStudent(classroomStudents), nil
}

func (csr *classroomRepositoryImpl) StoreClassroomStudent(classroomId int, studentId []int) errs.Error {
	tx, err := csr.db.Begin()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError(err.Error())
	}

	stmt, err := tx.Prepare(InsertClassroomStudent)

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	for _, sId := range studentId {
		_, err = stmt.Exec(classroomId, sId)

		if err != nil {
			tx.Rollback()
			return errs.NewInternalServerError(err.Error())
		}
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError(err.Error())
	}

	return nil
}

func (csr *classroomRepositoryImpl) DestroyClassroomStudent(classroomStudentId int) errs.Error {
	// hanya habus data di tabel classroom_student
	_, err := csr.db.Exec(DeleteClassroomStudent, classroomStudentId)

	if err != nil {
		return errs.NewInternalServerError(err.Error())
	}

	return nil
}
