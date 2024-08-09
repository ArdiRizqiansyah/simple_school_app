package classroom_student_pg

const (
	GetAllInClassroomWithStudent = `
		SELECT 
			cs.id, 
			cs.classroom_id, 
			cs.student_id, 
			cs.created_at, 
			cs.updated_at,
			s.id,
			s.name, 
			s.nis,
			s.place_birth,
			s.date_birth,
			s.created_at,
			s.updated_at
		FROM 
			classroom_students cs 
		JOIN 
			students s ON cs.student_id = s.id
		WHERE
			cs.classroom_id = $1
	`

	InsertClassroomStudent = `
		INSERT INTO classroom_students (classroom_id, student_id)
		VALUES ($1, $2)
		RETURNING id, classroom_id, student_id, created_at, updated_at
	`

	DeleteClassroomStudent = `
		DELETE FROM classroom_students
		WHERE id = $1
		RETURNING id, classroom_id, student_id, created_at, updated_at
	`
)
