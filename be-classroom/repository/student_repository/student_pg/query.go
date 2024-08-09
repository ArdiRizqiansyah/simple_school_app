package student_pg

const (
	createStudentQuery = `
		INSERT INTO students (name, nis, place_birth, date_birth)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, nis, place_birth, date_birth, created_at, updated_at
	`

	fetchAllStudentQuery = `
		SELECT id, name, nis, place_birth, date_birth, created_at, updated_at
		FROM students
	`

	fetchStudentByIdQuery = `
		SELECT id, name, nis, place_birth, date_birth, created_at, updated_at
		FROM students
		WHERE id = $1
	`

	fetchStudentDontHaveInClassroomQuery = `
		SELECT id, name, nis, place_birth, date_birth, created_at, updated_at
		FROM students
		WHERE id NOT IN (
			SELECT student_id
			FROM classroom_students
			WHERE classroom_id = $1
		)
	`

	updateStudentQuery = `
		UPDATE students
		SET
			name = $2,
			nis = $3,
			place_birth = $4,
			date_birth = $5
		WHERE
			id = $1
		RETURNING id, name, nis, place_birth, date_birth, created_at, updated_at
	`

	deleteStudentQuery = `
		DELETE FROM students
		WHERE id = $1
		RETURNING id, name, nis, place_birth, date_birth, created_at, updated_at
	`

	countStudentQuery = `
		SELECT COUNT(*)
		FROM students
	`

	fetchAllWithPaginationQuery = `
		SELECT id, name, nis, place_birth, date_birth, created_at, updated_at
		FROM students
		LIMIT $1 OFFSET $2
	`
)
