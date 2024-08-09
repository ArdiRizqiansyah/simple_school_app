package classroom_pg

const (
	createClassroomQuery = `
		INSERT INTO classrooms (name)
		VALUES ($1)
		RETURNING id, name, created_at, updated_at
	`

	fetchAllClassroomQuery = `
		SELECT id, name, created_at, updated_at
		FROM classrooms
	`

	fetchClassroomByIdQuery = `
		SELECT id, name, created_at, updated_at
		FROM classrooms
		WHERE id = $1
	`

	updateClassroomQuery = `
		UPDATE classrooms
		SET 
			name = $2
		WHERE 
			id = $1
		RETURNING id, name, created_at, updated_at
	`

	deleteClassroomQuery = `
		DELETE FROM classrooms
		WHERE id = $1
		RETURNING id, name, created_at, updated_at
	`

	countClassroomQuery = `
		SELECT COUNT(*)
		FROM classrooms
	`

	fetchAllWithPaginationQuery = `
		SELECT id, name, created_at, updated_at
		FROM classrooms
		LIMIT $1 OFFSET $2
	`
)
