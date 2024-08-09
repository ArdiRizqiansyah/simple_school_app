package student_pg

import (
	"be-classroom/dto"
	"be-classroom/entity"
	"be-classroom/pkg/errs"
	"be-classroom/repository/student_repository"
	"database/sql"
	"math"
)

type studentRepositoryImpl struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) student_repository.StudentRepository {
	return &studentRepositoryImpl{
		db: db,
	}
}

func (sr *studentRepositoryImpl) FindAllStudents() ([]dto.StudentResponse, errs.Error) {
	var students []dto.StudentResponse

	rows, err := sr.db.Query(fetchAllStudentQuery)
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	for rows.Next() {
		var student dto.StudentResponse
		err := rows.Scan(
			&student.Id,
			&student.Name,
			&student.Nis,
			&student.PlaceBirth,
			&student.DateBirth,
			&student.CreatedAt,
			&student.UpdatedAt,
		)
		if err != nil {
			return nil, errs.NewInternalServerError(err.Error())
		}

		students = append(students, student)
	}

	return students, nil
}

func (sr *studentRepositoryImpl) FindStudentById(studentId int) (*dto.StudentResponse, errs.Error) {
	var student dto.StudentResponse

	err := sr.db.QueryRow(fetchStudentByIdQuery, studentId).Scan(
		&student.Id,
		&student.Name,
		&student.Nis,
		&student.PlaceBirth,
		&student.DateBirth,
		&student.CreatedAt,
		&student.UpdatedAt,
	)

	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &student, nil
}

func (sr *studentRepositoryImpl) FindStudentDontHaveInClassroom(classroomId int) ([]dto.StudentResponse, errs.Error) {
	var students []dto.StudentResponse

	rows, err := sr.db.Query(fetchStudentDontHaveInClassroomQuery, classroomId)
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	for rows.Next() {
		var student dto.StudentResponse
		err := rows.Scan(
			&student.Id,
			&student.Name,
			&student.Nis,
			&student.PlaceBirth,
			&student.DateBirth,
			&student.CreatedAt,
			&student.UpdatedAt,
		)
		if err != nil {
			return nil, errs.NewInternalServerError(err.Error())
		}

		students = append(students, student)
	}

	return students, nil
}

func (sr *studentRepositoryImpl) StoreStudent(studentPayload *entity.Student) (*dto.StudentResponse, errs.Error) {
	var student dto.StudentResponse

	err := sr.db.QueryRow(
		createStudentQuery,
		studentPayload.Name,
		studentPayload.Nis,
		studentPayload.PlaceBirth,
		studentPayload.DateBirth,
	).Scan(
		&student.Id,
		&student.Name,
		&student.Nis,
		&student.PlaceBirth,
		&student.DateBirth,
		&student.CreatedAt,
		&student.UpdatedAt,
	)

	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &student, nil
}

func (sr *studentRepositoryImpl) UpdateStudent(studentId int, studentPayload *entity.Student) (*dto.StudentResponse, errs.Error) {
	var student dto.StudentResponse

	err := sr.db.QueryRow(
		updateStudentQuery,
		studentId,
		studentPayload.Name,
		studentPayload.Nis,
		studentPayload.PlaceBirth,
		studentPayload.DateBirth,
	).Scan(
		&student.Id,
		&student.Name,
		&student.Nis,
		&student.PlaceBirth,
		&student.DateBirth,
		&student.CreatedAt,
		&student.UpdatedAt,
	)

	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &student, nil
}

func (sr *studentRepositoryImpl) DestroyStudent(studentId int) errs.Error {
	_, err := sr.db.Exec(deleteStudentQuery, studentId)
	if err != nil {
		return errs.NewInternalServerError(err.Error())
	}

	return nil
}

func (sr *studentRepositoryImpl) CountStudents() (int, errs.Error) {
	var count int

	err := sr.db.QueryRow(countStudentQuery).Scan(&count)
	if err != nil {
		return 0, errs.NewInternalServerError(err.Error())
	}

	return count, nil
}

func (sr *studentRepositoryImpl) FindAllStudentsWithPagination(page, perPage int) (*dto.StudentWithPagination, errs.Error) {
	var students []dto.StudentResponse
	var total int

	rows, err := sr.db.Query(fetchAllWithPaginationQuery, perPage, (page-1)*perPage)
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	for rows.Next() {
		var student dto.StudentResponse
		err := rows.Scan(
			&student.Id,
			&student.Name,
			&student.Nis,
			&student.PlaceBirth,
			&student.DateBirth,
			&student.CreatedAt,
			&student.UpdatedAt,
		)
		if err != nil {
			return nil, errs.NewInternalServerError(err.Error())
		}

		students = append(students, student)
	}

	err = sr.db.QueryRow(countStudentQuery).Scan(&total)
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	totalPages := int(math.Ceil(float64(total) / float64(perPage)))

	return &dto.StudentWithPagination{
		Data: students,
		Page: dto.Page{
			CurrentPage: page,
			PerPage:     perPage,
			From:        (page - 1) * perPage,
			To:          page * perPage,
			TotalData:   total,
			TotalPage:   totalPages,
			LastPage:    totalPages,
		},
	}, nil
}
