package classroom_pg

import (
	"be-classroom/dto"
	"be-classroom/entity"
	"be-classroom/pkg/errs"
	"be-classroom/repository/classroom_repository"
	"database/sql"
	"math"
)

type classroomRepositoryImpl struct {
	db *sql.DB
}

func NewClassroomRepository(db *sql.DB) classroom_repository.ClassroomRepository {
	return &classroomRepositoryImpl{
		db: db,
	}
}

func (cr *classroomRepositoryImpl) FindAllClassrooms() ([]dto.ClassroomResponse, errs.Error) {
	var classrooms []dto.ClassroomResponse

	rows, err := cr.db.Query(fetchAllClassroomQuery)
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	for rows.Next() {
		var classroom dto.ClassroomResponse
		err := rows.Scan(
			&classroom.Id,
			&classroom.Name,
			&classroom.CreatedAt,
			&classroom.UpdatedAt,
		)
		if err != nil {
			return nil, errs.NewInternalServerError(err.Error())
		}

		classrooms = append(classrooms, classroom)
	}

	return classrooms, nil
}

func (cr *classroomRepositoryImpl) FindClassroomById(classroomId int) (*dto.ClassroomResponse, errs.Error) {
	var classroom dto.ClassroomResponse

	err := cr.db.QueryRow(fetchClassroomByIdQuery, classroomId).Scan(
		&classroom.Id,
		&classroom.Name,
		&classroom.CreatedAt,
		&classroom.UpdatedAt,
	)

	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &classroom, nil
}

func (cr *classroomRepositoryImpl) StoreClassroom(classroomPayload *entity.Classroom) (*dto.ClassroomResponse, errs.Error) {
	var classroom dto.ClassroomResponse

	err := cr.db.QueryRow(
		createClassroomQuery,
		classroomPayload.Name,
	).Scan(
		&classroom.Id,
		&classroom.Name,
		&classroom.CreatedAt,
		&classroom.UpdatedAt,
	)

	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &classroom, nil
}

func (cr *classroomRepositoryImpl) UpdateClassroom(classroomId int, classroomPayload *entity.Classroom) (*dto.ClassroomResponse, errs.Error) {
	var classroom dto.ClassroomResponse

	err := cr.db.QueryRow(
		updateClassroomQuery,
		classroomId,
		classroomPayload.Name,
	).Scan(
		&classroom.Id,
		&classroom.Name,
		&classroom.CreatedAt,
		&classroom.UpdatedAt,
	)

	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &classroom, nil
}

func (cr *classroomRepositoryImpl) DestroyClassroom(classroomId int) errs.Error {
	_, err := cr.db.Exec(deleteClassroomQuery, classroomId)
	if err != nil {
		return errs.NewInternalServerError(err.Error())
	}

	return nil
}

func (cr *classroomRepositoryImpl) CountClassrooms() (int, errs.Error) {
	var count int

	err := cr.db.QueryRow(countClassroomQuery).Scan(&count)
	if err != nil {
		return 0, errs.NewInternalServerError(err.Error())
	}

	return count, nil
}

func (cr *classroomRepositoryImpl) FindAllClassroomsWithPagination(page, perPage int) (*dto.ClassroomWithPagination, errs.Error) {
	var classrooms []dto.ClassroomResponse
	var total int

	rows, err := cr.db.Query(fetchAllWithPaginationQuery, perPage, (page-1)*perPage)
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	for rows.Next() {
		var classroom dto.ClassroomResponse
		err := rows.Scan(
			&classroom.Id,
			&classroom.Name,
			&classroom.CreatedAt,
			&classroom.UpdatedAt,
		)
		if err != nil {
			return nil, errs.NewInternalServerError(err.Error())
		}

		classrooms = append(classrooms, classroom)
	}

	total, err = cr.CountClassrooms()
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	totalPages := int(math.Ceil(float64(total) / float64(perPage)))

	return &dto.ClassroomWithPagination{
		Data: classrooms,
		Page: dto.Page{
			CurrentPage: page,
			PerPage:     perPage,
			From:        (page - 1) * perPage,
			To:          (page - 1) + len(classrooms),
			TotalData:   total,
			TotalPage:   totalPages,
			LastPage:    totalPages,
		},
	}, nil
}
