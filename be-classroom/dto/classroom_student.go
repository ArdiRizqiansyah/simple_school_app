package dto

type ClassroomStudentCreateRequest struct {
	StudentId []int `json:"student_id" valid:"required~Student ID can't be empty" example:"[1,2,3]"`
}

type ClassroomStudentResponse struct {
	Id          int             `json:"id"`
	ClassroomId int             `json:"classroom_id"`
	StudentId   int             `json:"student_id"`
	Student     StudentResponse `json:"student"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
	// Classroom   ClassroomResponse `json:"classroom"`
}
